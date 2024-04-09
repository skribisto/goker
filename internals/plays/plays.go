package plays

import (
	"goker/pkg/cards"
	"goker/pkg/log"
	scores "goker/pkg/poker"
)

type Play struct {
	ID string // Unique identifier for the game
	//StartTime  time.Time        // Time when the game started
	//EndTime    time.Time        // Time when the game ended
	Blinds  BlindInfo      // Information about blinds
	Players []*Player      // ordered, names should be unique !
	Bets    *map[int][]Bet // History of bets made per Rounds
	Deck    *cards.Deck    // whole ordered (but random) deck
	Round   int            // how many cards are revealed
}

type Bet struct {
	PlayerID int // ID of the player making the bet
	Amount   int // Amount of the bet
	//Timestamp time.Time // Time when the bet was made
}

// BlindInfo represents information about blinds.
type BlindInfo struct {
	SmallBlind int // Small blind amount
	BigBlind   int // Big blind amount. Not always twice the small blind ? BigBlind can also be used to compute bet limits
	//CurrentRound int // Current round of blinds
}

// Player represents a player in the game.
type Player struct {
	ID    int    // Player place in play ("at the table")
	Name  string // Player's name
	Stack int    // Player's stack (chips)
	//Hand       *cards.Hand // Player's hole cards
	//---- could be removed ? if we know players place in the Play, and Deck is ordered, we can deduce this
	StillPlays bool // Last action taken by the player (e.g., bet, fold)
	//Strategy   int
}

//Strategy
/*
-> take into account when a player reveal cards when folding (or winning without showdown) or not
-> keep track of other's strategy
-> Different lever:
	aggressivity in bets (i.e. confidence in its strategy)
	Adjustment according to stack size compared to others
	Adjustment according to time spend at a table / blinds values

*/

func NewPlayer(name string) (*Player, error) {
	return &Player{
		ID:         0,
		Name:       name,
		Stack:      10000,
		StillPlays: true,
	}, nil
}

func NewPlay(players []*Player) (*Play, error) {
	d, err := cards.NewDeck()
	if err != nil {
		return nil, log.Errorf("could not instantiate NewDeck %w", err)
	}

	if len(players) < 2 {
		return nil, log.Error("not enough player, goker is not fun alone")
	}

	var betsMap = make(map[int][]Bet)

	for i := range players {
		players[i].ID = i
		//log.Debugf("%v", players[i].Hand)
	}

	p := &Play{
		//ID:        string, // Example function to generate a unique ID
		//StartTime:  time.Now(),
		Blinds:  BlindInfo{SmallBlind: 5, BigBlind: 10},
		Players: players,
		Bets:    &betsMap,
		Deck:    d,
		Round:   0,
	}

	log.Debug("Player0 is always the Dealer")

	bigBlindPlayerId := 2 % len(p.Players) //Manages heads up
	log.Debugf("bigBlindPlayerId : %v", bigBlindPlayerId)

	if err := p.PutBet(1, p.Blinds.SmallBlind); err != nil {
		return nil, err
	}
	if err := p.PutBet(bigBlindPlayerId, p.Blinds.BigBlind); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Play) ComputePot() (int, error) {
	pot := 0
	for round := 0; round < p.Round; round++ {
		for _, bet := range (*p.Bets)[round] {
			pot += bet.Amount
		}
	}
	if pot == 0 {
		return 0, log.Error("pot cannot be 0") // Blinds
	}
	return pot, nil
}

//
func (p *Play) WinRound(winnerIds []int) error {
	//Prints the player(s) who wins, and how much
	p.Round = 4 //recognize play is ended
	pot, err := p.ComputePot()
	if err != nil {
		return err
	}
	//now would be a good time to split the pot

	log.Infof("Player %v won %v$", winnerIds, pot)
	for _, winnerId := range winnerIds {
		gains := int(pot / (len(winnerIds)))
		p.Players[winnerId].Stack += gains
		log.Infof("Player %v wins %v, GG !", p.Players[winnerId].Name, gains)
	}

	return nil
}

//Run every time a player answers to evaluate if round can end (by win, or go to next round)
func (p *Play) EndRound() error {
	//Players still playing needs to have all the same sum of bets && at least 2 bets for blinds player
	var playersStillPlaying []int
	bigBlindPlayerId := 2 % len(p.Players) //Manages heads up

	for i := range p.Players {
		if p.Players[i].StillPlays {
			playersStillPlaying = append(playersStillPlaying, p.Players[i].ID)
		}
	}
	if len(playersStillPlaying) == 1 {
		return p.WinRound([]int{playersStillPlaying[0]})
	}

	sumBetAmountByPlayerId := make(map[int]int)
	sumBetByPlayerId := make(map[int]int) // could be used later for max 3 bet rounds, right now only to manage blind bets

	for i := range (*p.Bets)[p.Round] {
		sumBetByPlayerId[(*p.Bets)[p.Round][i].PlayerID] += 1
		sumBetAmountByPlayerId[(*p.Bets)[p.Round][i].PlayerID] += (*p.Bets)[p.Round][i].Amount
	}

	if p.Round == 0 {
		if sumBetByPlayerId[1] < 2 && p.Players[1].StillPlays {
			return log.Error("Player 1 did not speak though he/she put the small blind")
		}
		if sumBetByPlayerId[bigBlindPlayerId] < 2 && p.Players[bigBlindPlayerId].StillPlays {
			return log.Errorf("Player %v did not speak though he/she put the small blind", bigBlindPlayerId)
		}
	}

	for id := range playersStillPlaying {
		if sumBetAmountByPlayerId[id] != sumBetAmountByPlayerId[playersStillPlaying[0]] {
			return log.Error("Not all still playing players have bet the same")
		}
	}

	log.Debug("We can go to next round")

	if p.Round == 3 {
		//To move into Showdown func
		playerScoreCards := make(map[int]*scores.ScoreCard, len(playersStillPlaying))

		for id := range playersStillPlaying {
			var playerHand []cards.Card // board + player cards

			playerCards, err := p.GetPlayerCards(id)
			if err != nil {
				return err
			}
			board, err := p.GetBoard()
			if err != nil {
				return err
			}
			for i := range *board {
				playerHand = append(playerHand, (*board)[i])
			}
			for i := range *playerCards {
				playerHand = append(playerHand, (*playerCards)[i])
			}

			sc, err := scores.Score(playerHand)
			if err != nil {
				return err
			}

			playerScoreCards[id] = sc
		}

		highestScoreCard := playerScoreCards[0]
		winnerIds := []int{0}

		for playerId := 1; playerId < len(playerScoreCards); playerId++ {
			comparator, err := scores.CompareScoreCards(playerScoreCards[playerId], highestScoreCard)
			if err != nil {
				return err
			}
			switch comparator {
			case 1:
				highestScoreCard = playerScoreCards[playerId]
				winnerIds = []int{playerId}
			case -1:
				log.Debugf("player %v did not beat winner", playerId)
			default:
				log.Debug("Oh fuck, a tie...")
				winnerIds = append(winnerIds, playerId)
			}
		}
		return p.WinRound(winnerIds)
	}
	p.Round++

	return nil
}

func (p *Play) BeginRound() error {
	if p.Round == 0 {
		return nil
	}
	potValue, err := p.ComputePot()
	if err != nil {
		return err
	}
	log.Infof("POT value : %v$", potValue)

	board, err := p.GetBoard()
	if err != nil {
		return err
	}
	log.Infof("BOARD: %v", *board)
	return nil
}

func (p *Play) CanCheck(playerId int) (bool, error) {
	//log.Debugf("Player %v tries to check, checking if she can", playerId)
	playerLastBet := 0
	currentRoundBets := (*p.Bets)[p.Round]
	for _, bet := range currentRoundBets {
		if bet.PlayerID == playerId {
			playerLastBet = bet.Amount
		}
	}

	if len(currentRoundBets) == 0 {
		return true, nil
	} else if (*p.Bets)[p.Round][len(currentRoundBets)-1].Amount > playerLastBet {
		return false, nil
	}
	return true, nil
}

func (p *Play) PutBet(playerId int, amount int) error {
	if err := p.checkPlayerStillPlays(playerId); err != nil {
		return err
	}
	stack := p.Players[playerId].Stack
	if stack-amount < 0 {
		return log.Error("not enough money to bet that")
	}
	if amount == 0 {
		canCheck, err := p.CanCheck(playerId)
		if err != nil {
			return err
		} else if !canCheck {
			return log.Error("cannot check")
		}
	} else if amount > 0 && amount < p.Blinds.SmallBlind {
		return log.Error("cannot bet lower than small blind")
	}

	p.Players[playerId].StillPlays = true
	p.Players[playerId].Stack -= amount

	(*p.Bets)[p.Round] = append((*p.Bets)[p.Round], Bet{PlayerID: playerId, Amount: amount})

	return nil
}

func (p *Play) Fold(playerId int) error {
	if err := p.checkPlayerStillPlays(playerId); err != nil {
		return err
	}
	p.Players[playerId].StillPlays = false

	//Try blindly to EndRound, there might be no need to continue
	if err := p.EndRound(); err != nil {
		log.Debugf("Because player %v is folding we tried unsuccessfully to EndRound, %w", playerId, err)
	}

	return nil
}

func (p *Play) GetBoard() (*[]cards.Card, error) {
	var board []cards.Card
	deck := *p.Deck
	startIndex := len(p.Players) * 2

	endIndex := startIndex + p.Round //1st card of flop, or turn or river

	if p.Round >= 1 {
		endIndex += 2
	}
	if endIndex > len(deck) {
		return nil, log.Error("too many players, go out of deck cards")
	}

	board = deck[startIndex:endIndex]

	return &board, nil
}

func (p *Play) GetPlayerCards(playerId int) (*[]cards.Card, error) {
	var playerCards []cards.Card

	if playerId > len(p.Players) {
		return nil, log.Error("Player not in this play")
	}
	deck := *p.Deck
	startIndex := playerId * 2
	endIndex := startIndex + 2

	playerCards = deck[startIndex:endIndex]

	//log.Debugf("Player %v cards : %v", p.Players[playerId].Name, playerCards)
	return &playerCards, nil
}

func (p *Play) checkPlayerStillPlays(playerId int) error {
	if playerId > len(p.Players) {
		return log.Error("Player not in this play")
	}
	if !p.Players[playerId].StillPlays {
		return log.Error("player does not play anymore")
	}

	return nil
}

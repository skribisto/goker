package plays

import (
	"github.com/skribisto/goker/pkg/cards"
	"github.com/skribisto/goker/pkg/log"
	scores "github.com/skribisto/goker/pkg/poker"
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
	Strategy   int
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

func NewPlayer(name string, strategy int) (*Player, error) {
	return &Player{
		ID:         0,
		Name:       name,
		Stack:      100,
		StillPlays: true,
		Strategy:   strategy,
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

	log.Debug("Dealer is always last player")

	smallBlindPlayerId := 0
	bigBlindPlayerId := 1

	//Manages heads up
	if len(p.Players) == 2 {
		smallBlindPlayerId = 1
		bigBlindPlayerId = 0
	}
	log.Debugf("bigBlindPlayerId : %v", bigBlindPlayerId)
	log.Debugf("smallBlindPlayerId : %v", smallBlindPlayerId)

	if err := p.PutBet(smallBlindPlayerId, p.Blinds.SmallBlind); err != nil {
		return nil, err
	}
	if err := p.PutBet(bigBlindPlayerId, p.Blinds.BigBlind); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Play) GetPot() (int, error) {
	pot := 0
	for round := 0; round < p.Round; round++ {
		for _, bet := range (*p.Bets)[round] {
			pot += bet.Amount
		}
	}
	if p.Round != 0 && pot == 0 {
		return 0, log.Error("pot cannot be 0") // Blinds
	}
	return pot, nil
}

func (p *Play) WinRound(winnerIds []int) error {
	//Prints the player(s) who wins, and how much
	p.Round = 4 //recognize play is ended
	pot, err := p.GetPot()
	if err != nil {
		return err
	}

	for _, winnerId := range winnerIds {
		gains := int(pot / (len(winnerIds)))
		p.Players[winnerId].Stack += gains
		log.GLogf("Player %v wins $%v, GG !", p.Players[winnerId].Name, gains)
	}

	return nil
}

//Run every time a player answers to evaluate if round can end (by win, or go to next round)
func (p *Play) EndRound() error {
	//Players still playing needs to have all the same sum of bets && at least 2 bets for blinds player
	var playersStillPlaying []int

	for i := range p.Players {
		if p.Players[i].StillPlays {
			playersStillPlaying = append(playersStillPlaying, p.Players[i].ID)
		}
	}
	if len(playersStillPlaying) == 1 {
		return p.WinRound([]int{playersStillPlaying[0]})
	}

	log.Debugf("playersStillPlaying: %v", playersStillPlaying)

	sumBetAmountByPlayerId := make(map[int]int)
	sumBetByPlayerId := make(map[int]int) // could be used later for max 3 bet rounds, right now only to manage blind bets

	for _, bet := range (*p.Bets)[p.Round] {
		sumBetByPlayerId[bet.PlayerID] += 1
		sumBetAmountByPlayerId[bet.PlayerID] += bet.Amount
	}
	log.Debugf("sumBetByPlayerId: %v", sumBetByPlayerId)

	for _, id := range playersStillPlaying {
		canCheck, err := p.CanCheck(id)
		if err != nil {
			return err
		}
		if !canCheck {
			return log.Error("Not all still playing players have bet the same")
		}
		haveTalked := false
		//manages blinds
		if p.Round == 0 {
			//No need to manage heads up, blind player ids are always <= 1
			if (id <= 1 && sumBetByPlayerId[id] > 1) || (id > 1 && sumBetByPlayerId[id] > 0) {
				haveTalked = true
			}
		} else if sumBetByPlayerId[id] > 0 {
			haveTalked = true
		}
		log.Debugf("sumBetByPlayerId: %v", sumBetByPlayerId)
		log.Debugf("sumBetByPlayerId[id]: %v", sumBetByPlayerId[id])
		log.Debugf("haveTalked: %v", haveTalked)
		log.Debugf("id: %v", id)

		if !haveTalked {
			return log.Error("Not all still playing players express themselves #N.W.A.")
		}
	}

	log.Debug("We can go to next round")

	if p.Round == 3 {
		//To move into Showdown func ?
		playerScoreCards := make(map[int]*scores.ScoreCard, len(playersStillPlaying))

		for _, id := range playersStillPlaying {
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

			log.GLogf("%v had: %v", p.Players[id].Name, *playerCards)
		}

		var highestScoreCard *scores.ScoreCard = nil
		var winnerIds []int

		//need to browse map keys, not index
		for playerId := range playerScoreCards {
			if highestScoreCard == nil {
				highestScoreCard = playerScoreCards[playerId]
				winnerIds = []int{playerId}
				continue
			}
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

		for _, id := range winnerIds {
			log.GLogf("%v WINS with %v", p.Players[id].Name, highestScoreCard)
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
	potValue, err := p.GetPot()
	if err != nil {
		return err
	}
	log.GLogf("New pot value : %v$", potValue)

	return nil
}

func (p *Play) GetMinBetToCall(playerId int) (int, error) {
	//log.Debugf("Player %v tries to check, checking if she can", playerId)
	sumBetAmountByPlayerId := make(map[int]int)

	for _, bet := range (*p.Bets)[p.Round] {
		sumBetAmountByPlayerId[bet.PlayerID] += bet.Amount
	}
	log.Debugf("sumBetAmountByPlayerId at GetMinBetToCall %v", sumBetAmountByPlayerId)

	maxBet := 0
	for id := range sumBetAmountByPlayerId {
		if id == playerId {
			continue
		}
		if sumBetAmountByPlayerId[id] > maxBet {
			maxBet = sumBetAmountByPlayerId[id]
		}
	}

	minBetToCall := maxBet - sumBetAmountByPlayerId[playerId]

	log.Debugf("Player%v needs to bet $%v to call", playerId, minBetToCall)

	return minBetToCall, nil
}

func (p *Play) CanCheck(playerId int) (bool, error) {
	minBet, err := p.GetMinBetToCall(playerId)
	if err != nil {
		return false, err
	}
	if minBet != 0 {
		log.Debugf("Cannot check, need at least a bet of %v$", minBet)
		return false, nil
	}
	return true, nil
}

//Call or check (call with 0$)
func (p *Play) Call(playerId int) error {
	minBet, err := p.GetMinBetToCall(playerId)
	if err != nil {
		return err
	}

	return p.PutBet(playerId, minBet)
}

func (p *Play) Raise(playerId int, raiseAmount int) error {
	if raiseAmount < p.Blinds.SmallBlind {
		return log.Errorf("cannot raise of amount %v, needs to be > small blind", raiseAmount)
	}
	minBet, err := p.GetMinBetToCall(playerId)
	if err != nil {
		return err
	}

	return p.PutBet(playerId, minBet+raiseAmount)
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

//Very dumb, basic and unique strat
func (p *Play) ComputePlayerDecision(playerId int) (string, error) {
	if playerId > len(p.Players) {
		return "", log.Error("Player not in this play")
	}
	if !p.Players[playerId].StillPlays {
		return "", log.Error("player does not play anymore")
	}

	decision := "check"

	canCheck, err := p.CanCheck(playerId)
	if err != nil {
		log.Fatalf("got error during CanCheck %w", err)
	}
	if !canCheck {
		decision = "fold"
	}

	//Default behavior is check/fold, managing other cases now

	var playerHand []cards.Card // board + player cards

	playerCards, err := p.GetPlayerCards(playerId)
	if err != nil {
		return "", err
	}
	board, err := p.GetBoard()
	if err != nil {
		return "", err
	}
	for i := range *board {
		playerHand = append(playerHand, (*board)[i])
	}
	for i := range *playerCards {
		playerHand = append(playerHand, (*playerCards)[i])
	}

	sc, err := scores.Score(playerHand)
	if err != nil {
		return "", err
	}

	bestScoreField := sc.GetBestScoreField()

	if p.Round > 0 {
		boardSc, err := scores.Score(*board)
		if err != nil {
			return "", err
		}
		boardBestScoreField := boardSc.GetBestScoreField()
		if bestScoreField == boardBestScoreField {
			log.Debug("No competitive advantage") // not true, but simplifying
			//could manage later the potential straight or flush
			return decision, nil
		}
	}

	//should inject more advanced calls to stats.go here
	switch bestScoreField {
	case "HighCard":
		var valueSum uint8
		for _, card := range *playerCards {
			valueSum += card.Value
		}
		if valueSum > 15 {
			//A2 -> 78 hole cards
			decision = "call"
		}
	case "Pair":
		pairRank := playerHand[sc.Pair[0]].Value
		for _, card := range playerHand {
			if pairRank > card.Value {
				pairRank++
			}
		}
		if pairRank > uint8(p.Round) {
			//Arbitrary, always call on pre-flop pair
			//then if amongst the lowest pair => check/fold
			decision = "call"
		}
	case "DoublePair":
		decision = "call"
	default:
		decision = "raise"
	}

	log.Debugf("Intermediate decision after checking scoreCard: %v", decision)

	if decision == "check" || decision == "fold" {
		//Maybe we're waiting for a miracle ?
		maxSuit := sc.GetMaxSuit()
		maxSraight := sc.GetMaxStraight()

		log.Debugf("maxSuit: %v and maxSraight: %v are on a boat... no wait", maxSuit, maxSraight)

		if p.Round == 1 {
			if maxSuit >= 3 || maxSraight >= 3 {
				decision = "call"
			}
		} else if p.Round < 3 && (maxSuit == 4 || maxSraight >= 4) {
			decision = "call"
		}
	}

	log.Debugf("Final decision after checking suit and straight chances: %v", decision)

	return decision, nil
}

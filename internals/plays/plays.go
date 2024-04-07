package plays

import (
	"goker/pkg/cards"
)

type Play struct {
	ID string // Unique identifier for the game
	//StartTime  time.Time        // Time when the game started
	//EndTime    time.Time        // Time when the game ended
	Blinds  BlindInfo   // Information about blinds
	Players []Player    // ordered
	Bets    []Bet       // History of bets made during the game
	Board   cards.Board // Community cards on the board
	Result  PlayResult  // Result of the game
}

// Contains who wins, and what
type PlayResult struct {
	Winner  string         // ID of the winning player
	Payouts map[string]int // Payouts to each player
}

type Bet struct {
	PlayerID string // ID of the player making the bet
	Amount   int    // Amount of the bet
	//Timestamp time.Time // Time when the bet was made
}

// BlindInfo represents information about blinds.
type BlindInfo struct {
	SmallBlind   int // Small blind amount
	BigBlind     int // Big blind amount
	CurrentRound int // Current round of blinds
}

//ComputePot()

func NewGame() *Play {
	return &Play{
		//ID:        string, // Example function to generate a unique ID
		//StartTime:  time.Now(),
		Blinds:  bets.BlindInfo{SmallBlind: 5, BigBlind: 10, CurrentRound: 1},
		Players: make([]Player, 0),
		Bet:     make([]Bet, 0),
		Result:  PlayResult{},
	}
}

func (p *Play) AddPlayer() {
	err := NewPlayer()
	if err != nil {

	}
	p.Players = append(p.Players, player)
}

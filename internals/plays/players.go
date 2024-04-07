package plays

import (
	"goker/pkg/cards"
)

// Player represents a player in the game.
type Player struct {
	ID         string     // Unique identifier for the player
	Name       string     // Player's name
	Stack      int        // Player's stack (chips)
	Hand       cards.Hand // Player's hole cards
	LastAction string     // Last action taken by the player (e.g., bet, fold)
	Strategy   int
}

//Call()
//Fold()
//Raise()

func (p *Play) Bet(bet Bet) {
	p.Bets = append(p.Bets, bet)
}

func NewPlayer() (*Player, error) {
	return &Player{
		ID:         "X",
		Name:       "Skrib",
		Stack:      10000,
		Hand:       nil,
		LastAction: "",
		Strategy:   0,
	}, nil
}

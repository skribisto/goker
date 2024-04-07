package cards

import (
	"errors"
	"math/rand"
	"strconv"
)

type Card struct {
	Value uint8 // 2 to 14
	Suit  uint8 // 0 to 3, resp. clubs diamonds hearts spades
}

type Hand []Card  // 2 cards for players
type Deck []Card  // all 52 cards, randomly ordered
type Board []Card // flop + turn + river (+player hands for score)

func (c Card) String() string {
	var stringSuit string
	var stringVal string

	switch c.Value {
	case 11:
		stringVal = "V"
	case 12:
		stringVal = "D"
	case 13:
		stringVal = "R"
	case 14:
		stringVal = "A"
	default:
		stringVal = strconv.Itoa(int((c.Value)))
	}

	switch c.Suit {
	case 1:
		stringSuit = "♣"
	case 2:
		stringSuit = "♦"
	case 3:
		stringSuit = "♥"
	case 0:
		stringSuit = "♠"
	default:
		stringSuit = "XX"
	}
	return stringVal + stringSuit
}

func NewDeck() (*Deck, error) {
	var val uint8
	var suit uint8
	d := new(Deck)
	i := 0

	for suit = 0; suit < 4; suit++ {
		for val = 2; val <= 14; val++ {
			c := Card{
				Value: val,
				Suit:  suit,
			}
			*d = append(*d, c)
			i++
		}
	}
	//Always shuffle ?
	//log.Print("Before shuffle", d)
	d.Shuffle()
	return d, nil
}

func (d *Deck) Shuffle() {
	//rand.Seed(time.Now().UnixNano())
	rand.Seed(1)
	rand.Shuffle(len(*d), func(i, j int) { (*d)[i], (*d)[j] = (*d)[j], (*d)[i] })
}

func (d *Deck) Deal() (Card, error) {
	if len(*d) == 0 {
		return Card{}, errors.New("out of cards") // Return an empty card and false if the deck is empty
	}
	// Get the next card to be dealt
	card := (*d)[0]

	// Remove the dealt card from the deck
	*d = (*d)[1:]

	return card, nil
}

func (d *Deck) DealHand() (*Hand, error) {
	hand := new(Hand)

	for i := 0; i < 2; i++ {
		card, err := d.Deal()
		if err != nil {
			return nil, err
		}
		*hand = append(*hand, card)
	}

	return hand, nil
}

func (d *Deck) DealFlop() (*Board, error) {
	board := new(Board)

	for i := 0; i < 3; i++ {
		card, err := d.Deal()
		if err != nil {
			return nil, err
		}
		*board = append(*board, card)
	}

	return board, nil
}

func (d *Deck) DealTurnOrRiver(board *Board) error {
	if len(*board) != 3 || len(*board) != 4 {
		return errors.New("now correct amount of card in board")
	}
	card, err := d.Deal()
	if err != nil {
		return err
	}

	*board = append(*board, card)

	return nil
}

/*
func (h *Hand) String() string {
	s := make([]string, len(h.Cards))
	for i, card := range h.Cards {
		s[i] = card.String()
	}
	return strings.Join(s, " ")
}
*/

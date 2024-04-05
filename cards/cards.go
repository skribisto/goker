package cards

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

//const SUITS = [4]string{"spades", "hearts", "clubs", "diamonds"}

type Card struct {
	Value uint8 // 2 to 14
	Suit  uint8 // 0 to 3, resp. clubs diamonds hearts spades
}

type Deck []Card

type Board []Card

type Hand []Card //2 for texas hold'em, but can be different, e.g. Omaha

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
	return string(stringVal + stringSuit)
}

func (d *Deck) New() Deck {
	var val uint8
	var suit uint8
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
	return *d
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
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

/*
func (h *Hand) String() string {
	s := make([]string, len(h.Cards))
	for i, card := range h.Cards {
		s[i] = card.String()
		//s[i] = fmt.Sprint(card)
	}
	return strings.Join(s, " ")
}
*/

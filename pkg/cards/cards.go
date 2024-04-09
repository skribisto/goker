package cards

import (
	"math/rand"
	"strconv"
	"time"
)

type Card struct {
	Value uint8 // 2 to 14
	Suit  uint8 // 0 to 3, resp. clubs diamonds hearts spades
}

type Deck []Card // all 52 cards, randomly ordered

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
	rand.Seed(time.Now().UnixNano())
	//rand.Seed(1)
	rand.Shuffle(len(*d), func(i, j int) { (*d)[i], (*d)[j] = (*d)[j], (*d)[i] })
}

package cards

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Card struct {
	id int
    suit string
}

//TODO implement ascii art
//TODO manage suits
func (c Card) String() string {
	switch c.id {
	case 11:
		return "V"
	case 12:
		return "D"
	case 13:
		return "R"
	case 1:
		return "A"
	default:
		return strconv.Itoa(c.id)
	}
}

type Hand struct {
	cards []Card
}

func (h *Hand) String() string {
	s := make([]string, len(h.cards))
	for i, card := range h.cards {
		s[i] = card.String()
		//s[i] = fmt.Sprint(card)
	}
	return strings.Join(s, " ")
}

//Adapt to non-blackjack rules
func (h *Hand) Score() (totalScore int) {
	numberOfAces := 0
	for _, card := range h.cards {
		var cardScore int
		if card.id == 1 {
			cardScore = 11
			numberOfAces++
		} else if card.id < 10 {
			cardScore = card.id
		} else {
			cardScore = 10
		}
		totalScore += cardScore
	}

	for numberOfAces > 0 && totalScore > 21 {
		totalScore -= 10 //In case of multiple Aces, consider the minimal amount as ones
		numberOfAces--
	}

	return
}

//TODO create deck constraint (deal out of a deck / group of deck
func (h *Hand) DealRandom(n int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		randomId := rand.Intn(12)
		randomId++
		c := Card{randomId,""}
		h.cards = append(h.cards, c)
	}
}

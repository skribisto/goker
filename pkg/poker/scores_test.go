package scores

import (
	"goker/pkg/cards"
	"goker/pkg/log"
	"testing"
)

type scoreTests struct {
	name              string
	board             []cards.Card
	expectedScoreCard ScoreCard
}

var testBoards = []scoreTests{
	{
		name: "StraightFlush1",
		board: []cards.Card{
			{Value: 14, Suit: 2},
			{Value: 14, Suit: 0},
			{Value: 14, Suit: 1},
			{Value: 13, Suit: 1},
			{Value: 12, Suit: 1},
			{Value: 11, Suit: 1},
			{Value: 10, Suit: 1},
		},
		expectedScoreCard: ScoreCard{
			StraightFlush: true,
			Flush:         []uint8{2, 3, 4, 5, 6},
			Straight:      []uint8{2, 3, 4, 5, 6},
		},
	}, {
		name: "Four1",
		board: []cards.Card{
			{Value: 14, Suit: 2},
			{Value: 14, Suit: 0},
			{Value: 14, Suit: 1},
			{Value: 14, Suit: 3},
			{Value: 12, Suit: 1},
			{Value: 11, Suit: 1},
			{Value: 10, Suit: 1},
		},
		expectedScoreCard: ScoreCard{
			Four:     []uint8{0, 1, 2, 3},
			HighCard: []uint8{4, 5, 6},
		},
	}, {
		name: "FullHouse1",
		board: []cards.Card{
			{Value: 14, Suit: 0},
			{Value: 14, Suit: 1},
			{Value: 14, Suit: 2},
			{Value: 12, Suit: 3},
			{Value: 11, Suit: 3},
			{Value: 8, Suit: 2},
			{Value: 8, Suit: 0},
		},
		expectedScoreCard: ScoreCard{
			FullHouse: true,
			Three:     []uint8{0, 1, 2},
			Pair:      []uint8{5, 6},
		},
	}, {
		name: "FullHouse2",
		board: []cards.Card{
			{Value: 14, Suit: 0},
			{Value: 14, Suit: 1},
			{Value: 14, Suit: 2},
			{Value: 8, Suit: 3},
			{Value: 8, Suit: 2},
			{Value: 8, Suit: 0},
			{Value: 2, Suit: 3},
		},
		expectedScoreCard: ScoreCard{
			FullHouse: true,
			Three:     []uint8{0, 1, 2},
			Pair:      []uint8{3, 4},
		},
	}, {
		name: "FullHouse3",
		board: []cards.Card{
			{Value: 14, Suit: 0},
			{Value: 14, Suit: 1},
			{Value: 14, Suit: 2},
			{Value: 12, Suit: 3},
			{Value: 8, Suit: 3},
			{Value: 8, Suit: 2},
			{Value: 8, Suit: 0},
		},
		expectedScoreCard: ScoreCard{
			FullHouse: true,
			Three:     []uint8{0, 1, 2},
			Pair:      []uint8{4, 5},
		},
	}, {
		name: "FullHouse4",
		board: []cards.Card{
			{Value: 12, Suit: 3},
			{Value: 11, Suit: 0},
			{Value: 11, Suit: 1},
			{Value: 11, Suit: 2},
			{Value: 8, Suit: 3},
			{Value: 8, Suit: 2},
			{Value: 8, Suit: 0},
		},
		expectedScoreCard: ScoreCard{
			FullHouse: true,
			Three:     []uint8{1, 2, 3},
			Pair:      []uint8{4, 5},
		},
	}, {
		name: "FullHouse5",
		board: []cards.Card{
			{Value: 12, Suit: 3},
			{Value: 11, Suit: 0},
			{Value: 11, Suit: 1},
			{Value: 11, Suit: 2},
			{Value: 9, Suit: 3},
			{Value: 8, Suit: 2},
			{Value: 8, Suit: 0},
		},
		expectedScoreCard: ScoreCard{
			FullHouse: true,
			Three:     []uint8{1, 2, 3},
			Pair:      []uint8{5, 6},
		},
	}, {
		name: "Flush1",
		board: []cards.Card{
			{Value: 14, Suit: 0},
			{Value: 14, Suit: 1},
			{Value: 14, Suit: 2},
			{Value: 12, Suit: 1},
			{Value: 11, Suit: 1},
			{Value: 10, Suit: 1},
			{Value: 9, Suit: 1},
		},
		expectedScoreCard: ScoreCard{
			Flush: []uint8{1, 3, 4, 5, 6},
		},
	}, {
		name: "Straight1",
		board: []cards.Card{
			{Value: 14, Suit: 0},
			{Value: 14, Suit: 1},
			{Value: 14, Suit: 2},
			{Value: 13, Suit: 1},
			{Value: 12, Suit: 3},
			{Value: 11, Suit: 3},
			{Value: 10, Suit: 0},
		},
		expectedScoreCard: ScoreCard{
			Straight: []uint8{2, 3, 4, 5, 6},
		},
	}, {
		name: "DoublePair1",
		board: []cards.Card{
			{Value: 14, Suit: 0},
			{Value: 14, Suit: 1},
			{Value: 13, Suit: 1},
			{Value: 12, Suit: 3},
			{Value: 11, Suit: 3},
			{Value: 8, Suit: 2},
			{Value: 8, Suit: 0},
		},
		expectedScoreCard: ScoreCard{
			DoublePair: []uint8{5, 6},
			Pair:       []uint8{0, 1},
			HighCard:   []uint8{2, 3, 4},
		},
	}, {
		name: "DoublePair2",
		board: []cards.Card{
			{Value: 14, Suit: 0},
			{Value: 13, Suit: 0},
			{Value: 13, Suit: 1},
			{Value: 12, Suit: 3},
			{Value: 11, Suit: 3},
			{Value: 8, Suit: 2},
			{Value: 8, Suit: 0},
		},
		expectedScoreCard: ScoreCard{
			DoublePair: []uint8{5, 6},
			Pair:       []uint8{1, 2},
			HighCard:   []uint8{0, 3, 4},
		},
	}, {
		name: "Pair1",
		board: []cards.Card{
			{Value: 14, Suit: 0},
			{Value: 14, Suit: 1},
			{Value: 13, Suit: 1},
			{Value: 12, Suit: 3},
			{Value: 11, Suit: 3},
			{Value: 9, Suit: 2},
			{Value: 8, Suit: 0},
		},
		expectedScoreCard: ScoreCard{
			HighCard: []uint8{3, 4, 5, 6, 2},
			Pair:     []uint8{0, 1},
		},
	}, {
		name: "HighCard1",
		board: []cards.Card{
			{Value: 14, Suit: 0},
			{Value: 13, Suit: 1},
			{Value: 12, Suit: 3},
			{Value: 11, Suit: 3},
			{Value: 9, Suit: 2},
			{Value: 8, Suit: 0},
			{Value: 6, Suit: 1},
		},
		expectedScoreCard: ScoreCard{
			HighCard: []uint8{0, 1, 2, 3, 4, 5, 6},
		},
	},
}

func TestScore(t *testing.T) {
	// {
	// 	Board: &board,
	// 	StraightFlush: false,
	// 	Four: nil,
	// 	FullHouse: false,
	// 	Flush: nil,
	// 	Straight: nil,
	// 	Three: nil,
	// 	DoublePair: nil,
	// 	Pair: nil,
	// 	HighCard: nil,
	// }

	for i := range testBoards {
		sc, err := Score(testBoards[i].board)
		if err != nil {
			t.Fatalf("Test-%v could not Score the board: %v", testBoards[i].name, err.Error())
		}
		sc.Board = nil
		log.Debugf("Comparing %v with %v", sc, testBoards[i].expectedScoreCard)

		//Surely there's a better way...
		if sc.StraightFlush != testBoards[i].expectedScoreCard.StraightFlush {
			t.Fatalf("Test-%v ScoreCards StraightFlush are different, got: %v, expected: %v", testBoards[i].name, sc.StraightFlush, testBoards[i].expectedScoreCard.StraightFlush)
		} else if len(sc.Four) != len(testBoards[i].expectedScoreCard.Four) {
			t.Fatalf("Test-%v ScoreCards Four are different, got: %v, expected: %v", testBoards[i].name, sc.Four, testBoards[i].expectedScoreCard.Four)
		} else if len(sc.Four) > 0 {
			for j := range sc.Four {
				if testBoards[i].expectedScoreCard.Four[j] != sc.Four[j] {
					t.Fatalf("Test-%v ScoreCards Four are different, got: %v, expected: %v", testBoards[i].name, sc.Four, testBoards[i].expectedScoreCard.Four)
				}
			}
		}
		if sc.FullHouse != testBoards[i].expectedScoreCard.FullHouse {
			t.Fatalf("Test-%v ScoreCards FullHouse are different, got: %v, expected: %v", testBoards[i].name, sc.FullHouse, testBoards[i].expectedScoreCard.FullHouse)
		} else if len(sc.Flush) != len(testBoards[i].expectedScoreCard.Flush) {
			t.Fatalf("Test-%v ScoreCards Flush are different, got: %v, expected: %v", testBoards[i].name, sc.Flush, testBoards[i].expectedScoreCard.Flush)
		} else if len(sc.Flush) > 0 {
			for j := range sc.Flush {
				if testBoards[i].expectedScoreCard.Flush[j] != sc.Flush[j] {
					t.Fatalf("Test-%v ScoreCards Flush are different, got: %v, expected: %v", testBoards[i].name, sc.Flush, testBoards[i].expectedScoreCard.Flush)
				}
			}
		} else if len(sc.Straight) != len(testBoards[i].expectedScoreCard.Straight) {
			t.Fatalf("Test-%v ScoreCards Straight are different, got: %v, expected: %v", testBoards[i].name, sc.Straight, testBoards[i].expectedScoreCard.Straight)
		} else if len(sc.Straight) > 0 {
			for j := range sc.Straight {
				if testBoards[i].expectedScoreCard.Straight[j] != sc.Straight[j] {
					t.Fatalf("Test-%v ScoreCards Straight are different, got: %v, expected: %v", testBoards[i].name, sc.Straight, testBoards[i].expectedScoreCard.Straight)
				}
			}
		} else if len(sc.Three) != len(testBoards[i].expectedScoreCard.Three) {
			t.Fatalf("Test-%v ScoreCards Three are different, got: %v, expected: %v", testBoards[i].name, sc.Three, testBoards[i].expectedScoreCard.Three)
		} else if len(sc.Three) > 0 {
			for j := range sc.Three {
				if testBoards[i].expectedScoreCard.Three[j] != sc.Three[j] {
					t.Fatalf("Test-%v ScoreCards Three are different, got: %v, expected: %v", testBoards[i].name, sc.Three, testBoards[i].expectedScoreCard.Three)
				}
			}
		} else if len(sc.DoublePair) != len(testBoards[i].expectedScoreCard.DoublePair) {
			t.Fatalf("Test-%v ScoreCards DoublePair are different, got: %v, expected: %v", testBoards[i].name, sc.DoublePair, testBoards[i].expectedScoreCard.DoublePair)
		} else if len(sc.DoublePair) > 0 {
			for j := range sc.DoublePair {
				if testBoards[i].expectedScoreCard.DoublePair[j] != sc.DoublePair[j] {
					t.Fatalf("Test-%v ScoreCards DoublePair are different, got: %v, expected: %v", testBoards[i].name, sc.DoublePair, testBoards[i].expectedScoreCard.DoublePair)
				}
			}
		} else if len(sc.Pair) != len(testBoards[i].expectedScoreCard.Pair) {
			t.Fatalf("Test-%v ScoreCards Pair are different, got: %v, expected: %v", testBoards[i].name, sc.Pair, testBoards[i].expectedScoreCard.Pair)
		} else if len(sc.Pair) > 0 {
			for j := range sc.Pair {
				if testBoards[i].expectedScoreCard.Pair[j] != sc.Pair[j] {
					t.Fatalf("Test-%v ScoreCards Pair are different, got: %v, expected: %v", testBoards[i].name, sc.Pair, testBoards[i].expectedScoreCard.Pair)
				}
			}
		} else if len(sc.HighCard) != len(testBoards[i].expectedScoreCard.HighCard) {
			t.Fatalf("Test-%v ScoreCards HighCard are different, got: %v, expected: %v", testBoards[i].name, sc, testBoards[i].expectedScoreCard)
		} else if len(sc.HighCard) > 0 {
			for j := range sc.HighCard {
				if testBoards[i].expectedScoreCard.HighCard[j] != sc.HighCard[j] {
					t.Fatalf("Test-%v ScoreCards HighCard are different, got: %v, expected: %v", testBoards[i].name, sc.HighCard, testBoards[i].expectedScoreCard.HighCard)
				}
			}
		}
	}

}

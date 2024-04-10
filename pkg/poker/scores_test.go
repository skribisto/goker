package scores

import (
	"goker/pkg/cards"
	"testing"
)

type scoreTests struct {
	board             []cards.Card
	expectedScoreCard ScoreCard
}

var scoreTestStraightFlush1 = scoreTests{
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
		HighCard:      []uint8{0, 1},
		Three:         []uint8{0, 1, 2},
	},
}
var scoreTestFlush1 = scoreTests{
	board: []cards.Card{
		{Value: 14, Suit: 0},
		{Value: 14, Suit: 1},
		{Value: 14, Suit: 2},
		{Value: 13, Suit: 1},
		{Value: 12, Suit: 1},
		{Value: 11, Suit: 1},
		{Value: 10, Suit: 1},
	},
	expectedScoreCard: ScoreCard{
		Flush:    []uint8{2, 3, 4, 5, 6},
		HighCard: []uint8{3, 4, 5, 6},
		Three:    []uint8{0, 1, 2},
	},
}

var scoreTestStraight1 = scoreTests{
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
		Flush:    []uint8{2, 3, 4, 5, 6},
		HighCard: []uint8{3, 4, 5, 6},
		Three:    []uint8{0, 1, 2},
	},
}
var scoreTestPair1 = scoreTests{
	board: []cards.Card{
		{Value: 14, Suit: 0},
		{Value: 14, Suit: 1},
		{Value: 9, Suit: 2},
		{Value: 13, Suit: 1},
		{Value: 12, Suit: 3},
		{Value: 11, Suit: 3},
		{Value: 10, Suit: 0},
	},
	expectedScoreCard: ScoreCard{
		HighCard: []uint8{3, 4, 5, 6, 2},
		Pair:     []uint8{0, 1},
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

	var tests []scoreTests

	tests = append(tests, scoreTestStraightFlush1)
	tests = append(tests, scoreTestFlush1)
	tests = append(tests, scoreTestStraight1)

	for i := range tests {
		sc, err := Score(tests[i].board)
		if err != nil {
			t.Errorf("could not Score the board: %v", err.Error())
		}
		sc.Board = nil

		//Surely there's a better way...
		if sc.StraightFlush != tests[i].expectedScoreCard.StraightFlush {
			t.Errorf("Test%v ScoreCards StraightFlush are different, got: %v, expected: %v", i, sc, tests[i].expectedScoreCard)
		} else if len(sc.Four) != len(tests[i].expectedScoreCard.Four) {
			t.Errorf("Test%v ScoreCards Four are different, got: %v, expected: %v", i, sc, tests[i].expectedScoreCard)
		} else if sc.FullHouse != tests[i].expectedScoreCard.FullHouse {
			t.Errorf("Test%v ScoreCards FullHouse are different, got: %v, expected: %v", i, sc, tests[i].expectedScoreCard)
		} else if len(sc.Flush) != len(tests[i].expectedScoreCard.Flush) {
			t.Errorf("Test%v ScoreCards Flush are different, got: %v, expected: %v", i, sc, tests[i].expectedScoreCard)
		} else if len(sc.Straight) != len(tests[i].expectedScoreCard.Straight) {
			t.Errorf("Test%v ScoreCards Straight are different, got: %v, expected: %v", i, sc, tests[i].expectedScoreCard)
		} else if len(sc.Three) != len(tests[i].expectedScoreCard.Three) {
			t.Errorf("Test%v ScoreCards Three are different, got: %v, expected: %v", i, sc, tests[i].expectedScoreCard)
		} else if len(sc.DoublePair) != len(tests[i].expectedScoreCard.DoublePair) {
			t.Errorf("Test%v ScoreCards DoublePair are different, got: %v, expected: %v", i, sc, tests[i].expectedScoreCard)
		} else if len(sc.Pair) != len(tests[i].expectedScoreCard.Pair) {
			t.Errorf("Test%v ScoreCards Pair are different, got: %v, expected: %v", i, sc, tests[i].expectedScoreCard)
		} else if len(sc.HighCard) != len(tests[i].expectedScoreCard.HighCard) {
			t.Errorf("Test%v ScoreCards HighCard are different, got: %v, expected: %v", i, sc, tests[i].expectedScoreCard)
		}
	}

}

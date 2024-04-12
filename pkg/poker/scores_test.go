package scores

import (
	"goker/pkg/cards"
	"testing"
)

type scoreTests struct {
	board             []cards.Card
	expectedScoreCard ScoreCard
}

//boards needs to be ordered
var testBoards = map[string]scoreTests{
	"StraightFlush1": {
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
	}, "Four1": {
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
	}, "FullHouse1": {
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
	}, "FullHouse2": {
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
	}, "FullHouse3": {
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
	}, "FullHouse4": {
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
	}, "FullHouse5": {
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
	}, "Flush1": {
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
	}, "Flush2": {
		board: []cards.Card{
			{Value: 14, Suit: 0},
			{Value: 12, Suit: 1},
			{Value: 11, Suit: 1},
			{Value: 10, Suit: 1},
			{Value: 9, Suit: 1},
			{Value: 4, Suit: 1},
			{Value: 2, Suit: 1},
		},
		expectedScoreCard: ScoreCard{
			Flush: []uint8{1, 2, 3, 4, 5},
		},
	}, "Flush3": {
		board: []cards.Card{
			{Value: 14, Suit: 1},
			{Value: 13, Suit: 1},
			{Value: 12, Suit: 1},
			{Value: 11, Suit: 1},
			{Value: 9, Suit: 1},
			{Value: 8, Suit: 1},
			{Value: 7, Suit: 1},
		},
		expectedScoreCard: ScoreCard{
			Flush: []uint8{0, 1, 2, 3, 4},
		},
	}, "Straight1": {
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
			Straight: []uint8{0, 3, 4, 5, 6},
		},
	}, "Straight2": {
		board: []cards.Card{
			{Value: 12, Suit: 0},
			{Value: 11, Suit: 1},
			{Value: 10, Suit: 2},
			{Value: 10, Suit: 1},
			{Value: 9, Suit: 3},
			{Value: 8, Suit: 3},
			{Value: 3, Suit: 0},
		},
		expectedScoreCard: ScoreCard{
			Straight: []uint8{0, 1, 2, 4, 5},
		},
	}, "DoublePair1": {
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
	}, "Three1": {
		board: []cards.Card{
			{Value: 14, Suit: 0},
			{Value: 13, Suit: 0},
			{Value: 13, Suit: 1},
			{Value: 13, Suit: 3},
			{Value: 11, Suit: 3},
			{Value: 9, Suit: 2},
			{Value: 8, Suit: 0},
		},
		expectedScoreCard: ScoreCard{
			Three:    []uint8{1, 2, 3},
			HighCard: []uint8{0, 4, 5, 6},
		},
	}, "DoublePair2": {
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
	}, "DoublePair3": {
		board: []cards.Card{
			{Value: 12, Suit: 2},
			{Value: 12, Suit: 0},
			{Value: 10, Suit: 3},
			{Value: 8, Suit: 1},
			{Value: 5, Suit: 3},
			{Value: 5, Suit: 0},
			{Value: 2, Suit: 0},
		},
		expectedScoreCard: ScoreCard{
			DoublePair: []uint8{4, 5},
			Pair:       []uint8{0, 1},
			HighCard:   []uint8{2, 3, 6},
		},
	}, "Pair1": {
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
			Pair:     []uint8{0, 1},
			HighCard: []uint8{3, 4, 5, 6, 2},
		},
	}, "Pair2": {
		board: []cards.Card{
			{Value: 14, Suit: 0},
			{Value: 14, Suit: 1},
			{Value: 13, Suit: 1},
			{Value: 12, Suit: 3},
			{Value: 9, Suit: 3},
			{Value: 8, Suit: 2},
			{Value: 5, Suit: 0},
		},
		expectedScoreCard: ScoreCard{
			Pair:     []uint8{0, 1},
			HighCard: []uint8{2, 3, 4, 5, 6},
		},
	}, "Pair3": {
		board: []cards.Card{
			{Value: 14, Suit: 0},
			{Value: 14, Suit: 1},
			{Value: 13, Suit: 1},
			{Value: 12, Suit: 3},
			{Value: 10, Suit: 3},
			{Value: 8, Suit: 2},
			{Value: 5, Suit: 0},
		},
		expectedScoreCard: ScoreCard{
			Pair:     []uint8{0, 1},
			HighCard: []uint8{2, 3, 4, 5, 6},
		},
	}, "HighCard1": {
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
			HighCard: []uint8{0, 1, 2, 3, 4},
		},
	}, "PartialBoard1": {
		board: []cards.Card{
			{Value: 14, Suit: 0},
			{Value: 11, Suit: 3},
			{Value: 9, Suit: 2},
			{Value: 6, Suit: 0},
			{Value: 6, Suit: 1},
		},
		expectedScoreCard: ScoreCard{
			HighCard: []uint8{0, 1, 2},
			Pair:     []uint8{3, 4},
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

	for test, st := range testBoards {
		sc, err := Score(st.board)
		if err != nil {
			t.Fatalf("Test-%v could not Score the board: %v", test, err.Error())
		}
		sc.Board = nil
		t.Logf("Testing: %v", test)
		//t.Logf("Comparing %#v with %#v", sc, st.expectedScoreCard)

		//Surely there's a better way...
		if sc.StraightFlush != st.expectedScoreCard.StraightFlush {
			t.Errorf("Test-%v ScoreCards StraightFlush are different, got: %v, expected: %v", test, sc.StraightFlush, st.expectedScoreCard.StraightFlush)
		} else if len(sc.Four) != len(st.expectedScoreCard.Four) {
			t.Errorf("Test-%v ScoreCards Four are different, got: %v, expected: %v", test, sc.Four, st.expectedScoreCard.Four)
		} else if len(sc.Four) > 0 {
			for j := range sc.Four {
				if st.expectedScoreCard.Four[j] != sc.Four[j] {
					t.Errorf("Test-%v ScoreCards Four are different, got: %v, expected: %v", test, sc.Four, st.expectedScoreCard.Four)
				}
			}
		} else if sc.FullHouse != st.expectedScoreCard.FullHouse {
			t.Errorf("Test-%v ScoreCards FullHouse are different, got: %v, expected: %v", test, sc.FullHouse, st.expectedScoreCard.FullHouse)
		} else if len(sc.Flush) != len(st.expectedScoreCard.Flush) {
			t.Errorf("Test-%v ScoreCards Flush are different, got: %v, expected: %v", test, sc.Flush, st.expectedScoreCard.Flush)
		} else if len(sc.Flush) > 0 {
			for j := range sc.Flush {
				if st.expectedScoreCard.Flush[j] != sc.Flush[j] {
					t.Errorf("Test-%v ScoreCards Flush are different, got: %v, expected: %v", test, sc.Flush, st.expectedScoreCard.Flush)
				}
			}
		} else if len(sc.Straight) != len(st.expectedScoreCard.Straight) {
			t.Errorf("Test-%v ScoreCards Straight are different, got: %v, expected: %v", test, sc.Straight, st.expectedScoreCard.Straight)
		} else if len(sc.Straight) > 0 {
			for j := range sc.Straight {
				if st.expectedScoreCard.Straight[j] != sc.Straight[j] {
					t.Errorf("Test-%v ScoreCards Straight are different, got: %v, expected: %v", test, sc.Straight, st.expectedScoreCard.Straight)
				}
			}
		} else if len(sc.Three) != len(st.expectedScoreCard.Three) {
			t.Errorf("Test-%v ScoreCards Three are different, got: %v, expected: %v", test, sc.Three, st.expectedScoreCard.Three)
		} else if len(sc.Three) > 0 {
			for j := range sc.Three {
				if st.expectedScoreCard.Three[j] != sc.Three[j] {
					t.Errorf("Test-%v ScoreCards Three are different, got: %v, expected: %v", test, sc.Three, st.expectedScoreCard.Three)
				}
			}
		} else if len(sc.DoublePair) != len(st.expectedScoreCard.DoublePair) {
			t.Errorf("Test-%v ScoreCards DoublePair are different, got: %v, expected: %v", test, sc.DoublePair, st.expectedScoreCard.DoublePair)
		} else if len(sc.DoublePair) > 0 {
			for j := range sc.DoublePair {
				if st.expectedScoreCard.DoublePair[j] != sc.DoublePair[j] {
					t.Errorf("Test-%v ScoreCards DoublePair are different, got: %v, expected: %v", test, sc.DoublePair, st.expectedScoreCard.DoublePair)
				}
			}
		} else if len(sc.Pair) != len(st.expectedScoreCard.Pair) {
			t.Errorf("Test-%v ScoreCards Pair are different, got: %v, expected: %v", test, sc.Pair, st.expectedScoreCard.Pair)
		} else if len(sc.Pair) > 0 {
			for j := range sc.Pair {
				if st.expectedScoreCard.Pair[j] != sc.Pair[j] {
					t.Errorf("Test-%v ScoreCards Pair are different, got: %v, expected: %v", test, sc.Pair, st.expectedScoreCard.Pair)
				}
			}
		} else if len(sc.HighCard) != len(st.expectedScoreCard.HighCard) {
			t.Errorf("Test-%v ScoreCards HighCard are different, got: %v, expected: %v", test, sc, st.expectedScoreCard)
		} else if len(sc.HighCard) > 0 {
			for j := range sc.HighCard {
				if st.expectedScoreCard.HighCard[j] != sc.HighCard[j] {
					t.Errorf("Test-%v ScoreCards HighCard are different, got: %v, expected: %v", test, sc.HighCard, st.expectedScoreCard.HighCard)
				}
			}
		}
	}

}

func TestCompareScoreCards(t *testing.T) {
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

	var compareTests = map[string]struct {
		test1          string
		test2          string
		expectedResult int
	}{
		"SamePairDiffHighCards1": {
			test1:          "Pair2",
			test2:          "Pair3",
			expectedResult: -1,
		}, "Same1": {
			test1:          "Pair2",
			test2:          "Pair2",
			expectedResult: 0,
		}, "Same2": {
			test1:          "Pair1",
			test2:          "Pair1",
			expectedResult: 0,
		}, "Same3": {
			test1:          "Straight1",
			test2:          "Straight1",
			expectedResult: 0,
		}, "Same4": {
			test1:          "Straight2",
			test2:          "Straight2",
			expectedResult: 0,
		}, "Same5": {
			test1:          "StraightFlush1",
			test2:          "StraightFlush1",
			expectedResult: 0,
		}, "Same6": {
			test1:          "Flush1",
			test2:          "Flush1",
			expectedResult: 0,
		}, "Same7": {
			test1:          "Four1",
			test2:          "Four1",
			expectedResult: 0,
		}, "Same8": {
			test1:          "Three1",
			test2:          "Three1",
			expectedResult: 0,
		}, "Same9": {
			test1:          "DoublePair1",
			test2:          "DoublePair1",
			expectedResult: 0,
		}, "Same10": {
			test1:          "HighCard1",
			test2:          "HighCard1",
			expectedResult: 0,
		}, "SamePairDiffHighCards2": {
			test1:          "Pair3",
			test2:          "Pair2",
			expectedResult: 1,
		}, "SameFullHouse1": {
			test1:          "FullHouse1",
			test2:          "FullHouse2",
			expectedResult: 0,
		}, "SameFullHouse2": {
			test1:          "FullHouse2",
			test2:          "FullHouse3",
			expectedResult: 0,
		}, "SameFullHouse3": {
			test1:          "FullHouse1",
			test2:          "FullHouse3",
			expectedResult: 0,
		}, "SameFullHouse4": {
			test1:          "FullHouse4",
			test2:          "FullHouse5",
			expectedResult: 0,
		}, "DiffFullHouse1": {
			test1:          "FullHouse1",
			test2:          "FullHouse4",
			expectedResult: 1,
		}, "DiffFullHouse2": {
			test1:          "FullHouse1",
			test2:          "FullHouse5",
			expectedResult: 1,
		}, "StraightFlushOverStraight1": {
			test1:          "StraightFlush1",
			test2:          "Straight1",
			expectedResult: 1,
		}, "StraightFlushOverFlush1": {
			test1:          "StraightFlush1",
			test2:          "Flush1",
			expectedResult: 1,
		},
	}

	for testName, test := range compareTests {
		sc1 := testBoards[test.test1].expectedScoreCard
		sc2 := testBoards[test.test2].expectedScoreCard

		board1 := testBoards[test.test1].board
		board2 := testBoards[test.test2].board

		sc1.Board = &board1
		sc2.Board = &board2

		comparator, err := CompareScoreCards(&sc1, &sc2)
		if err != nil {
			t.Fatalf("Test-%v could not CompareScoreCards the expectedScoreCards: %v", testName, err.Error())
		}
		if comparator != test.expectedResult {
			t.Errorf("Test-%v, got %v but expected %v", testName, comparator, test.expectedResult)
		} else {
			t.Logf("Test-%v SUCCESS", testName)
		}
	}

}

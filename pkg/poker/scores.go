package scores

import (
	"sort"

	"github.com/skribisto/goker/pkg/cards"
	"github.com/skribisto/goker/pkg/log"
)

type ScoreCard struct {
	Board         *[]cards.Card
	StraightFlush bool
	Four          []uint8 // 4 of a kind, aka OAK4
	FullHouse     bool
	Flush         []uint8
	Straight      []uint8
	Three         []uint8 // aka OAK3
	DoublePair    []uint8 // aka doubleOAK2, contains lower pair
	Pair          []uint8 // aka OAK2, contains highest pair
	HighCard      []uint8 //may need up to 5 to discriminate
}

func (sc *ScoreCard) String() string {
	winCards := ""
	if sc.StraightFlush {
		for _, i := range sc.Straight {
			winCards += (*sc.Board)[i].String() + " "
		}
		return "Straight Flush:  " + winCards
	} else if len(sc.Four) != 0 {
		for _, i := range sc.Four {
			winCards += (*sc.Board)[i].String() + " "
		}
		return "Four of a kind: " + winCards + (*sc.Board)[sc.HighCard[0]].String()
	} else if sc.FullHouse {
		for _, i := range sc.Three {
			winCards += (*sc.Board)[i].String() + " "
		}
		for _, i := range sc.Pair {
			winCards += (*sc.Board)[i].String() + " "
		}
		return "Full House: " + winCards
	} else if len(sc.Flush) != 0 {
		for _, i := range sc.Flush {
			winCards += (*sc.Board)[i].String() + " "
		}
		return "Flush:  " + winCards
	} else if len(sc.Straight) != 0 {
		for _, i := range sc.Straight {
			winCards += (*sc.Board)[i].String() + " "
		}
		return "Straight:  " + winCards
	} else if len(sc.Three) != 0 {
		for _, i := range sc.Three {
			winCards += (*sc.Board)[i].String() + " "
		}
		return "Three of a kind: " + winCards + (*sc.Board)[sc.HighCard[0]].String() + " " + (*sc.Board)[sc.HighCard[1]].String()
	} else if len(sc.DoublePair) != 0 {
		for _, i := range sc.Pair {
			winCards += (*sc.Board)[i].String() + " "
		}
		for _, i := range sc.DoublePair {
			winCards += (*sc.Board)[i].String() + " "
		}
		return "Double Pair: " + winCards + (*sc.Board)[sc.HighCard[0]].String()
	} else if len(sc.Pair) > 0 {
		for _, i := range sc.Pair {
			winCards += (*sc.Board)[i].String() + " "
		}
		return "Pair: " + winCards + (*sc.Board)[sc.HighCard[0]].String() + " " + (*sc.Board)[sc.HighCard[1]].String() + " " + (*sc.Board)[sc.HighCard[2]].String()
	} else if len(sc.HighCard) > 0 {
		for _, i := range sc.HighCard {
			winCards += (*sc.Board)[i].String() + " "
		}
		return "High card: " + winCards
	}
	return winCards
}

func (sc *ScoreCard) GetBestScoreField() string {
	if sc.StraightFlush {
		return "StraightFlush"
	} else if len(sc.Four) != 0 {
		return "Four"
	} else if sc.FullHouse {
		return "FullHouse"
	} else if len(sc.Flush) != 0 {
		return "Flush"
	} else if len(sc.Straight) != 0 {
		return "Straight"
	} else if len(sc.Three) != 0 {
		return "Three"
	} else if len(sc.DoublePair) != 0 {
		return "DoublePair"
	} else if len(sc.Pair) > 0 {
		return "Pair"
	}
	return "HighCard"
}

func Score(b []cards.Card) (*ScoreCard, error) {
	//Create and fill a new ScoreCard
	sc := new(ScoreCard)

	// do not manipulate (order) the board, only the copy made for ScoreCard
	var board []cards.Card

	board = append(board, b...)

	// Make a copy of the board
	sc.Board = &board

	if len(board) > 7 || len(board) < 2 {
		return nil, log.Errorf("not enough (<2) or too much (>7) cards to evaluate score from: ,%v", board)
	}

	//sort by value desc
	sort.SliceStable(board, func(i, j int) bool {
		return board[i].Value > board[j].Value
	})

	err := sc.checkStraight()
	if err != nil {
		return nil, err
	}
	//log.Debugf("SC checkStraight GOT : %#v", sc)

	err = sc.checkFlush()
	if err != nil {
		return nil, err
	}
	//log.Debugf("SC checkFlush GOT : %#v", sc)

	if len(sc.Straight) > 0 && len(sc.Flush) > 0 {
		sc.StraightFlush = true
		//got our 5 cards
	} else if len(sc.Straight) > 0 || len(sc.Flush) > 0 {
		log.Debug("We have 5 cards, no need to go further")
	} else {
		err = sc.checkXOfAKind()
		if err != nil {
			return nil, err
		}
	}

	//log.Debugf("sc at the end of Score %#v", sc)
	log.Debugf("sc at the end of Score %v", sc)

	return sc, nil
}

//1 means sc1 wins, -1 means sc1 loose, 0 means tie
func CompareScoreCards(sc1, sc2 *ScoreCard) (int, error) {
	board1 := *sc1.Board
	board2 := *sc2.Board

	if sc1.StraightFlush {
		if len(sc1.Straight) != 5 {
			return 0, log.Error("flag StraightFlush but no 5 Straight value in scoreCard 1")
		}
		if sc2.StraightFlush {
			if len(sc2.Straight) != 5 {
				return 0, log.Error("flag StraightFlush but no 5 Straight value in scoreCard 2")
			}
			if board1[sc1.Straight[0]].Value < board2[sc2.Straight[0]].Value {
				return -1, nil
			} else if board1[sc1.Straight[0]].Value == board2[sc2.Straight[0]].Value {
				return 0, nil
			}
		}
		return 1, nil
	} else if len(sc1.Four) > 0 {
		if len(sc1.Four) != 4 {
			return 0, log.Error("no Four value in scoreCard 1")
		}
		if sc2.StraightFlush {
			return -1, nil
		}
		if len(sc2.Four) > 0 {
			if len(sc2.Four) != 4 {
				return 0, log.Error("no Four value in scoreCard 2")
			}
			if board1[sc1.Four[0]].Value < board2[sc2.Four[0]].Value {
				return -1, nil
			} else if board1[sc1.Four[0]].Value == board2[sc2.Four[0]].Value {
				if len(sc1.HighCard) < 1 || len(sc2.HighCard) < 1 {
					return 0, log.Error("no HighCard value in scoreCard 1 or 2")
				}
				if board1[sc1.HighCard[0]].Value < board2[sc2.HighCard[0]].Value {
					return -1, nil
				} else if board1[sc1.HighCard[0]].Value == board2[sc2.HighCard[0]].Value {
					return 0, nil
				}
			}
		}
		return 1, nil
	} else if sc1.FullHouse {
		if len(sc1.Three) != 3 || len(sc1.Pair) != 2 {
			return 0, log.Error("no 3 Three or 2 Pair value in scoreCard 1")
		}
		if sc2.StraightFlush || len(sc2.Four) > 0 {
			return -1, nil
		}
		if sc2.FullHouse {
			if len(sc2.Three) != 3 || len(sc2.Pair) != 2 {
				return 0, log.Error("no 3 Three or 2 Pair value in scoreCard 2")
			}
			if board1[sc1.Three[0]].Value < board2[sc2.Three[0]].Value {
				return -1, nil
			} else if board1[sc1.Three[0]].Value == board2[sc2.Three[0]].Value {
				if board1[sc1.Pair[0]].Value < board2[sc2.Pair[0]].Value {
					return -1, nil
				} else if board1[sc1.Pair[0]].Value == board2[sc2.Pair[0]].Value {
					return 0, nil
				}
			}
		}
		return 1, nil
	} else if len(sc1.Flush) > 0 {
		if len(sc1.Flush) != 5 {
			return 0, log.Error("no 5 Flush value in scoreCard 1")
		}
		if sc2.StraightFlush || len(sc2.Four) > 0 || sc2.FullHouse {
			return -1, nil
		}
		if len(sc2.Flush) > 0 {
			if len(sc2.Flush) != 5 {
				return 0, log.Error("no 5 Flush value in scoreCard 2")
			}
			for _, i := range sc1.Flush {
				//board is always sorted
				if board1[i].Value < board2[i].Value {
					return -1, nil
				} else if board1[i].Value > board2[i].Value {
					return 1, nil
				}
			}
			return 0, nil
		}
		return 1, nil
	} else if len(sc1.Straight) > 0 {
		if len(sc1.Straight) != 5 {
			return 0, log.Error("no 5 Straight value in scoreCard 1")
		}
		if sc2.StraightFlush || len(sc2.Four) > 0 || sc2.FullHouse || len(sc2.Flush) > 0 {
			return -1, nil
		}
		if len(sc2.Straight) > 0 {
			if len(sc2.Straight) != 5 {
				return 0, log.Error("no 5 Straight value in scoreCard 2")
			}
			if board1[sc1.Straight[0]].Value < board2[sc2.Straight[0]].Value {
				return -1, nil
			} else if board1[sc1.Straight[0]].Value == board2[sc2.Straight[0]].Value {
				//We know the rest is the same too as it's a straight
				return 0, nil
			}
		}
		return 1, nil
	} else if len(sc1.Three) > 0 {
		if len(sc1.Three) != 3 {
			return 0, log.Error("no Three value in scoreCard 1")
		}
		if sc2.StraightFlush || len(sc2.Four) > 0 || sc2.FullHouse || len(sc2.Flush) > 0 || len(sc2.Straight) > 0 {
			return -1, nil
		}
		if len(sc2.Three) > 0 {
			if len(sc2.Three) != 3 {
				return 0, log.Error("no Three value in scoreCard 2")
			}
			if board1[sc1.Three[0]].Value < board2[sc2.Three[0]].Value {
				return -1, nil
			} else if board1[sc1.Three[0]].Value == board2[sc2.Three[0]].Value {
				if len(sc1.HighCard) < 2 || len(sc2.HighCard) < 2 {
					return 0, log.Errorf("no (enough) HighCard value in scoreCard 1 %v or scoreCard 2: %v", sc1.HighCard, sc2.HighCard)
				}
				for _, i := range sc1.HighCard[:2] {
					if board1[i].Value < board2[i].Value {
						return -1, nil
					} else if board1[i].Value > board2[i].Value {
						return 1, nil
					}
				}
				return 0, nil
			}
		}
		return 1, nil
	} else if len(sc1.DoublePair) > 0 {
		if len(sc1.DoublePair) != 2 || len(sc1.Pair) != 2 {
			return 0, log.Error("no 2x Pair value in scoreCard 1")
		}
		if sc2.StraightFlush || len(sc2.Four) > 0 || sc2.FullHouse || len(sc2.Flush) > 0 || len(sc2.Straight) > 0 || len(sc2.Three) > 0 {
			return -1, nil
		}
		if len(sc2.DoublePair) > 0 {
			if len(sc2.DoublePair) != 2 || len(sc2.Pair) != 2 {
				return 0, log.Error("no 2x Pair value in scoreCard 2")
			}
			if board1[sc1.Pair[0]].Value < board2[sc2.Pair[0]].Value {
				return -1, nil
			} else if board1[sc1.Pair[0]].Value == board2[sc2.Pair[0]].Value {
				if board1[sc1.DoublePair[0]].Value < board2[sc2.DoublePair[0]].Value {
					return -1, nil
				} else if board1[sc1.DoublePair[0]].Value > board2[sc2.DoublePair[0]].Value {
					return 1, nil
				}
				//if same double pair, check high card
				if len(sc1.HighCard) < 1 || len(sc2.HighCard) < 1 {
					return 0, log.Error("no HighCard value in scoreCard 1 or 2")
				}
				if board1[sc1.HighCard[0]].Value < board2[sc2.HighCard[0]].Value {
					return -1, nil
				} else if board1[sc1.HighCard[0]].Value == board2[sc2.HighCard[0]].Value {
					return 0, nil
				}
			}
		}
		return 1, nil
	} else if len(sc1.Pair) > 0 {
		if len(sc1.Pair) != 2 {
			return 0, log.Error("no Pair value in scoreCard 1")
		}
		if sc2.StraightFlush || len(sc2.Four) > 0 || sc2.FullHouse || len(sc2.Flush) > 0 || len(sc2.Straight) > 0 || len(sc2.Three) > 0 || len(sc2.DoublePair) > 0 {
			return -1, nil
		}
		if len(sc2.Pair) > 0 {
			if len(sc2.Pair) != 2 {
				return 0, log.Error("no Pair value in scoreCard 2")
			}
			if board1[sc1.Pair[0]].Value < board2[sc2.Pair[0]].Value {
				return -1, nil
			} else if board1[sc1.Pair[0]].Value == board2[sc2.Pair[0]].Value {
				if len(sc1.HighCard) < 3 || len(sc2.HighCard) < 3 {
					return 0, log.Errorf("no (enough) HighCard value in scoreCard 1 %v or scoreCard 2: %v", sc1.HighCard, sc2.HighCard)
				}
				log.Debugf("Got equal pair, comparing high cards: %v and %v", sc1.HighCard, sc2.HighCard)

				for _, i := range sc1.HighCard[:3] {

					//board is always sorted
					if board1[i].Value < board2[i].Value {
						return -1, nil
					} else if board1[i].Value > board2[i].Value {
						return 1, nil
					}
				}
				return 0, nil
			}
		}
		return 1, nil
	}
	//High card remaining

	if sc2.StraightFlush || len(sc2.Four) > 0 || sc2.FullHouse || len(sc2.Flush) > 0 || len(sc2.Straight) > 0 || len(sc2.Three) > 0 || len(sc2.DoublePair) > 0 || len(sc2.Pair) > 0 {
		return -1, nil
	}
	if len(sc1.HighCard) != 5 || len(sc2.HighCard) != 5 {
		return 0, log.Error("no (enough) HighCard value in scoreCard 1 or 2")
	}
	for _, i := range sc1.HighCard {
		//board is always sorted
		if board1[i].Value < board2[i].Value {
			return -1, nil
		} else if board1[i].Value > board2[i].Value {
			return 1, nil
		}
	}
	//Tried everything, scores are the same
	return 0, nil
}

func (sc *ScoreCard) GetMaxStraight() int {
	//needs sorted board
	board := *(sc.Board)
	maxInARow := 1
	var straightCards []uint8
	currentValue := board[0].Value
	straightCards = append(straightCards, uint8(0))

	for i := range board {
		if board[i].Value == currentValue {
			continue
		} else if board[i].Value != (currentValue - 1) {
			straightCards = []uint8{uint8(i)}
		} else { // == currentValue - 1
			straightCards = append(straightCards, uint8(i))
		}
		currentValue = board[i].Value
		if len(straightCards) > maxInARow {
			maxInARow = len(straightCards)
		}
	}

	return maxInARow
}

func (sc *ScoreCard) GetMaxSuit() int {
	//needs sorted board
	board := *(sc.Board)
	maxSuit := 1
	countCardsBySuit := make(map[uint8][]uint8, 4) //4 suits

	for position, card := range board {
		countCardsBySuit[card.Suit] = append(countCardsBySuit[card.Suit], uint8(position))
	}

	var suit uint8
	for suit = 0; suit < 4; suit++ {
		if len(countCardsBySuit[suit]) > maxSuit {
			maxSuit = len(countCardsBySuit[suit])
		}
	}

	return maxSuit
}

func (sc *ScoreCard) checkStraight() error {
	//needs sorted board
	enoughRoomLeft := true
	board := *(sc.Board)
	var straightCards []uint8
	currentValue := board[0].Value
	straightCards = append(straightCards, uint8(0))

	for i := range board {
		if enoughRoomLeft && len(board)-i < 5 {
			enoughRoomLeft = false
		}
		if board[i].Value == currentValue {
			//is okay, skip this one
			log.Debugf("Pass nb %v == currentValue straightCards: %v", i, straightCards)
			continue
		} else if board[i].Value != (currentValue - 1) {
			if !enoughRoomLeft {
				break
			}
			straightCards = []uint8{uint8(i)}
			log.Debugf("Pass nb %v - != currentValue-1 - straightCards: %v", i, straightCards)
		} else { // == currentValue - 1
			straightCards = append(straightCards, uint8(i))
			log.Debugf("Pass nb %v == currentValue straightCards: %v", i, straightCards)
		}
		currentValue = board[i].Value

		if len(straightCards) == 5 {
			break
		}
	}

	if len(straightCards) != 5 {
		straightCards = []uint8{} // all or nothing
		log.Debug("did not find any straight")
	}
	sc.Straight = straightCards

	return nil
}

func (sc *ScoreCard) checkFlush() error {
	board := *sc.Board
	countCardsBySuit := make(map[uint8][]uint8, 4) //4 suits

	for position, card := range board {
		//log.Debugf("Card position: %v, suit: %v, value: %v", position, card.Suit, card.Value)
		countCardsBySuit[card.Suit] = append(countCardsBySuit[card.Suit], uint8(position))
	}
	//log.Debugf("countCardsBySuit: %v", countCardsBySuit)

	var suit uint8
	for suit = 0; suit < 4; suit++ {
		if len(countCardsBySuit[suit]) >= 5 {
			sc.Flush = countCardsBySuit[suit][:5]
			return nil
		}
	}
	log.Debug("did not find any flush")

	return nil
}

func (sc *ScoreCard) checkXOfAKind() error {
	//needs sorted board
	//Check Four, Three, Two of a kind, fullHouses and double pairs

	board := *sc.Board
	type oakCount struct {
		count   uint8
		indices []uint8
		value   uint8
	}
	log.Debugf("board at start of checkXOfAKind %v", board)
	/* Function that keeps track of occurences of int in a slice */
	countMap := make(map[uint8]oakCount)

	//keep track of occurences, and where (index) in the board the cards form a xoak
	for index, num := range board {
		oakCount := countMap[num.Value]
		oakCount.count++
		oakCount.value = num.Value
		oakCount.indices = append(oakCount.indices, uint8(index))
		countMap[num.Value] = oakCount
	}

	sortedCountMapKeys := make([]uint8, len(countMap))
	for i := range countMap {
		sortedCountMapKeys = append(sortedCountMapKeys, i)
	}
	sort.SliceStable(sortedCountMapKeys, func(i, j int) bool {
		return countMap[sortedCountMapKeys[i]].value > countMap[sortedCountMapKeys[j]].value
	})
	//log.Debugf("countMap containing oakCount %v", countMap)
	//log.Debugf("sortedCountMapKeys %v", sortedCountMapKeys)

	//Check xoak counters and fill scoreCard accordingly
	for _, sortedCountMapKey := range sortedCountMapKeys {
		//Browse ordered board and not the unordered map !
		oakCount := countMap[uint8(sortedCountMapKey)]
		if oakCount.count == 4 {
			sc.Four = oakCount.indices
			continue
		} else if oakCount.count == 3 && len(sc.Four) == 0 {
			if len(sc.Three) > 0 {
				//Found a lower oak3 because board is ordered
				//we need to split it to have a fullHouse
				//Don't forget board is ordered
				sc.Pair = oakCount.indices[:1]
				sc.HighCard = []uint8{}
				sc.FullHouse = true
				break
			}
			sc.Three = oakCount.indices
			if len(sc.Pair) > 0 {
				sc.HighCard = []uint8{}
				sc.FullHouse = true
				break
			}
		} else if oakCount.count == 2 && len(sc.Four) == 0 {
			//log.Debugf("oakCount of pair %#v", oakCount)
			//log.Debugf("sc pair before %v", sc.Pair)
			if len(sc.Pair) > 0 {
				if len(sc.DoublePair) > 0 {
					//Already found higher double pair, do nothing
					//we need to split it to have a double pair + high card
					//Don't forget board is ordered
					sc.HighCard = append(sc.HighCard, oakCount.indices[0])
					log.Debugf("Already found double pair %#v", sc)
					break
				}
				//Found lower pair
				sc.DoublePair = oakCount.indices
				log.Debugf("found double pair %v", sc.DoublePair)
				continue // already 4 cards out of 5
			}
			sc.Pair = oakCount.indices
			//log.Debugf("found pair %#v", sc.Pair)

			//this pair might trigger a fullHouse
			if len(sc.Three) > 0 {
				sc.HighCard = []uint8{}
				sc.FullHouse = true
				break
			}
		} else if oakCount.count == 1 {
			//fill HighCard in desc order of value because board is ordered
			sc.HighCard = append(sc.HighCard, oakCount.indices[0])
		}
	}
	if len(sc.HighCard) > 5 {
		sc.HighCard = sc.HighCard[:5]
	}
	//log.Debugf("sc début du check xoak %#v", sc)

	return nil
}

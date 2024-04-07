package scores

import (
	"errors"
	"goker/cards"
	"goker/pkg/log"
	"sort"
)

type ScoreCard struct {
	*cards.Board
	StraightFlush bool
	Four          []uint8 // 4 of a kind, aka OAK4
	FullHouse     bool
	Flush         []uint8
	Straight      []uint8
	Three         []uint8 // aka OAK3
	DoublePair    []uint8 // aka doubleOAK2, contains lower pair
	Pair          []uint8 // aka OAK2, contains highest pair
	HighCard      []uint8 //may need up to 7 to discriminate
}

func Score(board cards.Board) (*ScoreCard, error) {
	//Create and fill a new ScoreCard
	sc := new(ScoreCard)

	// Make a copy of the board
	sc.Board = &board

	if len(board) > 7 || len(board) < 2 {
		return nil, errors.New("not enough (<2) or too much (>7) cards to evaluate score from")
	}

	//sort by value desc
	sort.SliceStable(board, func(i, j int) bool {
		return board[i].Value > board[j].Value
	})

	err := sc.checkStraight()
	if err != nil {
		return nil, err
	}

	err = sc.checkFlush()
	if err != nil {
		return nil, err
	}

	if len(sc.Straight) > 0 && len(sc.Flush) > 0 {
		sc.StraightFlush = true
	}

	err = sc.checkXOfAKind()
	if err != nil {
		return nil, err
	}

	log.Debugf("sc at the end of Score %v", sc)

	return sc, nil
}

func CompareScoreCards(sc1, sc2 *ScoreCard) (int, error) {
	board1 := *sc1.Board
	board2 := *sc2.Board

	if sc1.StraightFlush {
		if len(sc1.Straight) != 5 {
			return 0, errors.New("flag StraightFlush but no 5 Straight value in scoreCard 1")
		}
		if sc2.StraightFlush {
			if len(sc2.Straight) != 5 {
				return 0, errors.New("flag StraightFlush but no 5 Straight value in scoreCard 2")
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
			return 0, errors.New("no Four value in scoreCard 1")
		}
		if sc2.StraightFlush {
			return -1, nil
		}
		if len(sc2.Four) > 0 {
			if len(sc2.Four) != 4 {
				return 0, errors.New("no Four value in scoreCard 2")
			}
			if board1[sc1.Four[0]].Value < board2[sc2.Four[0]].Value {
				return -1, nil
			} else if board1[sc1.Four[0]].Value == board2[sc2.Four[0]].Value {
				if len(sc1.HighCard) < 1 || len(sc2.HighCard) < 1 {
					return 0, errors.New("no HighCard value in scoreCard 1 or 2")
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
			return 0, errors.New("no 3 Three or 2 Pair value in scoreCard 1")
		}
		if sc2.StraightFlush || len(sc2.Four) > 0 {
			return -1, nil
		}
		if sc2.FullHouse {
			if len(sc2.Three) != 3 || len(sc2.Pair) != 2 {
				return 0, errors.New("no 3 Three or 2 Pair value in scoreCard 2")
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
			return 0, errors.New("no 5 Flush value in scoreCard 1")
		}
		if sc2.StraightFlush || len(sc2.Four) > 0 || sc2.FullHouse {
			return -1, nil
		}
		if len(sc2.Flush) > 0 {
			if len(sc2.Flush) != 5 {
				return 0, errors.New("no 5 Flush value in scoreCard 2")
			}
			for i := range sc1.Flush {
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
			return 0, errors.New("no 5 Straight value in scoreCard 1")
		}
		if sc2.StraightFlush || len(sc2.Four) > 0 || sc2.FullHouse || len(sc2.Flush) > 0 {
			return -1, nil
		}
		if len(sc2.Straight) > 0 {
			if len(sc2.Straight) != 5 {
				return 0, errors.New("no 5 Straight value in scoreCard 2")
			}
			if board1[sc1.Straight[0]].Value < board2[sc2.Straight[0]].Value {
				return -1, nil
			} else if board1[sc1.Four[0]].Value == board2[sc2.Four[0]].Value {
				//We know the rest is the same too as it's a straight
				return 0, nil
			}
		}
		return 1, nil
	} else if len(sc1.Three) > 0 {
		if len(sc1.Three) != 3 {
			return 0, errors.New("no Three value in scoreCard 1")
		}
		if sc2.StraightFlush || len(sc2.Four) > 0 || sc2.FullHouse || len(sc2.Flush) > 0 || len(sc2.Straight) > 0 {
			return -1, nil
		}
		if len(sc2.Three) > 0 {
			if len(sc2.Three) != 3 {
				return 0, errors.New("no Three value in scoreCard 2")
			}
			if board1[sc1.Three[0]].Value < board2[sc2.Three[0]].Value {
				return -1, nil
			} else if board1[sc1.Three[0]].Value == board2[sc2.Three[0]].Value {
				if len(sc1.HighCard) < 2 || len(sc2.HighCard) < 2 {
					return 0, errors.New("no (enough) HighCard value in scoreCard 1 or 2")
				}
				for i := range sc1.HighCard[:2] {
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
			return 0, errors.New("no 2x Pair value in scoreCard 1")
		}
		if sc2.StraightFlush || len(sc2.Four) > 0 || sc2.FullHouse || len(sc2.Flush) > 0 || len(sc2.Straight) > 0 || len(sc2.Three) > 0 {
			return -1, nil
		}
		if len(sc2.DoublePair) > 0 {
			if len(sc2.DoublePair) != 2 || len(sc2.Pair) != 2 {
				return 0, errors.New("no 2x Pair value in scoreCard 2")
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
					return 0, errors.New("no HighCard value in scoreCard 1 or 2")
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
			return 0, errors.New("no Pair value in scoreCard 1")
		}
		if sc2.StraightFlush || len(sc2.Four) > 0 || sc2.FullHouse || len(sc2.Flush) > 0 || len(sc2.Straight) > 0 || len(sc2.Three) > 0 || len(sc2.DoublePair) > 0 {
			return -1, nil
		}
		if len(sc2.Pair) > 0 {
			if len(sc2.Pair) != 2 {
				return 0, errors.New("no Pair value in scoreCard 2")
			}
			if board1[sc1.Pair[0]].Value < board2[sc2.Pair[0]].Value {
				return -1, nil
			} else if board1[sc1.Pair[0]].Value == board2[sc2.Pair[0]].Value {
				if len(sc1.HighCard) < 3 || len(sc2.HighCard) < 3 {
					return 0, errors.New("no (enough) HighCard value in scoreCard 1 or 2")
				}
				for i := range sc1.HighCard[:3] {
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
		return 0, errors.New("no (enough) HighCard value in scoreCard 1 or 2")
	}
	for i := range sc1.HighCard {
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

func (sc *ScoreCard) checkStraight() error {
	//needs sorted board
	enoughRoomLeft := true
	board := *(sc.Board)
	var straightCards []uint8

	for i := 1; i < len(board); i++ {
		if enoughRoomLeft && len(board)-i < 5 {
			enoughRoomLeft = false
		}
		if board[i].Value != board[i-1].Value {
			//is okay, skip this one
			continue
		} else if board[i].Value != board[i-1].Value-1 {
			if !enoughRoomLeft {
				return nil
			}
			straightCards = nil
		}
		straightCards = append(straightCards, uint8(i))

		if len(straightCards) == 5 {
			break
		}
	}

	sc.Straight = straightCards
	return nil
}

func (sc *ScoreCard) checkFlush() error {
	if len(sc.Flush) != 0 {
		//already checked
		return nil
	}
	board := *sc.Board
	enoughRoomLeft := true
	var flushCards []uint8

	for i := 1; i < len(board); i++ {
		if enoughRoomLeft && len(board)-i < 5 {
			enoughRoomLeft = false
		}
		if board[i].Suit != board[i-1].Suit {
			if !enoughRoomLeft {
				return nil
			}
			flushCards = nil //reset
		}
		flushCards = append(flushCards, uint8(i))

		if len(flushCards) == 5 {
			break
		}
	}

	sc.Flush = flushCards

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
	log.Debugf("countMap containing oakCount %v", countMap)
	log.Debugf("sortedCountMapKeys %v", sortedCountMapKeys)

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
				sc.FullHouse = true
				break
			}
			sc.Three = oakCount.indices
			if len(sc.Pair) > 0 {
				sc.FullHouse = true
				break
			}
		} else if oakCount.count == 2 && len(sc.Four) == 0 {
			//log.Debug("oakCount of pair", oakCount)
			//log.Debug("sc pair", sc)
			if len(sc.Pair) > 0 {
				if len(sc.DoublePair) > 0 {
					//Already found higher double pair, do nothing
					//we need to split it to have a double pair + high card
					//Don't forget board is ordered
					sc.HighCard = append(sc.HighCard, oakCount.indices[0])
					//log.Debug("Already found double pair", sc)
					break
				}
				//Found lower pair
				sc.DoublePair = oakCount.indices
				//log.Debug("found double pair", sc)
				continue // already 4 cards out of 5
			}
			//log.Debug("found pair", sc)
			sc.Pair = oakCount.indices
		} else if oakCount.count == 1 {
			//fill HighCard in desc order of value because board is ordered
			sc.HighCard = append(sc.HighCard, oakCount.indices[0])
		}
	}
	//log.Debug("sc début du check xoak", sc)

	return nil
}

package scores

import (
	"errors"
	"goker/cards"
	"sort"
)

type card cards.Card
type ScoreCard struct {
	StraightFlush uint8
	Four          uint8
	FullHouse     uint16
	Flush         bool
	Straight      uint8
	Three         uint8
	DoublePair    uint8
	Pair          uint8
	HighCard      []uint8 //may need up to 7 to discriminate
}

func Score(board []card) (*ScoreCard, error) {
	var sc *ScoreCard

	if len(board) > 7 || len(board) < 2 {
		return nil, errors.New("ot enough (<2) or too much (>7) cards to evaluate score from")
	}

	//sort by value desc
	sort.SliceStable(board, func(i, j int) bool {
		return board[i].Value > board[j].Value
	})

	/*
		//need to check in proper order
		hasStraightFlush, err := sc.checkStraightFlush(board)
		if err != nil {
			return nil,err
		}
		has4OAK, err := sc.checkXOfAKind(board)
		if err != nil {
			return nil,err
		}
	*/
	//for _, card := range board {
	//}

	return sc, nil
}

func (sc *ScoreCard) checkStraightFlush(board []card) (bool, error) {
	//needs sorted board
	if sc.StraightFlush > 0 {
		//already checked
		return true, nil
	}
	hasStraight, err := sc.checkStraight(board)
	if err != nil {
		return false, err
	}
	hasFlush, err := sc.checkFlush(board)
	if err != nil {
		return false, err
	}

	if hasStraight && hasFlush {
		sc.StraightFlush = sc.Straight
		return true, nil
	}
	return false, nil
}

func (sc *ScoreCard) checkStraight(board []card) (bool, error) {
	//needs sorted board
	enoughRoomLeft := true
	straightHighCard := board[0].Value

	for i := 1; i < len(board); i++ {
		if enoughRoomLeft && len(board)-i < 5 {
			enoughRoomLeft = false
		}
		if board[i].Value != board[i-1].Value-1 {
			if !enoughRoomLeft {
				return false, nil
			} else {
				straightHighCard = board[i].Value
			}
		}
	}

	sc.Straight = straightHighCard
	return true, nil
}

func (sc *ScoreCard) checkFlush(board []card) (bool, error) {
	if sc.Flush {
		//already checked
		return true, nil
	}
	enoughRoomLeft := true

	for i := 0; i < len(board); i++ {
		if enoughRoomLeft && len(board)-i < 5 {
			enoughRoomLeft = false
		}
		if board[i].Suit != board[i-1].Suit {
			if !enoughRoomLeft {
				return false, nil
			}
		}
	}

	sc.Flush = true

	return true, nil
}

func (sc *ScoreCard) checkXOfAKind(board []card, n int) (bool, error) {
	enoughRoomLeft := true
	highCard := board[0].Value

	//Manage double pairs
	for i := 1; i < len(board); i++ {
		if enoughRoomLeft && len(board)-i < n {
			enoughRoomLeft = false
		}
		if board[i].Value != board[i-1].Value {
			if !enoughRoomLeft {
				return false, nil
			} else {
				highCard = board[i].Value
			}
		}
	}
	switch n {
	case 4:
		sc.Four = highCard
	case 3:
		sc.Three = highCard
	case 2:
		sc.Pair = highCard
	default:
		return false, errors.New("invalid X of a kind to search")
	}
	return true, nil
}

func (sc *ScoreCard) checkFullHouse(board []card) (bool, error) {
	//needs sorted board
	has3, err := sc.checkXOfAKind(board, 3)
	if err != nil {
		return false, err
	}
	if !has3 {
		return false, nil
	}

	//filteredBoard := board
	//filter out OAK3
	//for i, card := range filteredBoard {
	//}
	has2, err := sc.checkXOfAKind(board, 2)
	if err != nil {
		return false, err
	}
	if !has2 {
		return false, err
	}

	sc.FullHouse = uint16(sc.Three*100 + sc.Pair)
	return true, nil
}

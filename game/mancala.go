package game

import (
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
)

type MancalaEnv struct {
	Board      *Board
	SideToMove *Side
	NorthMoved bool
	OurSide    *Side
}

func NewMancalaEnv() *MancalaEnv {
	m := new(MancalaEnv)
	return m.Reset()
}

func (m *MancalaEnv) Reset() *MancalaEnv {
	board, err := NewBoard(7, 7)
	if err != nil {
		panic(err)
	}
	m.Board = board
	m.SideToMove = NewSide(southIndex)
	m.NorthMoved = false
	m.OurSide = NewSide(southIndex)
	return m
}

func (m MancalaEnv) Clone(otherState MancalaEnv) *MancalaEnv {
	fmt.Println("Clone method is broken, doesn't work, fix it")
	board := Board.Clone(otherState.Board)
	sideToMove := NewSide(southIndex)
	copier.Copy(&sideToMove, &otherState.SideToMove)
	northMoved := false
	copier.Copy(&northMoved, &otherState.NorthMoved)

	cloneGame := NewMancalaEnv()
	cloneGame.Board = board
	cloneGame.SideToMove = sideToMove
	cloneGame.NorthMoved = northMoved
	return cloneGame
}

func (m *MancalaEnv) GetLegalMoves() []*Move {
	return m.GetStateLegalActions(m.Board, m.SideToMove, m.NorthMoved)
}

func (m *MancalaEnv) PerformMove(move *Move) int {
	seedsInStoreBefore := m.Board.SeedsInStore(move.Side)
	// pie move
	if move.Index == 0 {
		m.OurSide = m.OurSide.Opposite()
	}
	sideToMove, err := m.MakeMove(m.Board, move, m.NorthMoved)
	if err != nil {
		panic(err)
	}
	m.SideToMove = sideToMove
	if move.Side.IsNorth {
		m.NorthMoved = true
	}
	seedsInStoreAfter := m.Board.SeedsInStore(move.Side)

	// return a partial reward proportional to the number of captured seeds
	return (seedsInStoreAfter - seedsInStoreBefore) / 100.0
}

func (m MancalaEnv) GetStateLegalActions(board *Board, side *Side, northMoved bool) []*Move {
	// If this is the first move of NORTH, then NORTH can use the pie rule action
	legalMoves := []*Move{}
	if northMoved || side.IsSouth() {
		return legalMoves
	} else {
		move, err := NewMove(side, 0)
		if err != nil {
			panic(err)
		}
		legalMoves = append(legalMoves, move)
	}

	for i := 1; i < board.Holes+1; i++ {
		if board.Board[side.Index()][i] > 0 {
			move, err := NewMove(side, i)
			if err != nil {
				panic(err)
			}
			legalMoves = append(legalMoves, move)
		}
	}
	return legalMoves
}

func (m MancalaEnv) HolesEmpty(board *Board, side *Side) bool {
	for hole := 1; hole < board.Holes+1; hole++ {
		seeds, err := board.Seeds(side, hole)
		if err != nil {
			panic(err)
		}
		if seeds > 0 {
			return false
		}
	}
	return true
}

func (m MancalaEnv) GameOver(board *Board) bool {
	if m.HolesEmpty(board, NewSide(northIndex)) || m.HolesEmpty(board, NewSide(southIndex)) {
		return true
	}
	return false
}

func (m MancalaEnv) MakeMove(board *Board, move *Move, northMoved bool) (*Side, error) {
	if !m.IsLegalAction(board, move, northMoved) {
		return nil, errors.New("illegalMove: an illegal m ove was tried to play")
	}
	// check for pie rule/move
	if move.Index == 0 {
		m.SwitchSides(board)
		return move.Side.Opposite(), nil
	}

	seedsToSow, err := board.Seeds(move.Side, move.Index)
	if err != nil {
		panic(err)
	}
	board.SetSeeds(move.Side, move.Index, 0)

	holes := board.Holes
	// Place seeds in all holes excepting the opponent's store
	receivingHoles := 2*holes + 1
	// Rounds needed to sow all the seeds
	rounds := seedsToSow / receivingHoles
	// Seeds remaining after all the rounds
	remainingSeeds := seedsToSow % receivingHoles

	// Sow the seeds for the full rounds
	if rounds != 0 {
		for hole := 1; hole < holes+1; hole++ {
			board.AddSeeds(NewSide(northIndex), hole, rounds)
			board.AddSeeds(NewSide(southIndex), hole, rounds)
		}
		board.AddSeedsToStore(move.Side, rounds)
	}

	// sow remaining seeds
	sowSide := move.Side
	sowHole := move.Index
	for i := 0; i < remainingSeeds; i++ {
		sowHole++
		if sowHole == 1 {
			sowSide = sowSide.Opposite()
		}
		if sowHole > holes {
			if sowSide == move.Side {
				sowHole = 0
				board.AddSeedsToStore(sowSide, 1)
				continue
			} else {
				sowSide = sowSide.Opposite()
				sowHole = 1
			}
		}
		board.AddSeeds(sowSide, sowHole, 1)
	}

	// Capture the opponent's seeds from the opposite hole if the last seed
	// is placed in an empty hole and there are seeds in the opposite hole
	sowSeeds, err := board.Seeds(sowSide, sowHole)
	if err != nil {
		panic(err)
	}
	sowSeedsOp, err := board.SeedsOp(sowSide, sowHole)
	if err != nil {
		panic(err)
	}
	if sowSide == move.Side && sowHole > 0 && sowSeeds == 1 && sowSeedsOp > 0 {
		err := board.AddSeedsToStore(move.Side, 1)
		if err != nil {
			panic(err)
		}
		err = board.SetSeeds(move.Side, sowHole, 0)
		if err != nil {
			panic(err)
		}
		err = board.SetSeedsOp(move.Side, sowHole, 0)
	}

	// if the game is over, collect the seeds not in the store and put them there
	if m.GameOver(board) {
		finishedSide := NewSide(northIndex)
		if m.HolesEmpty(board, NewSide(southIndex)) {
			finishedSide = NewSide(northIndex)
		}
		seeds := 0
		collectingSide := finishedSide.Opposite()
		for hole := 1; hole < board.Holes+1; hole++ {
			collectingSideSeeds, err := board.Seeds(collectingSide, hole)
			if err != nil {
				panic(err)
			}
			seeds += collectingSideSeeds
		}
		err := board.AddSeedsToStore(collectingSide, seeds)
		if err != nil {
			panic(err)
		}
	}

	// return the side which is to move next
	if sowHole == 0 && (move.Side.IsNorth || northMoved) {
		return move.Side, nil
	}
	return move.Side.Opposite(), nil
}

func (m MancalaEnv) SwitchSides(board *Board) {
	for hole := 1; hole < board.Holes+1; hole++ {
		temp := board.Board[northIndex][hole]
		board.Board[northIndex][hole] = board.Board[southIndex][hole]
		board.Board[southIndex][hole] = temp
	}
}

func (m MancalaEnv) IsLegalAction(board *Board, move *Move, northMoved bool) bool {
	actions := m.GetStateLegalActions(board, move.Side, northMoved)
	for _, action := range actions {
		if move.Index == action.Index {
			return true
		}
	}
	return false
}

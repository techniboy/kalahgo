package game

import (
	"errors"
	"log"

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
		log.Panic(err)
	}
	m.Board = board
	m.SideToMove = NewSide(SideSouth)
	m.NorthMoved = false
	m.OurSide = NewSide(SideSouth)
	return m
}

func (m *MancalaEnv) Clone() *MancalaEnv {
	board := m.Board.Clone()
	sideToMove := NewSide(SideSouth)
	copier.Copy(&sideToMove, &m.SideToMove)
	northMoved := false
	copier.Copy(&northMoved, &m.NorthMoved)

	cloneGame := NewMancalaEnv()
	cloneGame.Board = board
	cloneGame.SideToMove = sideToMove
	cloneGame.NorthMoved = northMoved
	return cloneGame
}

func (m *MancalaEnv) LegalMoves() []*Move {
	return m.StateLegalActions(m.Board, m.SideToMove, m.NorthMoved)
}

func (m *MancalaEnv) PerformMove(move *Move) float64 {
	seedsInStoreBefore := m.Board.SeedsInStore(move.Side)
	// pie move
	if move.Index == 0 {
		m.OurSide = m.OurSide.Opposite()
	}
	sideToMove, err := m.MakeMove(m.Board, move, m.NorthMoved)
	if err != nil {
		log.Panic(err)
	}
	m.SideToMove = sideToMove
	if move.Side.IsNorth {
		m.NorthMoved = true
	}
	seedsInStoreAfter := m.Board.SeedsInStore(move.Side)
	// return a partial reward proportional to the number of captured seeds
	return float64((seedsInStoreAfter - seedsInStoreBefore)) / 100.0
}

func (m MancalaEnv) StateLegalActions(board *Board, side *Side, northMoved bool) []*Move {
	// If this is the first move of NORTH, then NORTH can use the pie rule action
	var legalMoves []*Move
	if northMoved || side.IsSouth() {
		legalMoves = []*Move{}
	} else {
		move, err := NewMove(side, 0)
		if err != nil {
			log.Panic(err)
		}
		legalMoves = append(legalMoves, move)
	}

	for i := 1; i < board.Holes+1; i++ {
		if board.Board[side.Index()][i] > 0 {
			move, err := NewMove(side, i)
			if err != nil {
				log.Panic(err)
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
			log.Panic(err)
		}
		if seeds > 0 {
			return false
		}
	}
	return true
}

func (m MancalaEnv) GameOver(board *Board) bool {
	if m.HolesEmpty(board, NewSide(SideNorth)) || m.HolesEmpty(board, NewSide(SideSouth)) {
		return true
	}
	return false
}

func (m *MancalaEnv) IsGameOver() bool {
	return m.GameOver(m.Board)
}

func (m MancalaEnv) MakeMove(board *Board, move *Move, northMoved bool) (*Side, error) {
	if !m.IsLegalAction(board, move, northMoved) {
		return nil, errors.New("illegalMove: an illegal move was tried to play")
	}
	// check for pie rule/move
	if move.Index == 0 {
		m.SwitchSides(board)
		return move.Side.Opposite(), nil
	}

	seedsToSow, err := board.Seeds(move.Side, move.Index)
	if err != nil {
		log.Panic(err)
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
			board.AddSeeds(NewSide(SideNorth), hole, rounds)
			board.AddSeeds(NewSide(SideSouth), hole, rounds)
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
			if sowSide.Index() == move.Side.Index() {
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
	if sowHole > 0 && sowSide.Index() == move.Side.Index() {
		sowSeeds, err := board.Seeds(sowSide, sowHole)
		if err != nil {
			log.Panic(err)
		}
		sowSeedsOp, err := board.SeedsOp(sowSide, sowHole)
		if err != nil {
			log.Panic(err)
		}
		if sowSeeds == 1 && sowSeedsOp > 0 {
			err := board.AddSeedsToStore(move.Side, 1+sowSeedsOp)
			if err != nil {
				log.Panic(err)
			}
			err = board.SetSeeds(move.Side, sowHole, 0)
			if err != nil {
				log.Panic(err)
			}
			err = board.SetSeedsOp(move.Side, sowHole, 0)
			if err != nil {
				log.Panic(err)
			}
		}
	}

	// if the game is over, collect the seeds not in the store and put them there
	if m.GameOver(board) {
		finishedSide := NewSide(SideNorth)
		if m.HolesEmpty(board, NewSide(SideSouth)) {
			finishedSide = NewSide(SideNorth)
		}
		seeds := 0
		collectingSide := finishedSide.Opposite()
		for hole := 1; hole < board.Holes+1; hole++ {
			collectingSideSeeds, err := board.Seeds(collectingSide, hole)
			if err != nil {
				log.Panic(err)
			}
			seeds += collectingSideSeeds
		}
		err := board.AddSeedsToStore(collectingSide, seeds)
		if err != nil {
			log.Panic(err)
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
		temp := board.Board[SideNorth][hole]
		board.Board[SideNorth][hole] = board.Board[SideSouth][hole]
		board.Board[SideSouth][hole] = temp
	}
}

func (m MancalaEnv) IsLegalAction(board *Board, move *Move, northMoved bool) bool {
	actions := m.StateLegalActions(board, move.Side, northMoved)
	for _, action := range actions {
		if move.Index == action.Index {
			return true
		}
	}
	return false
}

func (m *MancalaEnv) ComputeFinalReward(side *Side) int {
	return m.Board.SeedsInStore(side) - m.Board.SeedsInStore(side.Opposite())
}

func (m *MancalaEnv) ComputeEndGameReward(side *Side) (float64, error) {
	if !m.GameOver(m.Board) {
		return -1, errors.New("compute_end_game_reward should only be called at end of the game")
	}
	reward := m.ComputeFinalReward(side)
	if reward > 0 {
		return 1, nil
	} else if reward < 0 {
		return 0, nil
	}
	return 0.5, nil
}

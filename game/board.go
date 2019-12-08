package game

import (
	"errors"
	"log"
)

type Board struct {
	Holes int
	Board [][]int
}

func NewBoard(holes int, seeds int) (*Board, error) {
	b := new(Board)
	if holes < 1 {
		return nil, errors.New("valueError: there has to be atleast one hole")
	}
	if seeds < 0 {
		return nil, errors.New("valueError: there has to be non-negative no. of seeds")
	}
	b.Holes = holes
	b.Board = make([][]int, 2)
	for i := range b.Board {
		b.Board[i] = make([]int, holes+1)
	}
	for hole := 1; hole < holes+1; hole++ {
		b.Board[SideNorth][hole] = seeds
		b.Board[SideSouth][hole] = seeds
	}
	return b, nil
}

func (b *Board) Clone() *Board {
	cloneBoard, err := NewBoard(b.Holes, 0)
	if err != nil {
		log.Panic(err)
	}
	for hole := 0; hole < b.Holes+1; hole++ {
		cloneBoard.Board[SideNorth][hole] = b.Board[SideNorth][hole]
		cloneBoard.Board[SideSouth][hole] = b.Board[SideSouth][hole]
	}
	return cloneBoard
}

func (b Board) Seeds(side *Side, hole int) (int, error) {
	if hole < 1 || hole > b.Holes {
		return -1, errors.New("valueError: hole number must be between 1 and no. of holes")
	}
	return b.Board[side.Index()][hole], nil
}

func (b *Board) SetSeeds(side *Side, hole int, seeds int) error {
	if hole < 1 || hole > b.Holes {
		return errors.New("valueError: hole number must be between 1 and no. of holes")
	}
	if seeds < 0 {
		return errors.New("valueError: there has to be a non-negative no. of seeds")
	}
	b.Board[side.Index()][hole] = seeds
	return nil
}

func (b Board) SeedsOp(side *Side, hole int) (int, error) {
	if hole < 1 || hole > b.Holes {
		return -1, errors.New("valueError: hole number must be between 1 and no. of holes")
	}
	return b.Board[side.Opposite().Index()][b.Holes+1-hole], nil
}

func (b *Board) SetSeedsOp(side *Side, hole int, seeds int) error {
	if hole < 1 || hole > b.Holes {
		return errors.New("valueError: hole number must be between 1 and no. of holes")
	}
	if seeds < 0 {
		return errors.New("valueError: there has to be a non-negative no. of seeds")
	}
	b.Board[side.Opposite().Index()][b.Holes+1-hole] = seeds
	return nil
}

func (b *Board) AddSeeds(side *Side, hole int, seeds int) error {
	if hole < 1 || hole > b.Holes {
		return errors.New("valueError: hole number must be between 1 and no. of holes")
	}
	if seeds < 0 {
		return errors.New("valueError: there has to be a non-negative no. of seeds")
	}
	b.Board[side.Index()][hole] += seeds
	return nil
}

func (b *Board) AddSeedsToStore(side *Side, seeds int) error {
	if seeds < 0 {
		return errors.New("valueError: there has to be a non-negative no. of seeds")
	}
	b.Board[side.Index()][0] += seeds
	return nil
}

func (b Board) SeedsInStore(side *Side) int {
	return b.Board[side.Index()][0]
}

func (b *Board) SetSeedsInStore(side *Side, seeds int) error {
	if seeds < 0 {
		return errors.New("valueError: there has to be a non-negative no. of seeds")
	}
	b.Board[side.Index()][0] = seeds
	return nil
}

func (b *Board) IsSeedable(side *Side, hole int) bool {
	for otherHole := hole - 1; otherHole > 0; otherHole-- {
		seeds, err := b.Seeds(side, otherHole)
		if err != nil {
			log.Panic(err)
		}
		if seeds == hole-otherHole {
			return true
		}
	}
	return false
}

package game

import (
	"errors"
	"fmt"
)

type Board struct {
	Holes int
	Board [][]int
}

func (b *Board) NewBoard(holes int, seeds int) (*Board, error) {
	if holes < 1 {
		return nil, errors.New("valueError: there has to be atleast one hole")
	}
	if seeds < 0 {
		return nil, errors.New("valueError: there has to be non-negative no. of seeds")
	}

	b.Holes = holes
	b.Board[holes+1][2] = 0
	for hole := 1; hole < holes+1; hole++ {
		b.Board[northIndex][hole] = seeds
		b.Board[southIndex][hole] = seeds
	}
	return b, nil
}

func (b *Board) Clone(orginalBoard *Board) *Board {
	fmt.Println("function not implemented yet")
	return b
}

func (b Board) GetSeeds(side Side, hole int) (int, error) {
	if hole < 1 || hole > b.Holes {
		return -1, errors.New("valueError: hole number must be between 1 and no. of holes")
	}
	return b.Board[side.Index()][hole], nil
}

func (b *Board) SetSeeds(side Side, hole int, seeds int) error {
	if hole < 1 || hole > b.Holes {
		return errors.New("valueError: hole number must be between 1 and no. of holes")
	}
	if seeds < 0 {
		return errors.New("valueError: there has to be a non-negative no. of seeds")
	}
	b.Board[side.Index()][hole] = seeds
	return nil
}

func (b Board) GetSeedsOp(side Side, hole int) (int, error) {
	if hole < 1 || hole > b.Holes {
		return -1, errors.New("valueError: hole number must be between 1 and no. of holes")
	}
	return b.Board[side.Opposite().Index()][hole], nil
}

func (b *Board) SetSeedsOp(side Side, hole int, seeds int) error {
	if hole < 1 || hole > b.Holes {
		return errors.New("valueError: hole number must be between 1 and no. of holes")
	}
	if seeds < 0 {
		return errors.New("valueError: there has to be a non-negative no. of seeds")
	}
	b.Board[side.Opposite().Index()][hole] = seeds
	return nil
}

func (b Board) GetSeedsInStore(side Side) int {
	return b.Board[side.Index()][0]
}

func (b *Board) IsSeedable(side Side, hole int) bool {
	for otherHole := hole - 1; otherHole < 0; otherHole-- {
		seeds, err := b.GetSeeds(side, otherHole)
		if err != nil {
			panic(err)
		}
		if seeds == hole-otherHole {
			return true
		}
	}
	return false
}

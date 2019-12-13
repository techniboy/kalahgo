package game

import (
	"errors"
)

type Board struct {
	Holes  int
	LBoard []int
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
	b.LBoard = make([]int, 2*(holes+1))
	for hole := 1; hole < holes+1; hole++ {
		b.SetHoleVal(seeds, SideNorth, hole)
		b.SetHoleVal(seeds, SideSouth, hole)
	}
	return b, nil
}

func (b *Board) HoleVal(r int, c int) int {
	return b.LBoard[r*(b.Holes+1)+c]
}

func (b *Board) SetHoleVal(val int, r int, c int) {
	b.LBoard[r*(b.Holes+1)+c] = val
}

func (b *Board) Clone() *Board {
	cloneBoard, err := NewBoard(b.Holes, 0)
	if err != nil {
		panic(err)
	}
	for hole := 0; hole < b.Holes+1; hole++ {
		cloneBoard.SetHoleVal(b.HoleVal(SideNorth, hole), SideNorth, hole)
		cloneBoard.SetHoleVal(b.HoleVal(SideSouth, hole), SideSouth, hole)
	}
	return cloneBoard
}

func (b Board) Seeds(side *Side, hole int) (int, error) {
	if hole < 1 || hole > b.Holes {
		return -1, errors.New("valueError: hole number must be between 1 and no. of holes")
	}
	return b.HoleVal(side.Index(), hole), nil
}

func (b *Board) SetSeeds(side *Side, hole int, seeds int) error {
	if hole < 1 || hole > b.Holes {
		return errors.New("valueError: hole number must be between 1 and no. of holes")
	}
	if seeds < 0 {
		return errors.New("valueError: there has to be a non-negative no. of seeds")
	}
	b.SetHoleVal(seeds, side.Index(), hole)
	return nil
}

func (b Board) SeedsOp(side *Side, hole int) (int, error) {
	if hole < 1 || hole > b.Holes {
		return -1, errors.New("valueError: hole number must be between 1 and no. of holes")
	}
	return b.HoleVal(side.Opposite().Index(), b.Holes+1-hole), nil
}

func (b *Board) SetSeedsOp(side *Side, hole int, seeds int) error {
	if hole < 1 || hole > b.Holes {
		return errors.New("valueError: hole number must be between 1 and no. of holes")
	}
	if seeds < 0 {
		return errors.New("valueError: there has to be a non-negative no. of seeds")
	}
	b.SetHoleVal(seeds, side.Opposite().Index(), b.Holes+1-hole)
	return nil
}

func (b *Board) AddSeeds(side *Side, hole int, seeds int) error {
	if hole < 1 || hole > b.Holes {
		return errors.New("valueError: hole number must be between 1 and no. of holes")
	}
	if seeds < 0 {
		return errors.New("valueError: there has to be a non-negative no. of seeds")
	}
	b.SetHoleVal(b.HoleVal(side.Index(), hole)+seeds, side.Index(), hole)
	return nil
}

func (b *Board) AddSeedsToStore(side *Side, seeds int) error {
	if seeds < 0 {
		return errors.New("valueError: there has to be a non-negative no. of seeds")
	}
	b.SetHoleVal(b.HoleVal(side.Index(), 0)+seeds, side.Index(), 0)
	return nil
}

func (b Board) SeedsInStore(side *Side) int {
	return b.HoleVal(side.Index(), 0)
}

func (b *Board) SetSeedsInStore(side *Side, seeds int) error {
	if seeds < 0 {
		return errors.New("valueError: there has to be a non-negative no. of seeds")
	}
	b.SetHoleVal(seeds, side.Index(), 0)
	return nil
}

func (b *Board) IsSeedable(side *Side, hole int) bool {
	for otherHole := hole - 1; otherHole > 0; otherHole-- {
		seeds, err := b.Seeds(side, otherHole)
		if err != nil {
			panic(err)
		}
		if seeds == hole-otherHole {
			return true
		}
	}
	return false
}

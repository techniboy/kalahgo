package game

import (
	"errors"
	"log"
)

// Move represents a whole (if greater than 1) or the pie action if 0.
type Move struct {
	Side  *Side
	Index int
}

// NewMove is a constructor for Move, might be needed or not LOL
func NewMove(side *Side, index int) (*Move, error) {
	m := new(Move)
	if index < 0 || index > 7 {
		return nil, errors.New("valueError: move number must be strictly greater than 0 and less than 8")
	}
	m.Side = side
	m.Index = index
	return m, nil
}

func (m *Move) Clone() *Move {
	cloneMove, err := NewMove(NewSide(m.Side.Index()), m.Index)
	if err != nil {
		log.Panic(err)
	}
	return cloneMove
}

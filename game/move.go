package game

import "fmt"

// Move represents a whole (if greater than 1) or the pie action if 0.
type Move struct {
	Side  Side
	Index int
}

// NewMove is a constructor for Move, might be needed or not LOL
func (m *Move) NewMove(side Side, index int) *Move {
	if index < 0 || index > 7 {
		fmt.Println("valueError: move number must be strictly greater than 0 and less than 8")
	}
	m.Side = side
	m.Index = index
	return m
}

package game

var northIndex = 0
var southIndex = 1

type Side struct {
	IsNorth bool
}

// constructor for Side
func NewSide(sideIndex int) *Side {
	s := new(Side)
	northIndex = 0
	southIndex = 1
	if sideIndex == northIndex {
		s.IsNorth = true
	}
	return s
}

// return side index
func (s Side) Index() int {
	if s.IsNorth {
		return northIndex
	}
	return southIndex
}

// return opposite index
func (s *Side) Opposite() *Side {
	s.IsNorth = !s.IsNorth
	return s
}

// return side as string
func (s Side) ToString() string {
	if s.IsNorth {
		return "North"
	}
	return "South"
}

// checks whether side is south, think of this as syntactic sugar
func (s Side) IsSouth() bool {
	if s.IsNorth {
		return false
	}
	return true
}

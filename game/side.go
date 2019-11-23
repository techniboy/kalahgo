package game

var northIndex = 0
var southIndex = 1

// Side tells us which player(side) is playing
type Side struct {
	IsNorth bool
}

// NewSide is a constructor for Side
func NewSide(sideIndex int) *Side {
	s := new(Side)
	northIndex = 0
	southIndex = 1
	if sideIndex == northIndex {
		s.IsNorth = true
	}
	return s
}

// Index returns index of the side (0 north; 1 south)
func (s Side) Index() int {
	if s.IsNorth {
		return northIndex
	}
	return southIndex
}

// Opposite returns other side
func (s *Side) Opposite() *Side {
	if s.IsNorth {
		return NewSide(southIndex)
	}
	return NewSide(northIndex)
}

// ToString returns string representation of side
func (s Side) ToString() string {
	if s.IsNorth {
		return "North"
	}
	return "South"
}

// IsSouth returns whether to side playing is south (or not)
func (s Side) IsSouth() bool {
	if s.IsNorth {
		return false
	}
	return true
}

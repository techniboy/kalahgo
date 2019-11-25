package game

var SideNorth = 0
var SideSouth = 1

// Side tells us which player(side) is playing
type Side struct {
	IsNorth bool
}

// NewSide is a constructor for Side
func NewSide(sideIndex int) *Side {
	s := new(Side)
	SideNorth = 0
	SideSouth = 1
	if sideIndex == SideNorth {
		s.IsNorth = true
	}
	return s
}

// Index returns index of the side (0 north; 1 south)
func (s Side) Index() int {
	if s.IsNorth {
		return SideNorth
	}
	return SideSouth
}

// Opposite returns other side
func (s *Side) Opposite() *Side {
	if s.IsNorth {
		return NewSide(SideSouth)
	}
	return NewSide(SideNorth)
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

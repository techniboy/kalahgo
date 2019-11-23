package protocol

// MoveTurn describes types of moves
type MoveTurn struct {
	End   bool
	Again bool
	Move  int
}

/*
// getter for end in MoveTurn
func (moveTurn MoveTurn) End() bool {
	return moveTurn.end
}

// getter for again in MoveTurn
func (moveTurn MoveTurn) Again() bool {
	return moveTurn.again
}

// getter for move in MoveTurn
func (moveTurn MoveTurn) Move() int {
	return moveTurn.move
}

// setter for end in MoveTurn
func (moveTurn *MoveTurn) SetEnd(end bool) {
	moveTurn.end = end
}

// setter for again in MoveTurn
func (moveTurn *MoveTurn) SetAgain(again bool) {
	moveTurn.again = again
}

// setter for move in MoveTurn
func (moveTurn *MoveTurn) SetMove(move int) {
	moveTurn.move = move
}
*/

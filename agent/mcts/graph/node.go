package graph

import (
	"github.com/techniboy/kalahgo/game"
)

type Node struct {
	Visits          float64
	Reward          float64
	State           *game.MancalaEnv
	Children        []*Node
	Parent          *Node
	Move            *game.Move
	UnexploredMoves []*game.Move
}

func NewNode(state *game.MancalaEnv, move *game.Move, parent *Node) *Node {
	n := new(Node)
	n.State = state
	n.Parent = parent
	n.Move = move
	n.UnexploredMoves = state.LegalMoves()
	return n
}

func (n *Node) Clone() *Node {
	var (
		cloneMove   *game.Move
		cloneParent *Node
	)
	if n.Move == nil {
		cloneMove = nil
	} else {
		var err error
		cloneMove, err = game.NewMove(game.NewSide(n.Move.Side.Index()), n.Move.Index)
		if err != nil {
			panic(err)
		}
	}
	if n.Parent == nil {
		cloneParent = nil
	} else {
		cloneParent = n.Parent.Clone()
	}
	cloneNode := NewNode(n.State.Clone(), cloneMove, cloneParent)
	cloneNode.Visits = n.Visits
	cloneNode.Reward = n.Reward
	return cloneNode
}

func (n *Node) InsertChild(child *Node) {
	n.Children = append(n.Children, child)
	for i, move := range n.UnexploredMoves {
		if move.Side.Index() == child.Move.Side.Index() && move.Index == child.Move.Index {
			lenUnexplored := len(n.UnexploredMoves)
			n.UnexploredMoves[i] = n.UnexploredMoves[lenUnexplored-1]
			n.UnexploredMoves[lenUnexplored-1] = nil
			n.UnexploredMoves = n.UnexploredMoves[:lenUnexplored-1]
			break
		}
	}
}

func (n *Node) Update(reward float64) {
	n.Reward += reward
	n.Visits++
}

func (n *Node) UnvisitedChildren() []*Node {
	unvisitedChildren := []*Node{}
	for _, child := range n.Children {
		if child.Visits <= 1 {
			unvisitedChildren = append(unvisitedChildren, child)
		}
	}
	return unvisitedChildren
}

func (n *Node) IsFullyExplored() bool {
	return len(n.UnexploredMoves) == 0
}

func (n *Node) IsTerminal() bool {
	return len(n.State.LegalMoves()) == 0
}

package policy

import (
	"log"
	"math/rand"

	"github.com/techniboy/kalahgo/agent/mcts/graph"
)

func MctpSelect(n *graph.Node) *graph.Node {
	for !n.IsTerminal() {
		if !n.IsFullyExplored() {
			return MctpExpand(n)
		} else {
			var err error
			n, err = graph.SelectBestChild(n)
			if err != nil {
				log.Panic(err)
			}
		}
	}
	return n
}

func MctpExpand(n *graph.Node) *graph.Node {
	childExpansionMove := n.UnexploredMoves[rand.Intn(len(n.UnexploredMoves))]
	childState := n.State.Clone()
	childState.PerformMove(childExpansionMove)
	childNode := graph.NewNode(childState, childExpansionMove, n)
	// childNode.Update(moveReward)
	n.InsertChild(childNode)
	return childNode
}

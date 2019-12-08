package policy

import (
	"math/rand"

	"github.com/techniboy/kalahgo/game"

	"github.com/techniboy/kalahgo/agent/mcts/graph"
)

func McdpSimuilate(root *graph.Node) *game.MancalaEnv {
	node := root.Clone()
	for !node.IsTerminal() {
		randomLegalMove := node.State.LegalMoves()[rand.Intn(len(node.State.LegalMoves()))]
		node.State.PerformMove(randomLegalMove)
	}
	return node.State
}

package policy

import (
	"math/rand"
	"time"

	"github.com/techniboy/kalahgo/game"

	"github.com/techniboy/kalahgo/agent/mcts/graph"
)

func McdpSimuilate(root *graph.Node) *game.MancalaEnv {
	node := root.Clone()
	for !node.IsTerminal() {
		rand.Seed(time.Now().Unix())
		randomLegalMove := node.State.LegalMoves()[rand.Intn(len(node.State.LegalMoves()))]
		node.State.PerformMove(randomLegalMove)
	}
	return node.State
}

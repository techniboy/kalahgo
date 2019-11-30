package policy

import (
	"math/rand"

	"github.com/techniboy/kalahgo/agent/mcts/graph"
)

func McdpSimuilate(n *graph.Node) {
	for !n.IsTerminal() {
		randomLegalMove := n.State.LegalMoves()[rand.Intn(len(n.State.LegalMoves()))]
		n.Update(float64(n.State.PerformMove(randomLegalMove)))
	}
}

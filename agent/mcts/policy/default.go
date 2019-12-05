package policy

import (
	"math/rand"
	"time"

	"github.com/techniboy/kalahgo/game"

	"github.com/techniboy/kalahgo/agent/mcts/graph"
)

func McdpSimuilate(n *graph.Node) *game.MancalaEnv {
	for !n.IsTerminal() {
		rand.Seed(time.Now().Unix())
		randomLegalMove := n.State.LegalMoves()[rand.Intn(len(n.State.LegalMoves()))]
		n.Update(n.State.PerformMove(randomLegalMove))
	}
	return n.State
}

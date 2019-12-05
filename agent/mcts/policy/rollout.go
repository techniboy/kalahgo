package policy

import (
	"log"

	"github.com/techniboy/kalahgo/agent/mcts/graph"
	"github.com/techniboy/kalahgo/game"
)

func McrpBackpropagate(root *graph.Node, state *game.MancalaEnv) {
	node := root
	for node != nil {
		var side *game.Side
		if node.Parent != nil {
			side = node.Parent.State.SideToMove
		} else {
			side = node.State.SideToMove
		}
		endGameReward, err := state.ComputeEndGameReward(side)
		if err != nil {
			log.Panic(err)
		}
		node.Update(endGameReward)
		node = node.Parent
	}
}

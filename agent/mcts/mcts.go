package mcts

import (
	"log"
	"time"

	"github.com/techniboy/kalahgo/agent/mcts/graph"
	"github.com/techniboy/kalahgo/agent/mcts/policy"
	"github.com/techniboy/kalahgo/game"
)

type MCTS struct {
	RunTime time.Duration
}

func NewMCTS(runTime time.Duration) *MCTS {
	mcts := new(MCTS)
	mcts.RunTime = runTime
	return mcts
}

func (mcts *MCTS) Search(state *game.MancalaEnv) *game.Move {
	gameStateRoot := graph.NewNode(state.Clone(), nil, nil)
	gamesPlayed := 0
	startTime := time.Now()
	for time.Now().Sub(startTime) < mcts.RunTime {
		node := policy.MctpSelect(gameStateRoot)
		policy.McdpSimuilate(node)
		graph.Backpropagate(node)
		gamesPlayed++
	}
	maxChild, err := graph.SelectMaxChild(gameStateRoot)
	if err != nil {
		log.Panic(err)
	}
	return maxChild.Move
}

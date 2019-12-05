package mcts

import (
	"log"
	"sync"
	"time"

	"github.com/techniboy/kalahgo/agent/mcts/graph"
	"github.com/techniboy/kalahgo/agent/mcts/policy"
	"github.com/techniboy/kalahgo/game"
)

type MCTS struct {
	RunTime     float64
	Root        *graph.Node
	mutex       sync.Mutex
	GamesPlayed int
}

func NewMCTS() *MCTS {
	mcts := new(MCTS)
	mcts.Root = graph.NewNode(game.NewMancalaEnv(), nil, nil)
	return mcts
}

func (mcts *MCTS) Search() {
	for !mcts.Root.State.IsGameOver() {
		node := policy.MctpSelect(mcts.Root)
		finalState := policy.McdpSimuilate(node)
		mcts.mutex.Lock()
		policy.McrpBackpropagate(node, finalState)
		mcts.mutex.Unlock()
		mcts.GamesPlayed++
	}
}

func (mcts *MCTS) BestMove() *game.Move {
	legalMoves := mcts.Root.State.LegalMoves()
	if len(legalMoves) == 1 {
		return legalMoves[0]
	}
	for {
		mcts.mutex.Lock()
		if mcts.Root.Visits < 50 {
			mcts.mutex.Unlock()
			time.Sleep(1 * time.Second)
		} else {
			selectedNode, err := graph.SelectRobustChild(mcts.Root)
			if err != nil {
				log.Panic(err)
			}
			mcts.mutex.Unlock()
			return selectedNode.Move
		}
	}
}

func (mcts *MCTS) PerformMove(moveIndex int) {
	mcts.mutex.Lock()
	if mcts.Root == nil {
		mcts.mutex.Unlock()
		return
	}
	for _, child := range mcts.Root.Children {
		if child.Move.Index == moveIndex {
			mcts.Root = child
			mcts.Root.Parent = nil
			mcts.mutex.Unlock()
			return
		}
	}
	for _, unexploredMove := range mcts.Root.UnexploredMoves {
		if unexploredMove.Index == moveIndex {
			move, err := game.NewMove(mcts.Root.State.SideToMove, moveIndex)
			if err != nil {
				log.Panic(err)
			}
			mcts.Root.State.PerformMove(move)
			mcts.Root = graph.NewNode(mcts.Root.State, unexploredMove, nil)
			mcts.mutex.Unlock()
			return
		}
	}
	mcts.mutex.Unlock()
	log.Panic("No child with the same move was found")
}

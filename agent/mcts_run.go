package agent

import (
	"github.com/techniboy/kalahgo/agent/mcts"
	"github.com/techniboy/kalahgo/game"
	"github.com/techniboy/kalahgo/protocol"
)

func RunGameMCTS(gameConn *protocol.GameConnection) {
	mcts := mcts.NewMCTS()
	state := game.NewMancalaEnv()
	go mcts.Search()
	for {
		msg := protocol.ReadMsg(gameConn)
		msgType, err := protocol.GetMsgType(msg)
		if err != nil {
			panic(err)
		}
		messageType := protocol.NewMsgType()
		// start playing the game
		if msgType == messageType.START {
			first, err := protocol.InterpretStartMsg(msg)
			if err != nil {
				panic(err)
			}
			if first {
				move := mcts.BestMove()
				protocol.SendMsg(gameConn, protocol.CreateMoveMsg(move.Index))
			} else {
				state.OurSide = state.OurSide.Opposite()
			}
		} else if msgType == messageType.STATE {
			moveTurn, err := protocol.InterpretStateMsg(msg)
			if err != nil {
				panic(err)
			}
			moveToPerform, err := game.NewMove(state.SideToMove, moveTurn.Move)
			if err != nil {
				panic(err)
			}
			state.PerformMove(moveToPerform)
			mcts.PerformMove(moveToPerform.Index)
			if !moveTurn.End {
				if moveTurn.Again {
					move := mcts.BestMove()
					if move.Index == 0 {
						protocol.SendMsg(gameConn, protocol.CreateSwapMsg())
					} else {
						protocol.SendMsg(gameConn, protocol.CreateMoveMsg(move.Index))
					}
				}
			}
		} else if msgType == messageType.END {
			break
		} else {
			panic("Invalid message type:" + msgType)
		}
	}
}

package agent

import (
	"log"

	"github.com/techniboy/kalahgo/agent/mcts"
	"github.com/techniboy/kalahgo/game"
	"github.com/techniboy/kalahgo/protocol"
)

func RunGameMCTS(state *game.MancalaEnv, mcts *mcts.MCTS) {
	log.Println("starting game...")
	gameConn, err := protocol.NewGameConnection("127.0.0.1", "12345")
	if err != nil {
		log.Panic(err)
	}
	for {
		msg := protocol.ReadMsg(gameConn)
		msgType, err := protocol.GetMsgType(msg)
		if err != nil {
			log.Panic(err)
		}
		messageType := protocol.NewMsgType()
		// start playing the game
		if msgType == messageType.START {
			first, err := protocol.InterpretStartMsg(msg)
			if err != nil {
				log.Panic(err)
			}
			if first {
				move := mcts.Search(state)
				protocol.SendMsg(gameConn, protocol.CreateMoveMsg(move.Index))
			} else {
				state.OurSide = state.OurSide.Opposite()
			}
		} else if msgType == messageType.STATE {
			moveTurn, err := protocol.InterpretStateMsg(msg)
			if err != nil {
				log.Panic(err)
			}
			if moveTurn.Move == 0 {
				state.OurSide = state.OurSide.Opposite()
			}

			moveToPerform, err := game.NewMove(state.SideToMove, moveTurn.Move)
			if err != nil {
				log.Panic(err)
			}

			state.PerformMove(moveToPerform)
			if !moveTurn.End {
				if moveTurn.Again {
					move := mcts.Search(state)
					if move.Index == 0 {
						protocol.SendMsg(gameConn, protocol.CreateSwapMsg())
					} else {
						protocol.SendMsg(gameConn, protocol.CreateMoveMsg(move.Index))
					}
				}
			}
		} else if msgType == messageType.END {
			log.Println("Game engine ended the game")
			break
		} else {
			log.Println("Invalid message type:" + msgType)
		}
	}
}

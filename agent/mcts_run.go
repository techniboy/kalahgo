package agent

import (
	"log"
	"math/rand"
	"time"

	"github.com/techniboy/kalahgo/agent/mcts"
	"github.com/techniboy/kalahgo/game"
	"github.com/techniboy/kalahgo/protocol"
)

func RunGameMCTS() {
	log.Println("starting game...")
	rand.Seed(time.Now().Unix())
	gameConn, err := protocol.NewGameConnection("127.0.0.1", "12340")
	if err != nil {
		log.Panic(err)
	}
	mcts := mcts.NewMCTS()
	state := game.NewMancalaEnv()
	go mcts.Search()
	// go mcts.Search()
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
				move := mcts.BestMove()
				protocol.SendMsg(gameConn, protocol.CreateMoveMsg(move.Index))
			} else {
				state.OurSide = state.OurSide.Opposite()
			}
		} else if msgType == messageType.STATE {
			moveTurn, err := protocol.InterpretStateMsg(msg)
			if err != nil {
				log.Panic(err)
			}
			moveToPerform, err := game.NewMove(state.SideToMove, moveTurn.Move)
			if err != nil {
				log.Panic(err)
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
			log.Printf("total games played = %d", mcts.GamesPlayed)
			log.Println("Game engine ended the game")
			break
		} else {
			log.Println("Invalid message type:" + msgType)
		}
	}
}

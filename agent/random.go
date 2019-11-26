package agent

import (
	"log"
	"math/rand"
	"time"

	"github.com/techniboy/kalahgo/game"
	"github.com/techniboy/kalahgo/protocol"
)

func mindlessRandom(state *game.MancalaEnv) *game.Move {
	legalMoves := state.LegalMoves()
	return legalMoves[rand.Intn(len(legalMoves))]
}

func RunGame(state *game.MancalaEnv) {
	log.Println("starting game...")
	rand.Seed(time.Now().Unix())
	gameConn, err := protocol.NewGameConnection("127.0.0.1", "12340")
	if err != nil {
		log.Panic(err)
	}
	ourSide := game.NewSide(game.SideSouth)
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
				move := mindlessRandom(state)
				protocol.SendMsg(gameConn, protocol.CreateMoveMsg(move.Index))
			} else {
				ourSide = ourSide.Opposite()
			}
		} else if msgType == messageType.STATE {
			moveTurn, err := protocol.InterpretStateMsg(msg)
			if err != nil {
				log.Panic(err)
			}
			if moveTurn.Move == 0 {
				ourSide = ourSide.Opposite()
			}

			moveToPerform, err := game.NewMove(state.SideToMove, moveTurn.Move)
			if err != nil {
				log.Panic(err)
			}

			state.PerformMove(moveToPerform)
			if !moveTurn.End {
				if moveTurn.Again {
					move := mindlessRandom(state)
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

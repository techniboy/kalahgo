package agent

import (
	"fmt"
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
	rand.Seed(time.Now().Unix())
	ourSide := game.NewSide(game.SideSouth)
	for {
		msg := protocol.ReadMsg()
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
				move := mindlessRandom(state)
				protocol.SendMsg(protocol.CreateMoveMsg(move.Index))
			} else {
				ourSide = ourSide.Opposite()
			}
		} else if msgType == messageType.STATE {
			moveTurn, err := protocol.InterpretStateMsg(msg)
			if err != nil {
				panic(err)
			}
			if moveTurn.Move == 0 {
				ourSide = ourSide.Opposite()
			}

			moveToPerform, err := game.NewMove(state.SideToMove, moveTurn.Move)
			if err != nil {
				panic(err)
			}

			state.PerformMove(moveToPerform)
			if !moveTurn.End {
				if moveTurn.Again {
					move := mindlessRandom(state)
					if move.Index == 0 {
						protocol.SendMsg(protocol.CreateSwapMsg())
					} else {
						protocol.SendMsg(protocol.CreateMoveMsg(move.Index))
					}
				}
			}
		} else if msgType == messageType.END {
			break
		} else {
			fmt.Println("Invalid message type:" + msgType)
		}
	}
}

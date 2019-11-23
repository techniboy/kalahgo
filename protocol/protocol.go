package protocol

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Send a message to the game engine
func SendMsg(msg string) {
	fmt.Println(msg)
}

// Receives a message from the game engine. Messages are terminated by
// a '\n' character.
func ReadMsg() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return text
}

// Creates a move message
func CreateMoveMsg(hole int) string {
	return "MOVE;" + string(hole)
}

func GetMsgType(msg string) string {
	if strings.HasPrefix(msg, "START;") {
		return START
	} else if strings.HasPrefix(msg, "CHANGE;") {
		return STATE
	} else if strings.HasPrefix(msg, "END\n") {
		return END
	} else {
		return "invalidMessageError: could not determine message type."
	}

}

// Interprets a "state_change" message. Should be called if
// GetMsgType(msg) returns MsgType.STATE
func InterpretStateMsg(msg string) MoveTurn {
	moveTurn := MoveTurn{}

	if msg[len(msg)-1] != '\n' {
		fmt.Println("invalidMessageError: message not terminated with 0x0A character.")
	}

	msgParts := strings.SplitN(msg, ":", 4)
	if len(msgParts) != 4 {
		fmt.Println("invalidMessageError: missing arguments.")
	}

	// msgParts[0] is "CHANGE"
	// 1st argument: the move (or swap)
	if msgParts[1] == "SWAP" {
		moveTurn.SetMove(0)
	} else {
		move, err := strconv.Atoi(msgParts[1])
		if err != nil {
			panic(err)
		}
		moveTurn.SetMove(move)
	}

	// 3rd argument: whose turn
	moveTurn.SetEnd(false)
	if msgParts[3] == "YOU\n" {
		moveTurn.SetAgain(true)
	} else if msgParts[3] == "OPP\n" {
		moveTurn.SetAgain(false)
	} else if msgParts[3] == "END\n" {
		moveTurn.SetEnd(true)
		moveTurn.SetAgain(false)
	} else {
		fmt.Printf("invalidMessageError: illegal value for turn parameter")
	}
	return moveTurn
}

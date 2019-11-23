package protocol

import (
	"bufio"
	"fmt"
	"os"
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
		return "invalidMessageError: could not determine message type"
	}

}

func InterpretStateMsg(msg string) MoveTurn {

}

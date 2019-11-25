package protocol

import (
	"bufio"
	"errors"
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

func CreateSwapMsg() string {
	return "SWAP"
}

func GetMsgType(msg string) (string, error) {
	msgType := NewMsgType()
	if strings.HasPrefix(msg, "START;") {
		return msgType.START, nil
	} else if strings.HasPrefix(msg, "CHANGE;") {
		return msgType.STATE, nil
	} else if strings.HasPrefix(msg, "END\n") {
		return msgType.END, nil
	} else {
		return "", errors.New("invalidMessageError: could not determine message type")
	}

}

// Interprets a "new_match" message. Should be called if
// GettMsgType(msg) returns MsgType.START
func InterpretStartMsg(msg string) (bool, error) {
	if msg[len(msg)-1] != '\n' {
		return false, errors.New("invalidMessageError: message not terminated with 0x0A character")
	}

	// Message are of the form START:<POSITION> \n
	position := msg[6 : len(msg)-1]
	if position == "South" {
		return true, nil
	} else if position == "North" {
		return false, nil
	} else {
		return false, errors.New("invalidMessageError: illegal position parameter")
	}
	// IMPLEMENT ERROR HANDLING OR THIS WILL BREAK
	return false, nil
}

// Interprets a "state_change" message. Should be called if
// GetMsgType(msg) returns MsgType.STATE
func InterpretStateMsg(msg string) (*MoveTurn, error) {
	moveTurn := new(MoveTurn)

	if msg[len(msg)-1] != '\n' {
		return nil, errors.New("invalidMessageError: message not terminated with 0x0A character")
	}

	msgParts := strings.SplitN(msg, ":", 4)
	if len(msgParts) != 4 {
		return nil, errors.New("invalidMessageError: missing arguments")
	}

	// msgParts[0] is "CHANGE"
	// 1st argument: the move (or swap)
	if msgParts[1] == "SWAP" {
		moveTurn.Move = 0
	} else {
		move, err := strconv.Atoi(msgParts[1])
		if err != nil {
			panic(err)
		}
		moveTurn.Move = move
	}

	// 3rd argument: whose turn
	moveTurn.End = false
	if msgParts[3] == "YOU\n" {
		moveTurn.Again = true
	} else if msgParts[3] == "OPP\n" {
		moveTurn.Again = false
	} else if msgParts[3] == "END\n" {
		moveTurn.End = true
		moveTurn.Again = false
	} else {
		return nil, errors.New("invalidMessageError: illegal value for turn parameter")
	}
	return moveTurn, nil
}

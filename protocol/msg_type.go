package protocol

type MsgType struct {
	START string
	STATE string
	END   string
}

func NewMsgType() *MsgType {
	m := new(MsgType)
	// message announcing the start of the game ("new_match" message)
	m.START = "new_match"
	// message describing a move or a swap ("state_change" message)
	m.STATE = "state_change"
	// message informing about the end of the game ("game_over" message)
	m.END = "game_over"
	return m
}

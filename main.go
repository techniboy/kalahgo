package main

import (
	"github.com/techniboy/kalahgo/agent"
	"github.com/techniboy/kalahgo/protocol"
)

func main() {
	gameConn, err := protocol.NewGameConnection("127.0.0.1", "12340")
	if err != nil {
		panic(err)
	}
	agent.RunGameMCTS(gameConn)
}

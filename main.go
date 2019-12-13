package main

import (
	"log"
	"os"

	"github.com/techniboy/kalahgo/agent"
	"github.com/techniboy/kalahgo/protocol"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)

	gameConn, err := protocol.NewGameConnection("127.0.0.1", "12340")
	if err != nil {
		log.Panic(err)
	}
	agent.RunGameMCTS(gameConn)
}

package main

import (
	"log"
	"os"

	"github.com/techniboy/kalahgo/agent/mcts"

	"github.com/techniboy/kalahgo/agent"
	"github.com/techniboy/kalahgo/game"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)

	state := game.NewMancalaEnv()
	// agent.RunGameRandom(state)
	agent.RunGameMCTS(state, mcts.NewMCTS(1))
}

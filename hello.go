package main

import (
	"log"
	"os"

	"github.com/techniboy/kalahgo/agent"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)

	// agent.RunGameRandom(state)
	agent.RunGameMCTS()
}

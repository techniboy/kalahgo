package main

import (
	"github.com/techniboy/kalahgo/agent"
	"github.com/techniboy/kalahgo/game"
)

func main() {
	state := game.NewMancalaEnv()
	agent.RunGame(state)
}

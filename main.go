package main

import (
	"go-tetris/entities"
	"go-tetris/systems"
	"time"
)

func gameLoop(fn func()) {
	frameTime := (1.0 / 60.0) * float64(time.Second)
	for {
		prevTime := time.Now()
		fn()
		delta := time.Since(prevTime)
		pauseTime := frameTime - float64(delta)
		if pauseTime > 0 {
			time.Sleep(time.Duration(pauseTime))
		} else {
			time.Sleep(0)
		}
	}
}

func main() {
	game := entities.NewGame()
	renderSystem := systems.NewRenderSystem()
	gamePlaySystem := systems.NewGamePlaySystem()

	gamePlaySystem.Init(&game)
	renderSystem.Init(&game)
	defer gamePlaySystem.Close()
	defer renderSystem.Close()

	gameLoop(func() {
		gamePlaySystem.Tick()
		renderSystem.Tick()
	})
}

package main

import (
	"fmt"
	"go-tetris/entities"
	"go-tetris/systems"
	"time"
)

func gameLoop(fn func(dt time.Duration) error) error {
	frameTime := (1.0 / 60.0) * float64(time.Second)
	prevTime := time.Now()
	for {
		dt := time.Since(prevTime)
		prevTime := time.Now()
		err := fn(dt)
		if err != nil {
			return err
		}
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

	var err error
	defer func() {
		fmt.Println(err)
	}()
	defer gamePlaySystem.Close()
	defer renderSystem.Close()

	err = gameLoop(func(dt time.Duration) error {
		if err := gamePlaySystem.Tick(dt); err != nil {
			return err
		}
		if err := renderSystem.Tick(dt); err != nil {
			return err
		}
		return nil
	})
}

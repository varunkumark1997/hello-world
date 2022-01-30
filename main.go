package main

import (
	"github.com/varunkumark1997/hello-world/game"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	game.InitPlayArea()
	defer sdl.Quit()
	defer game.Window.Destroy()
	defer game.Surface.Free()
	game.InitSnake()
	game.InitFood()
	running := true
	for running {
		game.MoveSnake()
		game.EatFood()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}

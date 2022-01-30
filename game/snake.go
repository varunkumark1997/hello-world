package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	SnakeSpeed = 0.4
	Right      = "right"
	Left       = "left"
	Down       = "down"
	Up         = "up"
)

type SnakeChar struct {
	x, y      float64
	direction string
	length    int32
}

var (
	snake    sdl.Rect
	newSnake SnakeChar
)

func InitSnake() {
	xPos := 400
	yPos := 400
	newSnake = SnakeChar{
		float64(xPos),
		float64(yPos),
		"up",
		10,
	}
	snake = sdl.Rect{X: int32(newSnake.x), Y: int32(newSnake.y), W: 10, H: newSnake.length}
	updateScreen(&snake)
}

func MoveSnake() {
	Clear(&snake)
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_RIGHT] == 1 && newSnake.direction != Left {
		newSnake.direction = Right
	}
	if keys[sdl.SCANCODE_DOWN] == 1 && newSnake.direction != Up {
		newSnake.direction = Down
	}
	if keys[sdl.SCANCODE_LEFT] == 1 && newSnake.direction != Right {
		newSnake.direction = Left
	}
	if keys[sdl.SCANCODE_UP] == 1 && newSnake.direction != Down {
		newSnake.direction = Up
	}

	switch newSnake.direction {
	case Up:
		newSnake.y -= SnakeSpeed
	case Left:
		newSnake.x -= SnakeSpeed
	case Right:
		newSnake.x += SnakeSpeed
	case Down:
		newSnake.y += SnakeSpeed
	}
	snake.Y = int32(newSnake.y)
	snake.X = int32(newSnake.x)
	updateScreen(&snake)
}

func increaseSnake() {
	newSnake.length += 1
	switch newSnake.direction {
	case Up:
		snake.H += 1
	case Down:
		snake.H += 1
	case Left:
		snake.H += 1
	case Right:
		snake.H += 1
	}
	updateScreen(&snake)
}

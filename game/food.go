package game

import (
	"github.com/varunkumark1997/hello-world/util"
	"github.com/veandco/go-sdl2/sdl"
)

var Food sdl.Rect

func InitFood() {
	xPos := util.RandomPos(3, boardLen-3)
	yPos := util.RandomPos(3, boardWdth-3)
	Food = sdl.Rect{X: int32(xPos), Y: int32(yPos), W: 10, H: 10}
	updateScreen(&Food)
}

func NewFood() {
	Food.X = int32(util.RandomPos(3, boardLen-3))
	Food.Y = int32(util.RandomPos(3, boardWdth-3))
	updateScreen(&Food)
}

func EatFood() {
	sx := snake.X
	sy := snake.Y
	fx := Food.X
	fy := Food.Y
	if (sx-fx) < 2 && (sy-fy) < 2 {
		NewFood()
		increaseSnake()
	}
}

package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

var (
	err       error
	Surface   *sdl.Surface
	Window    *sdl.Window
	boardLen  = 800
	boardWdth = 800
)

func InitPlayArea() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	Window, err = sdl.CreateWindow("Snake Game", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(boardLen), int32(boardWdth), sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}

	Surface, err = Window.GetSurface()
	if err != nil {
		panic(err)
	}
	Clear(nil)
}

func updateScreen(rect *sdl.Rect) {
	Surface.FillRect(rect, 0xffffffff)
	Window.UpdateSurface()
}

func Clear(rect *sdl.Rect) {
	Surface.FillRect(rect, 0)
	Window.UpdateSurface()
}

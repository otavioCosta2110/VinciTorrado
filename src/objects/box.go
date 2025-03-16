package objects

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Box struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
	Color  rl.Color
}

func NewBox(x, y, width, height int32, color rl.Color) *Box {
	return &Box{
		X:      x,
		Y:      y,
		Width:  50,
		Height: 50,
		Color:  color,
	}
}

func (b *Box) Draw() {
	rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, b.Color)
}

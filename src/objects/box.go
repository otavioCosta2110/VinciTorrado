package objects

import (
	"otaviocosta2110/getTheBlueBlocks/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Box struct {
	Object system.Object
	Color  rl.Color
}

func NewBox(x, y, width, height int32, color rl.Color) *Box {
	return &Box{
		Object: system.Object{
			X:          x,
			Y:          y,
			Width:      50,
			Height:     50,
			KnockbackX: 0,
			KnockbackY: 0,
		},
		Color: color,
	}
}

func (b *Box) Draw() {
	rl.DrawRectangle(b.Object.X-b.Object.Width/2, b.Object.Y-b.Object.Height/2, b.Object.Width, b.Object.Height, b.Color)
}

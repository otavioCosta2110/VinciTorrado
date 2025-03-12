package enemy

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Enemy struct {
	X       int32
	Y       int32
	Speed   int32
	Width   int32
	Height  int32
	Sprite  rl.Texture2D
	Scale   int32
	Flipped bool
}

func NewEnemy(x, y, speed, width, height, scale int32, sprite rl.Texture2D) *Enemy {
	return &Enemy{
		X:      x,
		Y:      y,
		Width:  width,
		Speed:  speed,
		Height: height,
		Scale: scale,
		Sprite: sprite,
		Flipped: false,
	}
}

func (e *Enemy) DrawEnemy() {
	var width int32
	if e.Flipped {
		width = -e.Width
	} else {
		width = e.Width
	}

	sourceRec := rl.NewRectangle(0, 0, float32(width), float32(e.Height))
	destinationRec := rl.NewRectangle(float32(e.X), float32(e.Y), float32(e.Width)*float32(e.Scale), float32(e.Height)*float32(e.Scale))
	origin := rl.NewVector2(float32(e.Width)*float32(e.Scale)/2, float32(e.Height)*float32(e.Scale)/2)

	rl.DrawTexturePro(e.Sprite, sourceRec, destinationRec, origin, 0.0, rl.White)
}

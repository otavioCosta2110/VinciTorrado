package enemy

import (
	"otaviocosta2110/getTheBlueBlocks/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Enemy struct {
	Object  system.Object
	Speed   int32
	Sprite  rl.Texture2D
	Scale   int32
	Flipped bool
}

func NewEnemy(x, y, speed, width, height, scale int32, sprite rl.Texture2D) *Enemy {
	return &Enemy{
    Object: system.Object{
      X:       x,
      Y:       y,
      Width:   width * scale,
      Height:  height * scale,
    },
		Speed:   speed,
		Scale:   scale,
		Sprite:  sprite,
		Flipped: false,
	}
}

func (e *Enemy) DrawEnemy() {
	var width float32 = 32
	if e.Flipped {
		width = -float32(width)
	}

	sourceRec := rl.NewRectangle(
		float32(e.Object.Width)*32,
		float32(e.Object.Height)*32,
		width,
		float32(32),
	)

	destinationRec := rl.NewRectangle(
		float32(e.Object.X),
		float32(e.Object.Y),
		float32(e.Object.Width),
		float32(e.Object.Height),
	)

	origin := rl.NewVector2(
		destinationRec.Width/2,
		destinationRec.Height/2,
	)

	rl.DrawTexturePro(e.Sprite, sourceRec, destinationRec, origin, 0.0, rl.White)
}

package enemy

import (
	"math/rand"
	"otaviocosta2110/getTheBlueBlocks/src/physics"
	"otaviocosta2110/getTheBlueBlocks/src/sprites"
	"otaviocosta2110/getTheBlueBlocks/src/system"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	animationDelay int32 = 300
)

type Enemy struct {
	Object  system.Object
	Speed   int32
	Scale   int32
	Flipped bool
}

func NewEnemy(x, y, speed, width, height, scale int32, sprite sprites.Sprite) *Enemy {
	return &Enemy{
		Object: system.Object{
			X:             x,
			Y:             y,
			Width:         width * scale / 2,
			Height:        height * scale,
			KnockbackX:    0,
			KnockbackY:    0,
			FrameY:        0,
			FrameX:        0,
			LastFrameTime: time.Now(),
			Sprite: sprites.Sprite{
				SpriteWidth:  width,
				SpriteHeight: height,
				Texture:      sprite.Texture,
			},
		},
		Speed:   speed,
		Scale:   scale,
		Flipped: false,
	}
}

func (e *Enemy) DrawEnemy() {
	var width float32 = float32(e.Object.Sprite.SpriteWidth)
	if e.Flipped {
		width = -float32(width)
	}

	sourceRec := rl.NewRectangle(
		float32(e.Object.FrameX)*float32(e.Object.Sprite.SpriteWidth),
		float32(e.Object.FrameY)*float32(e.Object.Sprite.SpriteWidth),
		width,
		float32(e.Object.Sprite.SpriteHeight),
	)

	destinationRec := rl.NewRectangle(
		float32(e.Object.X),
		float32(e.Object.Y),
		float32(e.Object.Sprite.SpriteWidth)*float32(e.Scale),
		float32(e.Object.Sprite.SpriteHeight)*float32(e.Scale),
	)

	origin := rl.NewVector2(
		destinationRec.Width/2,
		destinationRec.Height/2,
	)

	rl.DrawTexturePro(e.Object.Sprite.Texture, sourceRec, destinationRec, origin, 0.0, rl.White)
}

func (e *Enemy) CheckAtk(player system.Object) bool {
	// var isAttacking = false

	// botei essas vars pra ca pra fazer a caixa de colisao aparecer sempre
	punchX := e.Object.X
	punchY := e.Object.Y - e.Object.Height/3

	punchWidth := e.Object.Width
	punchHeight := e.Object.Height / 2

	if e.Flipped {
		punchX -= punchWidth + punchWidth/2 //esquerda
	} else {
		punchX += punchWidth / 2 //direita, n sei pq ta assim
	}

	// cor da colisão do soco (debug)
	rl.DrawRectangle(punchX, punchY, punchWidth, punchHeight, rl.Red)

	punchObject := system.Object{
		X:      punchX,
		Y:      punchY,
		Width:  punchWidth,
		Height: punchHeight,
	}

	if physics.CheckCollision(punchObject, player) {
		// isAttacking = true

    framex := rand.Intn(2)
    println(framex)

    e.Object.UpdateAnimation(50, []int{framex}, []int{1})

		return true
	}
  e.Object.UpdateAnimation(300, []int{0, 1}, []int{0, 0})

	return false
}

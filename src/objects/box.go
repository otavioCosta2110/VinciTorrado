package objects

import (
	"otaviocosta2110/getTheBlueBlocks/src/physics"
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
			Width:      width,
			Height:     height,
			KnockbackX: 0,
			KnockbackY: 0,
		},
		Color: color,
	}
}

func (b *Box) Draw() {
	rl.DrawRectangle(b.Object.X-b.Object.Width/2, b.Object.Y-b.Object.Height/2, b.Object.Width, b.Object.Height, b.Color)
}

func (b *Box) Update(colliders []system.Object, screenHeight int32) {
	b.Object.X += b.Object.KnockbackX
	b.Object.Y += b.Object.KnockbackY

	b.Object.KnockbackX = int32(float64(b.Object.KnockbackX) * 0.85) // amortecimento, baita palavra

	b.Object.KnockbackY += 2

	groundLevel := screenHeight - 100
	if b.Object.Y+b.Object.Height/2 > groundLevel {
		b.Object.Y = groundLevel - b.Object.Height/2
		b.Object.KnockbackY = -b.Object.KnockbackY / 2

		if abs(b.Object.KnockbackY) < 5 {
			b.Object.KnockbackY = 0
		}

		if abs(b.Object.KnockbackX) < 3 {
			b.Object.KnockbackX = 0
		}
	}

	if b.Object.X-b.Object.Width/2 < 0 {
		b.Object.X = b.Object.Width / 2
		b.Object.KnockbackX = 0
	}
	if b.Object.X+b.Object.Width/2 > 1280 {
		b.Object.X = 1280 - b.Object.Width/2
		b.Object.KnockbackX = 0
	}

	for _, obj := range colliders {
		if physics.CheckCollision(b.Object, obj) {
			physics.ResolveCollision(&b.Object, &obj)
		}
	}
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

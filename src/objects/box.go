package objects

import (
	"otaviocosta2110/getTheBlueBlocks/src/physics"
	"otaviocosta2110/getTheBlueBlocks/src/screen"
	"otaviocosta2110/getTheBlueBlocks/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Box struct {
	Object    system.Object
	Color     rl.Color
	OriginalY int32
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

func (b *Box) Update(colliders []system.Object, s *screen.Screen) {
	b.Object.X += b.Object.KnockbackX
	b.Object.Y += b.Object.KnockbackY

	b.Object.KnockbackX = int32(float64(b.Object.KnockbackX) * 0.85)
	b.Object.KnockbackY = int32(float64(b.Object.KnockbackY) * 0.50)
	if b.Object.Y < s.ScenaryHeight {
		b.Object.KnockbackY += 2
	}

	groundLevel := s.ScenaryHeight + 320
	if b.Object.Y+b.Object.Height/2 > groundLevel {
		b.Object.Y = groundLevel - b.Object.Height/2
		b.Object.KnockbackY = -b.Object.KnockbackY * 3 / 5
		b.Object.KnockbackX = b.Object.KnockbackX * 4 / 5
	}

	if b.Object.X-b.Object.Width/2 < 0 {
		b.Object.X = b.Object.Width / 2
		b.Object.KnockbackX = -b.Object.KnockbackX * 6 / 10
		b.Object.KnockbackY = -abs(b.Object.KnockbackX) / 3
	}
	if b.Object.X+b.Object.Width/2 > s.ScenaryWidth {
		b.Object.X = s.ScenaryWidth - b.Object.Width/2
		b.Object.KnockbackX = -b.Object.KnockbackX * 6 / 10
		b.Object.KnockbackY = -abs(b.Object.KnockbackX) / 3
	}

	// Stop very small movements
	if abs(b.Object.KnockbackX) < 2 {
		b.Object.KnockbackX = 0
	}
	if abs(b.Object.KnockbackY) < 2 {
		b.Object.KnockbackY = 0
	}

	for _, obj := range colliders {
		if physics.CheckCollision(b.Object, obj) {
			tempObj := obj
			physics.ResolveCollision(&b.Object, &tempObj)
		}
	}
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

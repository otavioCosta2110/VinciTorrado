package objects

import (
	"otaviocosta2110/vincitorrado/src/enemy"
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Box struct {
	Object    system.Object
	Texture   rl.Texture2D
	OriginalY int32
}

func NewBox(x, y, width, height int32) *Box {
	texture := rl.LoadTexture("assets/props/box.png")

	return &Box{
		Object: system.Object{
			X:          x,
			Y:          y,
			Width:      width,
			Height:     height,
			KnockbackX: 0,
			KnockbackY: 0,
		},
		Texture: texture,
	}
}

func (b *Box) IsKicked() bool{
	return false
}

func (b *Box) HandleKick(_ *[]*equipment.Equipment, playerObj system.Object) {
		// knockbackMultiplier := int32(3)
		// b.Object.KnockbackX = 150 * knockbackMultiplier
		knockbackMultiplier := int32(10)
    b.Object.KnockbackX = 10 * knockbackMultiplier
    // b.Object.KnockbackY = -50 * knockbackMultiplier
		// } else {
		// 	b.Object.KnockbackX = kickPower * knockbackMultiplier
		// }
		// b.Object.KnockbackY = 0
}

func (b *Box) Draw() {
	rl.DrawTexturePro(
		b.Texture,
		rl.NewRectangle(0, 0, float32(b.Texture.Width), float32(b.Texture.Height)),
		rl.NewRectangle(
			float32(b.Object.X),
			float32(b.Object.Y),
			float32(b.Object.Width),
			float32(b.Object.Height),
		),
		rl.NewVector2(float32(b.Object.Width)/2, float32(b.Object.Height)/2),
		0,
		rl.White,
	)
}

func (b *Box) Update(colliders []system.Object, s *screen.Screen, em *enemy.EnemyManager) {
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
		b.Object.KnockbackX = -b.Object.KnockbackX * 8 / 10
		b.Object.KnockbackY = -abs(b.Object.KnockbackX) / 3
	}
	if b.Object.X+b.Object.Width/2 > s.ScenaryWidth {
		b.Object.X = s.ScenaryWidth - b.Object.Width/2
		b.Object.KnockbackX = -b.Object.KnockbackX * 8 / 10
		b.Object.KnockbackY = -abs(b.Object.KnockbackX) / 3
	}

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

	if em != nil && (abs(b.Object.KnockbackX) > 5 || abs(b.Object.KnockbackY) > 5) {
		for _, e := range em.ActiveEnemies {
			if physics.CheckCollision(b.Object, e.Object) {
				e.TakeDamageFromBox(b.Object)
			}
		}
	}
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func (b *Box) GetObject() system.Object {
	return b.Object
}

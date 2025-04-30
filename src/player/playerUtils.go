package player

import (
	"otaviocosta2110/vincitorrado/src/enemy"
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/system"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	invencibilityDuration = 1000
)

func (p *Player) Update(em *enemy.EnemyManager, screen screen.Screen) {
	physics.TakeKnockback(&p.Object)

	if p.Object.KnockbackX == 0 || p.Object.KnockbackY == 0 {
		p.CheckMovement(screen)
	}

	for _, enemy := range em.Enemies {
		enemyObject := enemy.GetObject()

		if p.CheckAtk(enemyObject) {
			enemy.TakeDamage(p.Damage, p.Object.X, p.Object.Y)
		}
	}
}

func (p *Player) Draw() {
	color := rl.White

	if p.isInvincible(invencibilityDuration) {
		if time.Since(p.LastDamageTaken).Milliseconds()/100%2 == 0 {
			color = rl.Fade(rl.White, 0.3)
		}
	}

	var width float32 = float32(p.Object.Sprite.SpriteWidth)
	if p.Flipped {
		width = -width
	}

	sourceRec := rl.NewRectangle(
		float32(p.Object.FrameX)*float32(p.Object.Sprite.SpriteWidth),
		float32(p.Object.FrameY)*float32(p.Object.Sprite.SpriteHeight),
		width,
		float32(p.Object.Sprite.SpriteHeight),
	)

	destinationRec := rl.NewRectangle(
		float32(p.Object.X),
		float32(p.Object.Y),
		float32(p.Object.Sprite.SpriteWidth)*float32(p.Object.Scale),
		float32(p.Object.Sprite.SpriteHeight)*float32(p.Object.Scale),
	)

	origin := rl.NewVector2(
		destinationRec.Width/2,
		destinationRec.Height/2,
	)

	rl.DrawTexturePro(
		p.Object.Sprite.Texture,
		sourceRec,
		destinationRec,
		origin,
		0.0,
		color,
	)

	if p.Equipment != nil && p.Equipment.IsEquipped && p.HatSprite.Texture.ID != 0 {
		rl.DrawTexturePro(
			p.HatSprite.Texture,
			sourceRec,
			destinationRec,
			origin,
			0.0,
			color,
		)
	}
}

func (p *Player) TakeDamage(damage int32, eObj system.Object) {
	if !p.isInvincible(invencibilityDuration) {
		if p.Health > 1 {
			p.Object.UpdateAnimation(100, []int{0, 1}, []int{3, 3})
			p.Health -= damage
			p.LastDamageTaken = time.Now()
			p.Object.SetKnockback(eObj)
		} else {
			system.GameOverFlag = true
		}
	}
}

func (p *Player) isInvincible(duration int) bool {
	return time.Since(p.LastDamageTaken).Milliseconds() <= int64(duration)
}

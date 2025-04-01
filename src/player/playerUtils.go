package player

import (
	"otaviocosta2110/getTheBlueBlocks/src/physics"
	"otaviocosta2110/getTheBlueBlocks/src/screen"
	"otaviocosta2110/getTheBlueBlocks/src/system"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	invencibilityDuration = 1000
)

func (p *Player) Update(e system.Live, screen screen.Screen) {
	physics.TakeKnockback(&p.Object)

	if p.Object.KnockbackX == 0 || p.Object.KnockbackY == 0 {
		p.CheckMovement(screen)
	}

	enemyObject := e.GetObject()

	if p.CheckAtk(enemyObject) {
		println("atk")
		// reseta a pos do inimigo
		e.SetObject(system.Object{
			X:              0,
			Y:              0,
			Width:          enemyObject.Width,          // Keep existing
			Height:         enemyObject.Height,         // Keep existing
			KnockbackX:     enemyObject.KnockbackX,     // Keep existing
			KnockbackY:     enemyObject.KnockbackY,     // Keep existing
			FrameX:         enemyObject.FrameX,         // Keep existing
			FrameY:         enemyObject.FrameY,         // Keep existing
			LastFrameTime:  enemyObject.LastFrameTime,  // Keep existing
			LastAttackTime: enemyObject.LastAttackTime, // Keep existing
			Sprite:         enemyObject.Sprite,         // Keep existing
			Scale:          enemyObject.Scale,          // Keep existing
		})
		e.TakeDamage(1, p.Object.X, p.Object.Y)
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
		width = -float32(width)
	}

	sourceRec := rl.NewRectangle(
		float32(p.Object.FrameX)*float32(p.Object.Sprite.SpriteWidth),
		float32(p.Object.FrameY)*float32(p.Object.Sprite.SpriteWidth),
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

	rl.DrawTexturePro(p.Object.Sprite.Texture, sourceRec, destinationRec, origin, 0.0, color)

	// faz a caixa vermelha pra ver colisao
	rl.DrawRectangleLines(
		int32(destinationRec.X-origin.X+float32(p.Object.Width)/2),
		int32(destinationRec.Y-origin.Y),
		int32(p.Object.Width),
		int32(p.Object.Height),
		rl.Red,
	)
}

func (p *Player) TakeDamage(damage int32, eX int32, eY int32) {
	if !p.isInvincible(invencibilityDuration) {
		if p.Health > 1 {
			p.Health -= damage
			p.LastDamageTaken = time.Now()
			p.setKnockback(eX, eY)
		} else {
			system.GameOverFlag = true
		}
	}
}

func (p *Player) setKnockback(eX int32, eY int32) {
	knockbackStrengthX := int32(15)
	knockbackStrengthY := int32(10)

	if p.Object.X < eX {
		p.Object.KnockbackX = -knockbackStrengthX
	} else {
		p.Object.KnockbackX = knockbackStrengthX
	}

	if p.Object.Y < eY/2 {
		p.Object.KnockbackY = -knockbackStrengthY
	} else {
		p.Object.KnockbackY = knockbackStrengthY
	}

}

func (p *Player) isInvincible(duration int) bool {
	if time.Since(p.LastDamageTaken).Milliseconds() > int64(duration) {
		return false
	}
	return true
}

package player

import (
	"otaviocosta2110/getTheBlueBlocks/src/system"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	invencibilityDuration = 1000
)

func (p *Player) DrawPlayer() {
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

	frameX := p.Object.FrameX
	// if rl.IsKeyPressed(rl.KeyX) {
    // p.Object.UpdateAnimation(50, []int{0,0}, []int{2,0})
 //    p.Object.FrameY = 2
	// }

	sourceRec := rl.NewRectangle(
		float32(frameX)*float32(p.Object.Sprite.SpriteWidth),
		float32(p.Object.FrameY)*float32(p.Object.Sprite.SpriteWidth),
		width,
		float32(p.Object.Sprite.SpriteHeight),
	)

	destinationRec := rl.NewRectangle(
		float32(p.Object.X),
		float32(p.Object.Y),
		float32(p.Object.Sprite.SpriteWidth)*float32(p.Scale),
		float32(p.Object.Sprite.SpriteHeight)*float32(p.Scale),
	)

	origin := rl.NewVector2(
		destinationRec.Width/2,
		destinationRec.Height/2,
	)

	rl.DrawTexturePro(p.Object.Sprite.Texture, sourceRec, destinationRec, origin, 0.0, color)

	// Desenha a caixa vermelha pra ver colisao
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

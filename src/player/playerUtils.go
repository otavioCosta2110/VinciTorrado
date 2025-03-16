package player

import (
	"otaviocosta2110/getTheBlueBlocks/src/system"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	invencibilityDuration = 1000
)

func (p *Player) UpdateAnimation(animationDelay int, framesX, framesY []int) {
	if time.Since(p.LastFrameTime).Milliseconds() > int64(animationDelay) {
		currentIndex := -1
		for i := range framesX {
			if p.FrameX == int32(framesX[i]) && p.FrameY == int32(framesY[i]) {
				currentIndex = i
				break
			}
		}

		if currentIndex == -1 {
			p.FrameX = int32(framesX[0])
			p.FrameY = int32(framesY[0])
		} else {
			nextIndex := (currentIndex + 1) % len(framesX)
			p.FrameX = int32(framesX[nextIndex])
			p.FrameY = int32(framesY[nextIndex])
		}

		p.LastFrameTime = time.Now()
	}
}

func (p *Player) DrawPlayer() {
	color := rl.White

	if p.isInvincible(invencibilityDuration) {
		if time.Since(p.LastDamageTaken).Milliseconds()/100%2 == 0 {
			color = rl.Fade(rl.White, 0.3)
		}
	}

	var width float32 = float32(p.Sprite.SpriteWidth)
	if p.Flipped {
		width = -float32(width)
	}

	sourceRec := rl.NewRectangle(
		float32(p.FrameX)*float32(p.Sprite.SpriteWidth),
		float32(p.FrameY)*float32(p.Sprite.SpriteWidth),
		width,
		float32(p.Sprite.SpriteHeight),
	)

	destinationRec := rl.NewRectangle(
		float32(p.Object.X),
		float32(p.Object.Y),
		float32(p.Sprite.SpriteWidth)*float32(p.Scale),
		float32(p.Sprite.SpriteHeight)*float32(p.Scale),
	)

	origin := rl.NewVector2(
		destinationRec.Width/2,
		destinationRec.Height/2,
	)

	rl.DrawTexturePro(p.Sprite.Texture, sourceRec, destinationRec, origin, 0.0, color)

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
	}else{
		p.Object.KnockbackY = knockbackStrengthY
  }

}

func (p *Player) isInvincible(duration int) bool {
	if time.Since(p.LastDamageTaken).Milliseconds() > int64(duration) {
		return false
	}
	return true
}

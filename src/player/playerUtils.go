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

	if !p.isInvincible(invencibilityDuration) {
		if time.Since(p.LastDamageTaken).Milliseconds()/100%2 == 0 {
			color = rl.Fade(rl.White, 0.3) 
		}
	}

	var width float32 = 32
	if p.Flipped {
		width = -float32(width)
	}

	sourceRec := rl.NewRectangle(
		float32(p.FrameX)*32,
		float32(p.FrameY)*32,
		width,
		float32(32),
	)

	destinationRec := rl.NewRectangle(
		float32(p.Object.X),
		float32(p.Object.Y),
		float32(p.Object.Width),
		float32(p.Object.Height),
	)

	origin := rl.NewVector2(
		destinationRec.Width/2,
		destinationRec.Height/2,
	)

	rl.DrawTexturePro(p.Sprite, sourceRec, destinationRec, origin, 0.0, color)

	// faz a caixa vermelha pra ver colisao
	rl.DrawRectangleLines(
		int32(destinationRec.X-origin.X),
		int32(destinationRec.Y-origin.Y),
		int32(destinationRec.Width),
		int32(destinationRec.Height),
		rl.Red,
	)
}

func (p *Player) TakeDamage(damage int32) {
	if p.isInvincible(invencibilityDuration) {
		if p.Health > 1 {
			p.Health -= damage
			p.LastDamageTaken = time.Now()
		} else {
			system.GameOverFlag = true
		}
	}
}

func (p *Player) isInvincible(duration int) bool {
	if time.Since(p.LastDamageTaken).Milliseconds() > int64(duration) {
		return true
	}
	return false
}

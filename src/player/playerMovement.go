package player

import (
	"otaviocosta2110/getTheBlueBlocks/src/screen"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	animationDelay int32 = 300
)

var (
	framesWalkingX = []int{0, 1}
	framesWalkingY = []int{0, 0}
)

func (player *Player) CheckMovement(screen screen.Screen) {
	if rl.IsKeyDown(rl.KeyLeft) && player.X > 0 {
		player.X -= player.Speed
		player.Flipped = true
		player.UpdateAnimation(300, framesWalkingX, framesWalkingY)
	}

	if rl.IsKeyDown(rl.KeyRight) && player.X < screen.Width-(player.Width*player.Scale)/2 {
		player.X += player.Speed
		player.Flipped = false
		player.UpdateAnimation(300, framesWalkingX, framesWalkingY)
	}

	if rl.IsKeyDown(rl.KeyUp) && player.Y > player.Height*player.Scale-player.Y {
		player.Y -= player.Speed
		player.UpdateAnimation(300, framesWalkingX, framesWalkingY)
	}

	if rl.IsKeyDown(rl.KeyDown) && player.Y < screen.Height-(player.Height*player.Scale)/2 {
		player.Y += player.Speed
		player.UpdateAnimation(300, framesWalkingX, framesWalkingY)
	}

}

func (player *Player) CheckAtk() {
	if rl.IsKeyPressed(rl.KeyZ) {
		player.UpdateAnimation(300, []int{0, 1}, []int{1, 1})
	}
}

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
      println(framesX[nextIndex], framesY[nextIndex])
		}
    println(framesX[0], framesY[0])

		p.LastFrameTime = time.Now()
	}
}

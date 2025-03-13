package player

import (
	"otaviocosta2110/getTheBlueBlocks/src/physics"
	"otaviocosta2110/getTheBlueBlocks/src/screen"

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
		player.UpdateAnimation(int(animationDelay), framesWalkingX, framesWalkingY)
	} else if rl.IsKeyDown(rl.KeyRight) && player.X < screen.Width-(player.Width*player.Scale)/2 {
		player.X += player.Speed
		player.Flipped = false
		player.UpdateAnimation(int(animationDelay), framesWalkingX, framesWalkingY)
	}

	if rl.IsKeyDown(rl.KeyUp) && player.Y > player.Height*player.Scale-player.Y {
		player.Y -= player.Speed
		player.UpdateAnimation(int(animationDelay), framesWalkingX, framesWalkingY)
	} else if rl.IsKeyDown(rl.KeyDown) && player.Y < screen.Height-(player.Height*player.Scale)/2 {
		player.Y += player.Speed
		player.UpdateAnimation(int(animationDelay), framesWalkingX, framesWalkingY)
	}
}

func (player *Player) CheckAtk(enemyX, enemyY, enemyWidth, enemyHeight int32) bool {
	var isAttacking = false
	if rl.IsKeyPressed(rl.KeyZ) {
		isAttacking = true

		player.UpdateAnimation(50, []int{0, 1}, []int{1, 1})

		punchX := player.X
		punchY := player.Y

		punchWidth := player.Width
		punchHeight := player.Height / 2

		if player.Flipped {
			punchX -= punchWidth * 2 //esquerda
		} else {
			punchX += player.Width //direita, n sei pq ta assim
		}

		// cor da colisão do soco (debug)
		rl.DrawRectangle(punchX, punchY, punchWidth, punchHeight, rl.Red)

		return physics.CheckCollision(punchX, punchY, enemyX, enemyY, punchWidth, punchHeight)
	}
	if !isAttacking {
		player.UpdateAnimation(int(animationDelay), []int{0}, []int{0})
	}
	return false
}

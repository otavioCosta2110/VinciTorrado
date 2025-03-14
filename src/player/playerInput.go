package player

import (
	"otaviocosta2110/getTheBlueBlocks/src/physics"
	"otaviocosta2110/getTheBlueBlocks/src/screen"
	"otaviocosta2110/getTheBlueBlocks/src/system"

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
	if rl.IsKeyDown(rl.KeyLeft) && player.Object.X > player.Object.Width/2 {
		player.Object.X -= player.Speed
		player.Flipped = true
		player.UpdateAnimation(int(animationDelay), framesWalkingX, framesWalkingY)

	} else if rl.IsKeyDown(rl.KeyRight) && player.Object.X < screen.Width-(player.Object.Width)/2 {
		player.Object.X += player.Speed
		player.Flipped = false
		player.UpdateAnimation(int(animationDelay), framesWalkingX, framesWalkingY)
	}

	if rl.IsKeyDown(rl.KeyUp) && player.Object.Y > player.Object.Height-player.Object.Y {
		player.Object.Y -= player.Speed
		player.UpdateAnimation(int(animationDelay), framesWalkingX, framesWalkingY)

	} else if rl.IsKeyDown(rl.KeyDown) && player.Object.Y < screen.Height-(player.Object.Height)/2 {
		player.Object.Y += player.Speed
		player.UpdateAnimation(int(animationDelay), framesWalkingX, framesWalkingY)
	}
}

func (player *Player) CheckAtk(enemyObj system.Object) bool {
	var isAttacking = false
	if rl.IsKeyPressed(rl.KeyZ) {
		isAttacking = true

		player.UpdateAnimation(50, []int{0, 1}, []int{1, 1})

		punchX := player.Object.X
		punchY := player.Object.Y

		punchWidth := player.Object.Width
		punchHeight := player.Object.Height / 2

		if player.Flipped {
			punchX -= punchWidth * 2 //esquerda
		} else {
			punchX += player.Object.Width //direita, n sei pq ta assim
		}

    punchObj := system.Object{
      punchX,
      punchY,
      punchWidth,
      punchHeight,
    }

		// cor da colisão do soco (debug)
		rl.DrawRectangle(punchX, punchY, punchWidth, punchHeight, rl.Red)

		return physics.CheckCollision(punchObj, enemyObj)
	}
	if !isAttacking {
		player.UpdateAnimation(int(animationDelay), []int{0}, []int{0})
	}
	return false
}

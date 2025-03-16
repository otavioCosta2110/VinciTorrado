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
		player.Object.UpdateAnimation(int(animationDelay), framesWalkingX, framesWalkingY)

	} else if rl.IsKeyDown(rl.KeyRight) && player.Object.X < screen.Width-(player.Object.Width)/2 {
		player.Object.X += player.Speed
		player.Flipped = false
		player.Object.UpdateAnimation(int(animationDelay), framesWalkingX, framesWalkingY)
	}

	if rl.IsKeyDown(rl.KeyUp) && player.Object.Y > player.Object.Height-player.Object.Y {
		player.Object.Y -= player.Speed
		player.Object.UpdateAnimation(int(animationDelay), framesWalkingX, framesWalkingY)

	} else if rl.IsKeyDown(rl.KeyDown) && player.Object.Y < screen.Height-(player.Object.Height)/2 {
		player.Object.Y += player.Speed
		player.Object.UpdateAnimation(int(animationDelay), framesWalkingX, framesWalkingY)
	}
}

func (player *Player) CheckAtk(enemyObj system.Object) bool {
	var isAttacking = false

	// botei essas vars pra ca pra fazer a caixa de colisao aparecer sempre
	punchX := player.Object.X
	punchY := player.Object.Y - player.Object.Height/3

	punchWidth := player.Object.Width
	punchHeight := player.Object.Height / 2

	if player.Flipped {
		punchX -= punchWidth + punchWidth/2 //esquerda
	} else {
		punchX += punchWidth / 2 //direita, n sei pq ta assim
	}

	// cor da colisão do soco (debug)
	rl.DrawRectangle(punchX, punchY, punchWidth, punchHeight, rl.Red)

	if rl.IsKeyPressed(rl.KeyZ) {
		isAttacking = true

		player.Object.UpdateAnimation(50, []int{0, 1}, []int{1, 1})

		punchObj := system.Object{
			X:      punchX,
			Y:      punchY,
			Width:  punchWidth,
			Height: punchHeight,
		}

		return physics.CheckCollision(punchObj, enemyObj)
	}
	if !isAttacking {
		player.Object.UpdateAnimation(int(animationDelay), []int{0}, []int{0})
	}
	return false
}

func (player *Player) CheckKick(enemyObj system.Object) bool {
	var isKicking = false

	// Define a área do chute
	kickX := player.Object.X
	kickY := player.Object.Y - player.Object.Height/3

	kickWidth := player.Object.Width
	kickHeight := player.Object.Height / 2

	if player.Flipped {
		kickX -= kickWidth + kickWidth/2 // Chute para a esquerda
	} else {
		kickX += kickWidth / 2 // Chute para a direita
	}

	// Desenha a caixa de colisão do chute (para debug)
	rl.DrawRectangle(kickX, kickY, kickWidth, kickHeight, rl.Blue)

	if rl.IsKeyPressed(rl.KeyX) {
		isKicking = true

		player.Object.UpdateAnimation(50, []int{0, 1}, []int{0, 2})

		kickObj := system.Object{
			X:      kickX,
			Y:      kickY,
			Width:  kickWidth,
			Height: kickHeight,
		}

		return physics.CheckCollision(kickObj, enemyObj)
	}

	if !isKicking {
		player.Object.UpdateAnimation(int(animationDelay), []int{0}, []int{0}) // Volta para a animação padrão
	}
	return false
}

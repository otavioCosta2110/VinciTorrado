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
	if player.Object.FrameY == 1 || player.Object.FrameY == 2 {
		return
	}

	if rl.IsKeyDown(rl.KeyLeft) && player.Object.X > player.Object.Width/2 {
		player.Object.X -= player.Speed
		player.Flipped = true
		player.Object.UpdateAnimation(int(animationDelay), framesWalkingX, framesWalkingY)

	} else if rl.IsKeyDown(rl.KeyRight) && float32(player.Object.X) < screen.Camera.Target.X+float32(screen.Width)/2.0-float32(player.Object.Width/2.0) {
		player.Object.X += player.Speed
		player.Flipped = false
		player.Object.UpdateAnimation(int(animationDelay), framesWalkingX, framesWalkingY)
	}

	if rl.IsKeyDown(rl.KeyUp) && player.Object.Y > player.Object.Height-player.Object.Y+(screen.ScenaryHeight+player.Object.Height) {
		player.Object.Y -= player.Speed
		player.Object.UpdateAnimation(int(animationDelay), framesWalkingX, framesWalkingY)

	} else if rl.IsKeyDown(rl.KeyDown) && player.Object.Y < screen.Height-(player.Object.Height)/2 {
		player.Object.Y += player.Speed
		player.Object.UpdateAnimation(int(animationDelay), framesWalkingX, framesWalkingY)
	}
}

func (player *Player) CheckAtk(enemyObj system.Object) bool {
	var isAttacking = false
	punchWidth := float32(player.Object.Width) 
	punchHeight := player.Object.Height / 2

	punchX := player.Object.X - player.Object.Width*2
	punchY := player.Object.Y - player.Object.Height/4

	if player.Flipped {
		punchX = (player.Object.X - player.Object.Width/2) 
	} else {
		punchX = ( player.Object.X + player.Object.Width/2 )
	}

	if rl.IsKeyPressed(rl.KeyZ) {
		isAttacking = true

		player.Object.UpdateAnimation(50, []int{0, 1}, []int{1, 1})

		punchObj := system.Object{
			X:      punchX,
			Y:      punchY,
			Width:  int32(punchWidth),
			Height: punchHeight,
		}

		return physics.CheckCollision(punchObj, enemyObj)
	}
	if !isAttacking {
		player.Object.UpdateAnimation(int(animationDelay), []int{0}, []int{0})
	}
	return false
}

func (player *Player) CheckKick(enemyObj *system.Object) bool {
	kickX := player.Object.X
	kickY := player.Object.Y

	kickWidth := player.Object.Width
	kickHeight := player.Object.Height / 2

	if player.Flipped {
		kickX -= kickWidth + kickWidth/2
	} else {
		kickX += kickWidth / 2
	}

	if rl.IsKeyPressed(rl.KeyX) {
		player.IsKicking = true
		player.Object.UpdateAnimation(50, []int{0, 0}, []int{2, 0})

		kickObj := system.Object{
			X:      kickX,
			Y:      kickY,
			Width:  kickWidth,
			Height: kickHeight,
		}

		// nao me pergunta porque precisa disso
		newEnObj := system.Object{
			X:      enemyObj.X - enemyObj.Width/2,
			Y:      enemyObj.Y - enemyObj.Height/2,
			Width:  enemyObj.Width,
			Height: enemyObj.Height,
		}

		if physics.CheckCollision(kickObj, *&newEnObj) {

			enemyObj.SetKnockback(kickObj)
			player.IsKicking = false
			return true
		}
	} else {
		player.IsKicking = false
	}

	return false
}

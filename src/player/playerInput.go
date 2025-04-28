package player

import (
	"otaviocosta2110/vincitorrado/src/audio"
	"otaviocosta2110/vincitorrado/src/objects"
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/system"
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
		punchX = (player.Object.X + player.Object.Width/2)
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

		if physics.CheckCollision(punchObj, enemyObj) {
			if !enemyObj.Destroyed{
				audio.PlayPunch()
			}

			return true
		}
	}
	if !isAttacking {
		player.Object.UpdateAnimation(int(animationDelay), []int{0}, []int{0})
	}
	return false
}

func (player *Player) CheckKick(box *objects.Box) bool {
	if rl.IsKeyPressed(rl.KeyX) && time.Since(player.LastKickTime) > player.KickCooldown {
		player.IsKicking = true
		player.LastKickTime = time.Now()
		player.Object.UpdateAnimation(50, []int{0}, []int{2})

		box.OriginalY = box.Object.Y

		kickWidth := player.Object.Width * 2
		kickHeight := int32(float32(player.Object.Height) * 1.5)

		kickX := player.Object.X
		kickY := player.Object.Y - player.Object.Height/4

		if player.Flipped {
			kickX -= kickWidth
		} else {
			kickX += player.Object.Width
		}

		rl.DrawRectangle(kickX, kickY, kickWidth, kickHeight, rl.NewColor(0, 0, 255, 128))

		kickObj := system.Object{
			X:      kickX,
			Y:      kickY,
			Width:  kickWidth,
			Height: kickHeight,
		}

		if physics.CheckCollision(kickObj, box.Object) {
			knockbackMultiplier := int32(3)

			if player.Flipped {
				box.Object.KnockbackX = -player.KickPower * knockbackMultiplier
			} else {
				box.Object.KnockbackX = player.KickPower * knockbackMultiplier
			}

			box.Object.KnockbackY = 0
			audio.PlayKick()
			return true
		}
	}
	player.IsKicking = false
	return false
}

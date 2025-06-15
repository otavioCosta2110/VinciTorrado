package player

import (
	"otaviocosta2110/vincitorrado/src/audio"
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/system"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	animationDelay  int32 = 300
	attackAnimSpeed       = 50
	kickAnimSpeed         = 50
)

var (
	framesWalkingX = []int32{0, 1}
	framesWalkingY = []int32{0, 0}
	framesAttackX  = []int32{0, 1}
	framesAttackY  = []int32{1, 1}
	frameKickX     = []int32{0}
	frameKickY     = []int32{3}
)

func (player *Player) CheckMovement(screen screen.Screen) {
	if player.Object.FrameY != 0 {
		return
	}

	if rl.IsKeyDown(rl.KeyLeft) && float32(player.Object.X) > screen.Camera.Target.X-float32(player.Screen.Width)/2+float32(player.Object.Width/2) {
		player.Object.X -= player.Speed
		player.Object.Flipped = true
		player.UpdateAnimation("walk")

	} else if rl.IsKeyDown(rl.KeyRight) && float32(player.Object.X) < screen.Camera.Target.X+float32(screen.Width)/2.0-float32(player.Object.Width/2.0) {
		player.Object.X += player.Speed
		player.Object.Flipped = false
		player.UpdateAnimation("walk")
	}

	if rl.IsKeyDown(rl.KeyUp) && player.Object.Y > player.Object.Height-player.Object.Y+(screen.ScenaryHeight+player.Object.Height) {
		player.Object.Y -= player.Speed
		player.UpdateAnimation("walk")

	} else if rl.IsKeyDown(rl.KeyDown) && player.Object.Y < screen.Height-(player.Object.Height)/2 {
		player.Object.Y += player.Speed
		player.UpdateAnimation("walk")
	}
}

func (player *Player) CheckAtk(enemyObj system.Object) bool {
	var isAttacking = false
	punchWidth := int32(float32(player.Object.Width))
	punchHeight := float32(player.Object.Height / 2)

	if player.Weapon != nil {
		punchWidth = int32(float32(player.Object.Width) + float32(player.Weapon.HitboxX))
		punchHeight = float32(player.Object.Height/2) + float32(player.Weapon.HitboxY)
	}

	punchX := player.Object.X - player.Object.Width*2
	punchY := player.Object.Y - player.Object.Height/4

	if player.Object.Flipped {
		punchX = (player.Object.X - player.Object.Width/2)
	} else {
		punchX = (player.Object.X + player.Object.Width/2)
	}

	if rl.IsKeyPressed(rl.KeyZ) {
		isAttacking = true

		player.UpdateAnimation("punch")

		punchObj := system.Object{
			X:      punchX,
			Y:      punchY,
			Width:  int32(punchWidth),
			Height: int32(punchHeight),
		}

		if physics.CheckCollision(punchObj, enemyObj) {
			if !enemyObj.Destroyed {
				audio.PlayPunch()
				if player.Weapon != nil {
					player.Weapon.Health -= 1
					if player.Weapon.Health <= 0 {
						audio.PlayWeaponBreaking()
						player.DropWeapon()
					}
				}
			}
			return true
		}
	}
	if !isAttacking {
		player.UpdateAnimation("default")
	}
	return false
}

func (p *Player) CheckKick(kickables []physics.Kickable, items *[]*equipment.Equipment) bool {
	kickedSomething := false
	if rl.IsKeyPressed(rl.KeyX) && time.Since(p.LastKickTime) > p.KickCooldown {
		p.IsKicking = true
		p.LastKickTime = time.Now()
		p.Object.FrameY = 3
		p.Object.FrameX = 0
		p.Object.UpdateAnimation(50, []int32{0}, []int32{3})

		kickWidth := p.Object.Width
		kickHeight := p.Object.Height
		kickX := p.Object.X - p.Object.Width
		kickY := p.Object.Y - p.Object.Height/4

		if p.Object.Flipped {
			kickX -= kickWidth
		} else {
			kickX += p.Object.Width
		}

		kickHitbox := system.Object{
			X:      kickX,
			Y:      kickY,
			Width:  kickWidth,
			Height: kickHeight,
		}

		for _, obj := range kickables {
			if !obj.IsKicked() && physics.CheckCollision(kickHitbox, obj.GetObject()) {
				obj.HandleKick(items, p.Object)
				audio.PlayKick()
				kickedSomething = true
			}
		}
	}
	return kickedSomething
}

func (player *Player) HandleAttackInput() {
	if rl.IsKeyPressed(rl.KeyZ) {
		player.IsAttacking = true
		player.LastAttackTime = time.Now()
		player.UpdateAnimation("attack")
	}
}

package player

import (
	"math/rand"
	"otaviocosta2110/vincitorrado/src/audio"
	"otaviocosta2110/vincitorrado/src/equipment"
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

	if rl.IsKeyDown(rl.KeyLeft) && float32(player.Object.X) > screen.Camera.Target.X-float32(player.Screen.Width)/2+float32(player.Object.Width/2) {
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
			if !enemyObj.Destroyed {
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

func (p *Player) CheckKick(boxes []*objects.Box, trashCans []*objects.TrashCan, items *[]*equipment.Equipment) bool {
	kickedSomething := false

	if rl.IsKeyPressed(rl.KeyX) && time.Since(p.LastKickTime) > p.KickCooldown {
		p.IsKicking = true
		p.LastKickTime = time.Now()
		p.Object.UpdateAnimation(50, []int{0}, []int{2})

		kickWidth := p.Object.Width * 2
		kickHeight := p.Object.Height
		kickX := p.Object.X
		kickY := p.Object.Y - p.Object.Height/4

		if p.Flipped {
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

		for _, box := range boxes {
			if physics.CheckCollision(kickHitbox, box.Object) {
				knockbackMultiplier := int32(3)
				if p.Flipped {
					box.Object.KnockbackX = -p.KickPower * knockbackMultiplier
				} else {
					box.Object.KnockbackX = p.KickPower * knockbackMultiplier
				}
				box.Object.KnockbackY = 0
				audio.PlayKick()
				kickedSomething = true
			}
		}

		for _, trash := range trashCans {
			if !trash.Kicked && physics.CheckCollision(kickHitbox, trash.Object) {
				trash.Kicked = true
				item := *trash.LootTable[rand.Intn(len(trash.LootTable))]
				item.Object.X = trash.Object.X
				item.Object.Y = trash.Object.Y
				item.IsDropped = true
				*items = append(*items, &item)
				audio.PlayKick()
				kickedSomething = true
			}
		}
	}

	p.IsKicking = false
	return kickedSomething
}

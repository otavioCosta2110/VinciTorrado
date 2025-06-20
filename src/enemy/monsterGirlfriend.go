// TODO: grunidos
package enemy

import (
	"math"
	"otaviocosta2110/vincitorrado/src/audio"
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/system"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Enemy) startCharge(p system.Player, s screen.Screen) {
	if e.EnemyType != "gf_monster" {
		return
	}

	audio.PlayGfRunningSound()

	e.IsCharging = true

	dirX := float32(p.GetObject().X - e.Object.X)

	dirY := float32((p.GetObject().Y + p.GetObject().Height/2) - e.Object.Y)
	if p.GetObject().Y > s.Height/2 {
		dirY = float32((p.GetObject().Y - p.GetObject().Height/2) - e.Object.Y)
	}

	length := float32(math.Sqrt(float64(dirX*dirX + dirY*dirY)))

	if dirX > 0 {
		e.Object.Flipped = false
	} else {
		e.Object.Flipped = true
	}

	if length > 0 {
		e.ChargeDirection = rl.Vector2{
			X: dirX / length,
			Y: dirY / length,
		}
	} else {
		e.ChargeDirection = rl.Vector2{X: 1, Y: 0}
	}
}

func (e *Enemy) handleCharge(p system.Player) {
	e.Object.X += int32(e.ChargeDirection.X * float32(e.Speed))
	e.Object.Y += int32(e.ChargeDirection.Y * float32(e.Speed))

	screenWidth := int32(rl.GetScreenWidth())
	screenHeight := int32(rl.GetScreenHeight())

	enemyLeft := e.Object.X - e.Object.Width/2
	enemyRight := e.Object.X + e.Object.Width/2
	enemyTop := e.Object.Y - e.Object.Height
	enemyBottom := e.Object.Y

	hitWall := false
	if enemyLeft <= 0 || enemyRight >= screenWidth-e.Object.Width/e.Object.Scale {
		hitWall = true
	}
	if enemyTop <= 0 || enemyBottom >= screenHeight-e.Object.Height/2 {
		hitWall = true
	}

	if hitWall {
		e.onChargeCollision()
		e.IsStunned = true
		return
	}

	gfMonsterHitbox := system.Object{
		X:         e.Object.X,
		Y:         e.Object.Y + e.Object.Height/4,
		Width:     e.Object.Width,
		Height:    e.Object.Height / 2,
		Flipped:   e.Object.Flipped,
		Scale:     e.Object.Scale,
		Destroyed: e.Object.Destroyed,
	}
	if physics.CheckCollision(gfMonsterHitbox, p.GetObject()) {
		p.TakeDamage(e.Damage, e.Object)
		return
	}

}

func (e *Enemy) UpdateGirlfriendHealth() {
	if e.EnemyType != "gf_monster" {
		return
	}

	if time.Since(e.LastHealthDecrease) > 5*time.Second {
		e.Health -= 4
		e.LastHealthDecrease = time.Now()

		if e.Health <= 0 {
			e.Object.Destroyed = true
		}
	}
}

func (e *Enemy) onChargeCollision() {
	e.IsCharging = false
	e.IsStunned = true
	e.StunEndTime = time.Now().Add(2 * time.Second)

	e.UpdateAnimation("gf_stunned")

	audio.StopGfRunningSound()
	audio.PlayGfHittingWall()
}

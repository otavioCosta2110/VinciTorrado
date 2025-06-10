// TODO: som de correr, grunidos, som de bater na parede
package enemy

import (
	"math"
	"otaviocosta2110/vincitorrado/src/audio"
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/system"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Enemy) startCharge(p system.Player) {
	if e.EnemyType != "gf_monster" {
		return
	}

	audio.PlayGfRunningSound()

	e.IsCharging = true

	e.Speed = 9

	dirX := float32(p.GetObject().X - e.Object.X)
	dirY := float32(p.GetObject().Y - e.Object.Y)
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
	if enemyLeft <= 0 || enemyRight >= screenWidth-e.Object.Width/2 {
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

	if physics.CheckCollision(e.Object, p.GetObject()) {
		p.TakeDamage(e.Damage, e.Object)
		e.onChargeCollision()
		return
	}

}

func (e *Enemy) UpdateGirlfriendHealth() {
	if e.EnemyType != "gf_monster" {
		return
	}

	if time.Since(e.LastHealthDecrease) > 5*time.Second {
		e.Health -= 1
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

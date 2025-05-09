package enemy

import (
	"math"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/system"
	"time"
)

func MoveEnemyTowardPlayer(p system.Player, e Enemy, s screen.Screen) Enemy {
	hitStunDuration := time.Millisecond * 300
	if time.Since(e.LastDamageTaken) < hitStunDuration {
		e.CanMove = false
	}

	if !e.CanMove {
		return e
	}

	if e.Object.FrameY == 1 {
		return e
	}

	playerX := float64(p.GetObject().X)
	playerY := float64(p.GetObject().Y)
	enemyX := float64(e.Object.X)
	enemyY := float64(e.Object.Y)

	distX := playerX - enemyX
	distY := playerY - enemyY
	distance := math.Sqrt(distX*distX + distY*distY)

	punchRange := float64(e.Object.Width) * 1.2

	if distance > punchRange {
		if distance > 0 {
			distX /= distance
			distY /= distance
		}

		e.Object.X += int32((distX * float64(e.Speed)) * 1.5)
		e.Object.Y += int32((distY * float64(e.Speed)) * 1.5)

		if distX > 0 {
			e.Object.Flipped = false
		} else if distX < 0 {
			e.Object.Flipped = true
		}
	}

	return e
}

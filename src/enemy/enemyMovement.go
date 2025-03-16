package enemy

import (
	"math"
	"otaviocosta2110/getTheBlueBlocks/src/player"
	"otaviocosta2110/getTheBlueBlocks/src/screen"
)

func MoveEnemyTowardPlayer(p player.Player, e Enemy, s screen.Screen) Enemy {
	if e.Object.FrameY == 1 {
		return e
	}

	distX := float64(p.Object.X - e.Object.X)
	distY := float64(p.Object.Y - e.Object.Y)
	distance := math.Sqrt(distX*distX + distY*distY) 

	stopDistance := float64(e.Object.Width) * 0.8 

	if distance > stopDistance { 
		if e.Object.X < p.Object.X {
			e.Flipped = false
			e.Object.X += e.Speed
		}
		if e.Object.X > p.Object.X {
			e.Flipped = true
			e.Object.X -= e.Speed
		}
		if e.Object.Y < p.Object.Y {
			e.Object.Y += e.Speed
		}
		if e.Object.Y > p.Object.Y {
			e.Object.Y -= e.Speed
		}
	}
	return e
}

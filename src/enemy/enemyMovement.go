package enemy

import (
	"math"
	"otaviocosta2110/getTheBlueBlocks/src/screen"
	"otaviocosta2110/getTheBlueBlocks/src/system"
)

func MoveEnemyTowardPlayer(p system.Player, e Enemy, s screen.Screen) Enemy {
	if e.Object.FrameY == 1 {
		return e
	}

	distX := float64(p.GetObject().X - e.Object.X)
	distY := float64(p.GetObject().Y - e.Object.Y)
	distance := math.Sqrt(distX*distX + distY*distY) 

	stopDistance := float64(e.Object.Width) * 0.8 

	if distance > stopDistance { 
		if e.Object.X < p.GetObject().X {
			e.Flipped = false
			e.Object.X += e.Speed
		}
		if e.Object.X > p.GetObject().X {
			e.Flipped = true
			e.Object.X -= e.Speed
		}
		if e.Object.Y < p.GetObject().Y {
			e.Object.Y += e.Speed
		}
		if e.Object.Y > p.GetObject().Y {
			e.Object.Y -= e.Speed
		}
	}
	return e
}

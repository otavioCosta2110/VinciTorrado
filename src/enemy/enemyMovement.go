package enemy

import (
	"otaviocosta2110/getTheBlueBlocks/src/player"
	"otaviocosta2110/getTheBlueBlocks/src/screen"
)

func MoveEnemyTowardPlayer(p player.Player, e Enemy, s screen.Screen) Enemy {
	if e.X < p.X {
		e.Flipped = false
		e.X += e.Speed
	}
	if e.X > p.X {
		e.Flipped = true
		e.X -= e.Speed
	}
	if e.Y < p.Y {
		e.Y += e.Speed
	}
	if e.Y > p.Y {
		e.Y -= e.Speed
	}

	return e
}

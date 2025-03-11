package enemy

import (
	"otaviocosta2110/getTheBlueBlocks/src/physics"
	"otaviocosta2110/getTheBlueBlocks/src/player"
)

func MoveEnemyTowardPlayer(p player.Player, e Enemy) (Enemy){
	if e.X < p.X && !physics.CheckCollision(e.X+e.Speed, e.Y, p.X, p.Y, p.Width, p.Height) {
		e.X += e.Speed
	}
	if e.X > p.X && !physics.CheckCollision(e.X-e.Speed, e.Y, p.X, p.Y, p.Width, p.Height) {
		e.X -= e.Speed
	}
	if e.Y < p.Y && !physics.CheckCollision(e.X, e.Y+e.Speed, p.X, p.Y, p.Width, p.Height) {
		e.Y += e.Speed
	}
	if e.Y > p.Y && !physics.CheckCollision(e.X, e.Y-e.Speed, p.X, p.Y, p.Width, p.Height) {
		e.Y -= e.Speed
	}

  return e
}

package player

import (
	"otaviocosta2110/getTheBlueBlocks/src/screen"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (player Player) CheckMovement(screen screen.Screen) Player {
	if rl.IsKeyDown(rl.KeyLeft) && player.X > 0 {
		player.X -= player.Speed
	}
	if rl.IsKeyDown(rl.KeyRight) && player.X+player.Width < screen.Width {
		player.X += player.Speed
	}
	if rl.IsKeyDown(rl.KeyUp) && player.Y > 0 {
		player.Y -= player.Speed
	}
	if rl.IsKeyDown(rl.KeyDown) && player.Y+player.Height < screen.Height {
		player.Y += player.Speed
	}

	return player
}


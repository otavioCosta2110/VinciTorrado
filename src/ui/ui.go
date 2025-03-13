package ui

import (
	"otaviocosta2110/getTheBlueBlocks/src/player"
	"otaviocosta2110/getTheBlueBlocks/src/screen"
  rl "github.com/gen2brain/raylib-go/raylib"

)
func DrawLife(s screen.Screen, p *player.Player) {
	barWidth := 200
	barHeight := 20
	healthPercentage := float32(p.Health) / float32(p.MaxHealth) 

	rl.DrawRectangle(20, 20, int32(barWidth), int32(barHeight), rl.Gray)

	rl.DrawRectangle(20, 20, int32(float32(barWidth)*healthPercentage), int32(barHeight), rl.Red)

	rl.DrawRectangleLines(20, 20, int32(barWidth), int32(barHeight), rl.Black)
}

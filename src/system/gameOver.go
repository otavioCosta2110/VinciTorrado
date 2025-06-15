package system

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"otaviocosta2110/vincitorrado/src/screen"
)

var GameOverFlag bool = false

func DrawGameOver(screen *screen.Screen) {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.NewColor(0, 0, 0, 200))

	text := "GAME OVER"
	fontSize := int32(72)
	textWidth := rl.MeasureText(text, fontSize)
	rl.DrawText(text,
		screen.Width/2-textWidth/2,
		screen.Height/2-fontSize,
		fontSize, rl.Red)

	restartText := "Press R to restart"
	restartFontSize := int32(36)
	restartWidth := rl.MeasureText(restartText, restartFontSize)
	rl.DrawText(restartText,
		screen.Width/2-restartWidth/2,
		screen.Height/2+restartFontSize,
		restartFontSize, rl.White)

	exitText := "Press ESC to return to menu"
	exitFontSize := int32(24)
	exitWidth := rl.MeasureText(exitText, exitFontSize)
	rl.DrawText(exitText,
		screen.Width/2-exitWidth/2,
		screen.Height/2+restartFontSize*2,
		exitFontSize, rl.Gray)
}

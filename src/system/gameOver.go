package system

import (
	rl "github.com/gen2brain/raylib-go/raylib"
  "otaviocosta2110/vincitorrado/src/screen"
)

var GameOverFlag bool = false


func GameOver(s *screen.Screen) {
	text := "Game Over"
	textWidth := rl.MeasureText(text, 100)

	xPos := (s.Width - textWidth) / 2
	yPos := s.Height / 2

	rl.DrawText(text, xPos, yPos, 100, rl.White)
}

package ui

import (
	"otaviocosta2110/vincitorrado/src/player"
	"otaviocosta2110/vincitorrado/src/screen"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	heartTexture rl.Texture2D
	textureLoaded bool = false
)

func DrawLife(s screen.Screen, p *player.Player) {
	// Load texture only once
	if !textureLoaded {
		heartTexture = rl.LoadTexture("assets/ui/heart.png")
		textureLoaded = true
	}

	scale := float32(p.Object.Scale)
	heartWidth := int32(heartTexture.Width * p.Object.Scale)   // Adjust based on your heart image size
	heartHeight := int32(heartTexture.Height * p.Object.Scale) // Adjust based on your heart image size
	padding := int32(5 * scale) // Scale padding too

	posX := int32(20)
	posY := int32(20)

	for i := int32(0); i < p.MaxHealth; i++ {
		destRec := rl.NewRectangle(
			float32(posX),
			float32(posY),
			float32(heartWidth),
			float32(heartHeight),
		)

		sourceRec := rl.NewRectangle(
			0,
			0,
			float32(heartTexture.Width),
			float32(heartTexture.Height),
		)

		color := rl.White
		if i >= p.Health {
			color = rl.Fade(rl.White, 0.3)
		}

		rl.DrawTexturePro(
			heartTexture,
			sourceRec,
			destRec,
			rl.NewVector2(0, 0),
			0,
			color,
		)

		posX += heartWidth + padding // Moved INSIDE the loop
	}
}

package ui

import (
	"otaviocosta2110/vincitorrado/src/enemy"
	"otaviocosta2110/vincitorrado/src/player"
	"otaviocosta2110/vincitorrado/src/screen"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	heartTexture  rl.Texture2D
	textureLoaded bool = false
)

func DrawLife(s screen.Screen, p *player.Player) {
	if p.Object.Destroyed {
		return
	}
	if !textureLoaded {
		heartTexture = rl.LoadTexture("assets/ui/heart.png")
		textureLoaded = true
	}

	scale := float32(p.Object.Scale)
	heartWidth := int32(heartTexture.Width * p.Object.Scale)
	heartHeight := int32(heartTexture.Height * p.Object.Scale)
	padding := int32(5 * scale)

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

		posX += heartWidth + padding
	}
}

func DrawBossHealthBar(boss *enemy.Enemy, screenWidth int32) {
	if boss == nil || boss.Object.Destroyed {
		return
	}

	barWidth := int32(300)
	barHeight := int32(30)
	margin := int32(20)

	posX := screenWidth - barWidth - margin
	posY := int32(30)

	rl.DrawRectangle(posX, posY, barWidth, barHeight, rl.DarkGray)

	healthPercent := float32(boss.Health) / float32(boss.MaxHealth)
	fillWidth := int32(float32(barWidth) * healthPercent)

	rl.DrawRectangle(posX, posY, fillWidth, barHeight, rl.Red)
	rl.DrawRectangleLines(posX, posY, barWidth, barHeight, rl.White)
	/*
		name := "Giiiiiiirlfriend"
		fontSize := int32(20)
		textWidth := rl.MeasureText(name, fontSize)
		rl.DrawText(
			name,
			posX+barWidth/2-textWidth/2,
			posY-fontSize-5,
			fontSize,
			rl.White,
		)*/
}

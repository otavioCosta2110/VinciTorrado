package ui

import (
	"otaviocosta2110/vincitorrado/src/screen"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StartMenu struct {
	BgTexture rl.Texture2D
}

func NewStartMenu() *StartMenu {
	return &StartMenu{
		BgTexture: rl.LoadTexture("assets/ui/landscape.png"),
	}
}

func NewEndingMenu() *StartMenu {
	return &StartMenu{
		BgTexture: rl.LoadTexture("assets/ui/ending.png"),
	}
}

func (sm *StartMenu) DrawStartMenu(screen *screen.Screen) bool {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Black)

	screenWidth := float32(screen.Width)
	screenHeight := float32(screen.Height)

	bgWidth := float32(sm.BgTexture.Width)
	bgHeight := float32(sm.BgTexture.Height)

	cropRect := rl.NewRectangle(0, 0, bgWidth, bgHeight)

	if bgWidth > screenWidth {
		cropRect.X = (bgWidth - screenWidth) / 2
		cropRect.Width = screenWidth
	}

	if bgHeight > screenHeight {
		cropRect.Y = (bgHeight - screenHeight) / 2
		cropRect.Height = screenHeight
	}

	rl.DrawTexturePro(
		sm.BgTexture,
		cropRect,
		rl.NewRectangle(0, 0, screenWidth, screenHeight),
		rl.NewVector2(0, 0),
		0,
		rl.White,
	)

	rl.DrawRectangle(0, 0, int32(screenWidth), int32(screenHeight), rl.NewColor(0, 0, 0, 180))

	title := "VINCI TORRADO"
	subtitle := "Press ENTER to start"

	titleFontSize := int32(72)
	subtitleFontSize := int32(36)

	titleWidth := rl.MeasureText(title, titleFontSize)
	subtitleWidth := rl.MeasureText(subtitle, subtitleFontSize)

	rl.DrawText(title,
		screen.Width/2-titleWidth/2+2,
		screen.Height/2-titleFontSize+2,
		titleFontSize, rl.Black)
	rl.DrawText(title,
		screen.Width/2-titleWidth/2,
		screen.Height/2-titleFontSize,
		titleFontSize, rl.White)

	rl.DrawText(subtitle,
		screen.Width/2-subtitleWidth/2+2,
		screen.Height/2+subtitleFontSize+2,
		subtitleFontSize, rl.Black)
	rl.DrawText(subtitle,
		screen.Width/2-subtitleWidth/2,
		screen.Height/2+subtitleFontSize,
		subtitleFontSize, rl.White)

	controls := "Controls: Arrows to move, Z to punch, X to kick, ESC for menu"
	controlsFontSize := int32(20)
	controlsWidth := rl.MeasureText(controls, controlsFontSize)
	rl.DrawText(controls,
		screen.Width/2-controlsWidth/2,
		screen.Height-50,
		controlsFontSize, rl.LightGray)

	return rl.IsKeyPressed(rl.KeyEnter)
}

func (sm *StartMenu) DrawEndingMenu(screen *screen.Screen) {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Black)

	screenWidth := float32(screen.Width)
	screenHeight := float32(screen.Height)

	bgWidth := float32(sm.BgTexture.Width)
	bgHeight := float32(sm.BgTexture.Height)

	cropRect := rl.NewRectangle(0, 0, bgWidth, bgHeight)

	if bgWidth > screenWidth {
		cropRect.X = (bgWidth - screenWidth) / 2
		cropRect.Width = screenWidth
	}

	if bgHeight > screenHeight {
		cropRect.Y = (bgHeight - screenHeight) / 2
		cropRect.Height = screenHeight
	}

	rl.DrawTexturePro(
		sm.BgTexture,
		cropRect,
		rl.NewRectangle(0, 0, screenWidth, screenHeight),
		rl.NewVector2(0, 0),
		0,
		rl.White,
	)

	rl.DrawRectangle(0, 0, int32(screenWidth), int32(screenHeight), rl.NewColor(0, 0, 0, 180))

	title := "The End"

	titleFontSize := int32(72)

	titleWidth := rl.MeasureText(title, titleFontSize)

	rl.DrawText(title,
		screen.Width/2-titleWidth/2+2,
		screen.Height/2-titleFontSize+2,
		titleFontSize, rl.Black)
	rl.DrawText(title,
		screen.Width/2-titleWidth/2,
		screen.Height/2-titleFontSize,
		titleFontSize, rl.White)
}

package main

import (
	"fmt"
	"otaviocosta2110/getTheBlueBlocks/src/enemy"

	"otaviocosta2110/getTheBlueBlocks/src/player"
	"otaviocosta2110/getTheBlueBlocks/src/screen"
	"otaviocosta2110/getTheBlueBlocks/src/sprites"
	"otaviocosta2110/getTheBlueBlocks/src/system"
	"otaviocosta2110/getTheBlueBlocks/src/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowWidth   int32 = 1280
	windowHeight  int32 = 720
	obstacleSpeed int32 = 2
	playerScale   int32 = 4
	playerSizeX   int32 = 32
	playerSizeY   int32 = 32
)

func main() {
	screen := screen.NewScreen(windowWidth, windowHeight, "jogo poggers")
	rl.InitWindow(screen.Width, screen.Height, screen.Title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	playerSprite := sprites.Sprite{
		SpriteWidth:  playerSizeX,
		SpriteHeight: playerSizeY,
		Texture:      rl.LoadTexture("assets/player.png"),
	}

	enemySprite := sprites.Sprite{
		SpriteWidth:  playerSizeX,
		SpriteHeight: playerSizeY,
		Texture:      rl.LoadTexture("assets/enemy.png"),
	}

	player := player.NewPlayer(screen.Width/2, screen.Height/2, playerSizeX, playerSizeY, 4, playerScale, playerSprite)
	enemyManager := enemy.EnemyManager{}

	// enemyManager.AddEnemy(enemy.NewEnemy(50, 80, obstacleSpeed, playerSizeX, playerSizeY, playerScale, enemySprite))
	// enemyManager.AddEnemy(enemy.NewEnemy(200, 150, obstacleSpeed, playerSizeX, playerSizeY, playerScale, enemySprite))
	enemyManager.AddEnemy(enemy.NewEnemy(1000000, 10000, obstacleSpeed, playerSizeX, playerSizeY, playerScale, enemySprite))

	for !rl.WindowShouldClose() {
		update(player, &enemyManager, screen)
		draw(player, &enemyManager, *screen)
	}
}

func update(p *player.Player, em *enemy.EnemyManager, screen *screen.Screen) {
	if system.GameOverFlag {
		return
	}

	em.Update(p, *screen)
	p.Update(em, *screen)

	screen.UpdateCamera(p.Object.X, p.Object.Y)
}

func draw(p *player.Player, em *enemy.EnemyManager, s screen.Screen) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode2D(s.Camera)

	chao := rl.LoadTexture("assets/chao.png")
	chao.Width *= playerScale
	chao.Height *= playerScale
	buildings := rl.LoadTexture("assets/predio.png")
	buildings.Width *= playerScale
	buildings.Height *= playerScale
	drawTiledBackground(chao, s.Camera, s.Width, s.Height)
	drawBuildings(buildings)

	p.Draw()
	em.Draw()

	rl.EndMode2D()

	if system.GameOverFlag {
		system.GameOver(&s)
	}

	ui.DrawLife(s, p)
	rl.DrawText(fmt.Sprintf("Player: %d, %d", p.Object.X, p.Object.Y), 10, 10, 10, rl.Black)
	rl.DrawText(fmt.Sprintf("Enemies: %d", len(em.Enemies)), 10, 25, 10, rl.Black)

	rl.EndDrawing()
}

func drawTiledBackground(texture rl.Texture2D, camera rl.Camera2D, screenWidth, screenHeight int32) {
	texWidth := texture.Width
	texHeight := texture.Height

	visibleStartX := int32(camera.Target.X) - screenWidth/2 - texWidth
	visibleStartY := int32(camera.Target.Y) - screenHeight/2 - texHeight
	visibleEndX := int32(camera.Target.X) + screenWidth/2 + texWidth
	visibleEndY := int32(camera.Target.Y) + screenHeight/2 + texHeight

	startTileX := visibleStartX / texWidth
	startTileY := visibleStartY / texHeight
	endTileX := visibleEndX/texWidth + 1
	endTileY := visibleEndY/texHeight + 1

	for y := startTileY; y <= endTileY; y++ {
		for x := startTileX; x <= endTileX; x++ {
			rl.DrawTexture(
				texture,
				x*texWidth,
				y*texHeight,
				rl.White,
			)
		}
	}
}

func drawBuildings(texture rl.Texture2D) {
	rl.DrawTexture(
		texture,
		0, 
		0, 
		rl.White,
	)
}

package main

import (
	"fmt"
	"otaviocosta2110/getTheBlueBlocks/src/enemy"
	// "otaviocosta2110/getTheBlueBlocks/src/objects"
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
	
	enemyManager.AddEnemy(enemy.NewEnemy(50, 80, obstacleSpeed, playerSizeX, playerSizeY, playerScale, enemySprite))
	enemyManager.AddEnemy(enemy.NewEnemy(200, 150, obstacleSpeed, playerSizeX, playerSizeY, playerScale, enemySprite))
	enemyManager.AddEnemy(enemy.NewEnemy(300, 300, obstacleSpeed, playerSizeX, playerSizeY, playerScale, enemySprite))

	// box := objects.NewBox(400, 400, 50, 50, rl.Brown)

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
}

func draw(p *player.Player, em *enemy.EnemyManager, s screen.Screen) {
	rl.BeginDrawing()
	
	rl.ClearBackground(rl.RayWhite)
	
	chao := rl.LoadTexture("assets/chao.png")
	chao.Width *= playerScale
	chao.Height *= playerScale

	tilesX := s.Width/(chao.Width) + 1 
	tilesY := s.Height/(chao.Height) + 1

	for y := 0; int32(y) < tilesY; y++ {
		for x := 0; int32(x) < tilesX; x++ {
			rl.DrawTexture(
				chao,
				(int32(x)*chao.Width),
				(int32(y)*chao.Height),
				rl.White,
			)
		}
	}

	if system.GameOverFlag {
		system.GameOver(&s)
		rl.EndDrawing()
		return
	}

	p.Draw()
	em.Draw()
	// box.Draw()
	ui.DrawLife(s, p)

	rl.DrawText(fmt.Sprintf("Player: %d, %d", p.Object.X, p.Object.Y), 10, 10, 10, rl.Black)
	rl.DrawText(fmt.Sprintf("Enemies: %d", len(em.Enemies)), 10, 25, 10, rl.Black)

	rl.EndDrawing()
}

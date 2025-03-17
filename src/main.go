package main

import (
	"fmt"
	"math/rand"
	"otaviocosta2110/getTheBlueBlocks/src/enemy"
	"otaviocosta2110/getTheBlueBlocks/src/objects"
	"otaviocosta2110/getTheBlueBlocks/src/physics"
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
	playerScale   int32 = 3
	playerSizeX   int32 = 32
	playerSizeY   int32 = 32
)

var enemyArray []enemy.Enemy

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
	enemy := enemy.NewEnemy(200, 200, obstacleSpeed, playerSizeX, playerSizeY, playerScale, enemySprite)

	box := objects.NewBox(400, 400, 50, 50, rl.Brown)

	for !rl.WindowShouldClose() {
		update(player, enemy, box, screen)
		draw(player, enemy, box, *screen)
	}
}

func update(p *player.Player, e *enemy.Enemy, box *objects.Box, screen *screen.Screen) {
	if system.GameOverFlag {
		return
	}

	physics.TakeKnockback(&p.Object)
  physics.TakeKnockback(&box.Object)

	if p.Object.KnockbackX == 0 || p.Object.KnockbackY == 0 {
		p.CheckMovement(*screen)
	}

	if p.CheckAtk(e.Object) {
		newEnemy := enemy.NewEnemy(rand.Int31n(screen.Width), rand.Int31n(screen.Height), e.Speed, playerSizeX, playerSizeY, playerScale, e.Object.Sprite)
		*e = *newEnemy
	}
  if physics.CheckCollision(p.Object, box.Object){
    physics.ResolveCollision(&box.Object, &p.Object)
  }

	p.CheckKick(&box.Object)
	

	if e.CheckAtk(p.Object) {
		p.TakeDamage(1, e.Object)
		return
	}

	*e = enemy.MoveEnemyTowardPlayer(*p, *e, *screen)
}

func draw(p *player.Player, e *enemy.Enemy, box *objects.Box, s screen.Screen) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	if system.GameOverFlag {
		system.GameOver(&s)
		rl.EndDrawing()
		return
	}

	p.DrawPlayer()
	e.DrawEnemy()
	box.Draw()
	ui.DrawLife(s, p)

	rl.DrawText(fmt.Sprintf("Player: %d, %d", p.Object.X, p.Object.Y), 10, 10, 10, rl.Black)
	rl.DrawText(fmt.Sprintf("Enemy: %d, %d", e.Object.X, e.Object.Y), 10, 25, 10, rl.Black)

	rl.EndDrawing()
}

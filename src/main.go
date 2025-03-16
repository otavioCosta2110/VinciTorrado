package main

import (
	"fmt"
	"math/rand"
	"otaviocosta2110/getTheBlueBlocks/src/enemy"
	"otaviocosta2110/getTheBlueBlocks/src/physics"
	"otaviocosta2110/getTheBlueBlocks/src/player"
	"otaviocosta2110/getTheBlueBlocks/src/screen"
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

// Collider represents any object that can collide with others

func main() {
	screen := screen.NewScreen(windowWidth, windowHeight, "jogo poggers")
	rl.InitWindow(screen.Width, screen.Height, screen.Title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	playerSprite := system.Sprite{
		SpriteWidth:  playerSizeX,
		SpriteHeight: playerSizeY,
		Texture:      rl.LoadTexture("assets/player.png"),
	}

	enemySprite := rl.LoadTexture("assets/enemy.png")

	player := player.NewPlayer(screen.Width/2, screen.Height/2, playerSizeX, playerSizeY, 0, 4, playerScale, playerSprite)
	enemy := enemy.NewEnemy(50, 80, obstacleSpeed, playerSizeX, playerSizeY, playerScale, enemySprite)

	for !rl.WindowShouldClose() {
		update(player, enemy, screen)
		draw(player, enemy, *screen)
	}
}

func update(p *player.Player, e *enemy.Enemy, screen *screen.Screen) {
	if system.GameOverFlag {
		return
	}

	prevPX, prevEX := p.Object.X, e.Object.X
	prevPY, prevEY := p.Object.Y, e.Object.Y

	p.CheckMovement(*screen)
	if p.CheckAtk(e.Object) {
		newEnemy := enemy.NewEnemy(rand.Int31n(screen.Width), rand.Int31n(screen.Height), e.Speed, e.Object.Width, e.Object.Height, 1, e.Sprite)
		*e = *newEnemy
	}

	*e = enemy.MoveEnemyTowardPlayer(*p, *e, *screen)

	if physics.CheckCollision(p.Object, e.Object) {
		p.TakeDamage(1)
		p.Object.X, e.Object.X = prevPX, prevEX
		p.Object.Y, e.Object.Y = prevPY, prevEY
		return
	}
}

func draw(p *player.Player, e *enemy.Enemy, s screen.Screen) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	if system.GameOverFlag {
		system.GameOver(&s)
		rl.EndDrawing()
		return
	}

	p.DrawPlayer()
	e.DrawEnemy()
	ui.DrawLife(s, p)

	rl.DrawText(fmt.Sprintf("Player: %d, %d", p.Object.X, p.Object.Y), 10, 10, 10, rl.Black)
	rl.DrawText(fmt.Sprintf("Points: %d", p.Points), 10, 40, 10, rl.Black)
	rl.DrawText(fmt.Sprintf("Enemy: %d, %d", e.Object.X, e.Object.Y), 10, 25, 10, rl.Black)

	rl.EndDrawing()
}

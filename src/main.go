package main

import (
	"fmt"
	"math/rand"
	"otaviocosta2110/getTheBlueBlocks/src/enemy"
	"otaviocosta2110/getTheBlueBlocks/src/physics"
	"otaviocosta2110/getTheBlueBlocks/src/player"
	"otaviocosta2110/getTheBlueBlocks/src/points"
	"otaviocosta2110/getTheBlueBlocks/src/screen"
	"otaviocosta2110/getTheBlueBlocks/src/system"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowWidth   int32 = 1280
	windowHeight  int32 = 720
	obstacleSpeed int32 = 3
	playerScale   int32 = 3
)

var enemyArray []enemy.Enemy

func main() {
	screen := screen.NewScreen(windowWidth, windowHeight, "jogo poggers")
	rl.InitWindow(screen.Width, screen.Height, screen.Title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	playerSprite := rl.LoadTexture("assets/player.png")
	print("\n\n\n", playerSprite.Width, "\n", playerSprite.Height, "\n")

	enemySprite := rl.LoadTexture("assets/enemy.png")

	player := player.NewPlayer(900, 499, 32, 32, 0, 4, playerScale, playerSprite)
	enemy := enemy.NewEnemy(50, 80, obstacleSpeed, 32, 32, playerScale, enemySprite)

	numberOfPoints := rand.Intn(14-3+1) + 3
	pointsObject := make([]points.Point, 0, numberOfPoints)

	for range numberOfPoints {
		p := points.NewPoint(*screen)
		pointsObject = append(pointsObject, p)
	}

	for !rl.WindowShouldClose() {
		update(player, enemy, pointsObject, screen)
		draw(player, enemy, pointsObject, *screen)
	}
}

func update(p *player.Player, e *enemy.Enemy, pointsObject []points.Point, screen *screen.Screen) {
	if system.GameOverFlag {
		return
	}

	prevX, prevY := p.X, p.Y

	p.CheckMovement(*screen)
	if p.CheckAtk(e.X, e.Y, e.Width, e.Height) {
		// Lógica para derrotar o inimigo e gerar um novo
		newEnemy := enemy.NewEnemy(rand.Int31n(screen.Width), rand.Int31n(screen.Height), e.Speed, e.Width, e.Height, e.Scale, e.Sprite)
		*e = *newEnemy // Copia os valores do novo inimigo para o inimigo existente
	}

	*e = enemy.MoveEnemyTowardPlayer(*p, *e, *screen)

	for i := range pointsObject {
		point := pointsObject[i]
		if physics.CheckCollision(p.X, p.Y, point.X, point.Y, p.Width, p.Height) {
			p.Points++
			pointsObject = slices.Delete(pointsObject, i, i+1)
			break
		}
	}

	if physics.CheckCollision(p.X, prevY, e.X, e.Y, p.Width*p.Scale/2, p.Height*p.Scale/2) {
		system.GameOverFlag = true
		p.X = prevX
		return
	}
	if physics.CheckCollision(prevX, p.Y, e.X, e.Y, p.Width*p.Scale/2, p.Height*p.Scale/2) {
		system.GameOverFlag = true
		p.Y = prevY
		return
	}
}

func draw(p *player.Player, e *enemy.Enemy, pointsObject []points.Point, s screen.Screen) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	if system.GameOverFlag {
		system.GameOver(&s)
		rl.EndDrawing()
		return
	}

	p.DrawPlayer()
	e.DrawEnemy()

	for _, point := range pointsObject {
		point.DrawPoint()
	}

	rl.DrawText(fmt.Sprintf("Player: %d, %d", p.X, p.Y), 10, 10, 10, rl.Black)
	rl.DrawText(fmt.Sprintf("Points: %d", p.Points), 10, 40, 10, rl.Black)
	rl.DrawText(fmt.Sprintf("Enemy: %d, %d", e.X, e.Y), 10, 25, 10, rl.Black)

	rl.EndDrawing()
}

package main

import (
	"fmt"
	"math/rand"
	"otaviocosta2110/getTheBlueBlocks/src/enemy"
	"otaviocosta2110/getTheBlueBlocks/src/physics"
	"otaviocosta2110/getTheBlueBlocks/src/player"
	"otaviocosta2110/getTheBlueBlocks/src/points"
	"otaviocosta2110/getTheBlueBlocks/src/screen"

	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowWidth  int32 = 1920
	windowHeight int32 = 1080
	squareSize   int32 = 20
	obstacleSpeed int32 = 3 
)
var GameOver bool = false


func main() {
	screen := screen.NewScreen(windowWidth, windowHeight, "Raylib Go - Solid Object Collision")
	rl.InitWindow(screen.Width, screen.Width, screen.Title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	player := player.NewPlayer(900, 499, squareSize, squareSize, 0, 4)
	enemy := enemy.NewEnemy(50, 80, obstacleSpeed, squareSize, squareSize)

  numberOfPoints := rand.Intn(14 - 3 + 1 ) + 3
  pointsObject := make([]points.Point, 0, numberOfPoints)

  for range numberOfPoints {
    p := points.NewPoint(*screen)
    pointsObject = append(pointsObject, p)
  }

	for !rl.WindowShouldClose() {
		update(player, enemy, pointsObject, screen)
		draw(player, enemy, pointsObject)
	}
}

func update(p *player.Player, e *enemy.Enemy, pointsObject []points.Point, screen *screen.Screen) {
  if GameOver {
    return 
  }

	prevX, prevY := p.X, p.Y

	*p = p.CheckMovement(*screen)

	*e = enemy.MoveEnemyTowardPlayer(*p, *e)

  for i := range pointsObject {
    point := pointsObject[i]
    if physics.CheckCollision(p.X, p.Y, point.X, point.Y, p.Width, p.Height) {
      p.Points++
      pointsObject = slices.Delete((pointsObject), i, i+1) 
      break
    }
  }

	if physics.CheckCollision(p.X, prevY, e.X, e.Y, p.Width, p.Height) {
    GameOver = true
		p.X = prevX 
	}
	if physics.CheckCollision(prevX, p.Y, e.X, e.Y, p.Width, p.Height) {
    GameOver = true
		p.Y = prevY
	}
}

func gameOver() {
	text := "Game Over"
	textWidth := rl.MeasureText(text, 100)

	xPos := (windowWidth - textWidth) / 2
	yPos := windowHeight / 2 - 325

	rl.DrawText(text, xPos, yPos, 100, rl.Black)
}

func draw(p *player.Player, e *enemy.Enemy, pointsObject []points.Point) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	if GameOver {
		gameOver()
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


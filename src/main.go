package main

import (
	"otaviocosta2110/vincitorrado/src/enemy"
	"otaviocosta2110/vincitorrado/src/objects"
	"otaviocosta2110/vincitorrado/src/player"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"
	"otaviocosta2110/vincitorrado/src/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowWidth   int32  = 1280
	windowHeight  int32  = 720
	windowTitle   string = "Vinci Torrado"
	obstacleSpeed int32  = 2
	playerScale   int32  = 4
	playerSizeX   int32  = 32
	playerSizeY   int32  = 32
)

func main() {
	rl.InitWindow(windowWidth, windowHeight, windowTitle)

	buildings := rl.LoadTexture("assets/scenes/predio.png")
	buildings.Width *= playerScale
	buildings.Height *= playerScale

	screen := screen.NewScreen(windowWidth, windowHeight, buildings.Width, buildings.Height, windowTitle)

	chao := rl.LoadTexture("assets/scenes/chao.png")
	chao.Width *= playerScale
	chao.Height *= playerScale

	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	playerSprite := sprites.Sprite{
		SpriteWidth:  playerSizeX,
		SpriteHeight: playerSizeY,
		Texture:      rl.LoadTexture("assets/player/player.png"),
	}

	player := player.NewPlayer(screen.Width/2, screen.Height/2, playerSizeX, playerSizeY, 2, playerScale, playerSprite)

	boxes := []*objects.Box{
		objects.NewBox(200, screen.Height-100, 50, 50),
	}

	// nesse arquivos tem informcoes sobre os inimigos, como a posicao inicial, vida, forca, etc
	enemies, err := enemy.LoadEnemiesFromJSON("assets/enemies/enemyInfo/1_00 enemyInfo.json", playerScale)
	if err != nil {
		panic(err)
	}

	enemyManager := enemy.EnemyManager{}
	for _, e := range enemies {
		enemyManager.AddEnemy(e)
	}

	screen.InitCamera(player.Object.X, player.Object.Y)

	for !rl.WindowShouldClose() {
		update(player, &enemyManager, screen, boxes)
		draw(player, &enemyManager, *screen, chao, buildings, boxes)
	}
}

func update(p *player.Player, em *enemy.EnemyManager, screen *screen.Screen, boxes []*objects.Box) {
	if system.GameOverFlag {
		return
	}

	p.CheckMovement(*screen)

	for _, box := range boxes {
		p.CheckKick(box)
	}

	for _, box := range boxes {
		box.Update([]system.Object{p.GetObject()}, screen, em)
	}

	em.Update(p, *screen)
	p.Update(em, *screen)
	canAdvance := len(em.ActiveEnemies) <= 0

	screen.UpdateCamera(p.Object.X, p.Object.Y, canAdvance)
}

func draw(p *player.Player, em *enemy.EnemyManager, s screen.Screen, chao rl.Texture2D, buildings rl.Texture2D, boxes []*objects.Box) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode2D(s.Camera)

	drawTiledBackground(chao, s.Camera, s.Width, s.Height)
	drawBuildings(buildings)

	for _, box := range boxes {
		box.Draw()
	}

	em.Draw()
	p.Draw()

	rl.EndMode2D()

	if system.GameOverFlag {
		system.GameOver(&s)
	}

	ui.DrawLife(s, p)
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

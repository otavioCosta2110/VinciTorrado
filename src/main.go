package main

import (
	"otaviocosta2110/vincitorrado/src/audio"
	"otaviocosta2110/vincitorrado/src/enemy"
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/objects"
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/player"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"
	"otaviocosta2110/vincitorrado/src/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowWidth  int32  = 1280
	windowHeight int32  = 720
	windowTitle  string = "Vinci Torrado"
	playerScale  int32  = 4
	playerSizeX  int32  = 32
	playerSizeY  int32  = 32
)

func main() {
	rl.InitWindow(windowWidth, windowHeight, windowTitle)

	buildings := rl.LoadTexture("assets/scenes/predio.png")
	buildings.Width *= playerScale
	buildings.Height *= playerScale

	chao := rl.LoadTexture("assets/scenes/chao.png")
	chao.Width *= playerScale
	chao.Height *= playerScale

	rl.InitAudioDevice()
	audio.LoadSounds()
	defer rl.CloseAudioDevice()
	defer audio.UnloadSounds()

	screen := screen.NewScreen(windowWidth, windowHeight, buildings.Width, buildings.Height, windowTitle)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	rl.SetExitKey(0)

	playerSprite := sprites.Sprite{
		SpriteWidth:  playerSizeX,
		SpriteHeight: playerSizeY,
		Texture:      rl.LoadTexture("assets/player/player.png"),
	}
	player := player.NewPlayer(screen.Width/2, screen.Height/2, playerSizeX, playerSizeY, 2, playerScale, playerSprite, screen)
	menu := ui.NewMenu(player, &playerSprite)

	items, err := equipment.LoadItemsFromJSON("assets/items/items.json")
	if err != nil {
		panic("Failed to load items: " + err.Error())
	}

	enemies, err := enemy.LoadEnemiesFromJSON(
		"assets/enemies/enemyInfo/1_00 enemyInfo.json",
		playerScale,
	)
	if err != nil {
		panic("Failed to load enemies: " + err.Error())
	}

	trashCans, err := objects.LoadTrashCansFromJSON("assets/props/props.json", items)
	if err != nil {
		panic("Failed to load trash cans: " + err.Error())
	}

	boxes := []*objects.Box{
		objects.NewBox(200, screen.Height-100, 50, 50),
	}

	var kickables []physics.Kickable
	for _, box := range boxes {
		kickables = append(kickables, box)
	}
	for _, trash := range trashCans {
		kickables = append(kickables, trash)
	}

	enemyManager := &enemy.EnemyManager{}
	for _, e := range enemies {
		enemyManager.AddEnemy(e)
	}

	screen.InitCamera(player.Object.X, player.Object.Y)

	for !rl.WindowShouldClose() {
		menu.Update()

		if !menu.IsVisible {
			update(player, enemyManager, screen, kickables, &items)
		}
		draw(player, enemyManager, *screen, chao, buildings, boxes, items, trashCans, *menu)
	}
}

func update(p *player.Player, em *enemy.EnemyManager, screen *screen.Screen, kickables []physics.Kickable, items *[]*equipment.Equipment) {
	if kicked := p.CheckKick(kickables, items); kicked {
		// som de chute
		audio.PlayKick()
	}

	if system.GameOverFlag {
		return
	}

	p.CheckMovement(*screen)

	for _, obj := range kickables {
		if box, ok := obj.(*objects.Box); ok {
			box.Update([]system.Object{p.GetObject()}, screen, em, nil)
		}
	}

	for _, e := range em.Enemies {
		if e.Object.Destroyed && e.Drop != nil && !e.DropCollected {
			dropWidth := int32(32 * e.Object.Scale)
			dropHeight := int32(32 * e.Object.Scale)
			dropY := e.Object.Y - 20

			dropBox := system.Object{
				X:      e.Object.X,
				Y:      dropY,
				Width:  dropWidth / 2,
				Height: dropHeight / 2,
			}
			e.Drop.IsDropped = true

			playerObj := p.GetObject()
			if physics.CheckCollision(playerObj, dropBox) {
				menu_select_sound := rl.LoadSound("assets/sounds/collect_item.mp3")
				rl.PlaySound(menu_select_sound)
				p.AddToInventory(e.Drop)
				e.DropCollected = true
				e.Drop.IsDropped = false
			}
		}
	}

	for _, item := range *items {
		if item.IsDropped {
			itemBox := system.Object{
				X:      item.Object.X,
				Y:      item.Object.Y,
				Width:  item.Object.Width / 2,
				Height: item.Object.Height / 2,
			}

			if physics.CheckCollision(p.GetObject(), itemBox) {
				p.AddToInventory(item)
				item.IsDropped = false
				collectSound := rl.LoadSound("assets/sounds/collect_item.mp3")
				rl.PlaySound(collectSound)
			}
		}
	}

	em.Update(p, *screen)
	p.Update(em, *screen)
	canAdvance := len(em.ActiveEnemies) <= 0
	screen.UpdateCamera(p.Object.X, p.Object.Y, canAdvance)

	return
}

func draw(p *player.Player, em *enemy.EnemyManager, s screen.Screen, chao rl.Texture2D, buildings rl.Texture2D, boxes []*objects.Box, items []*equipment.Equipment, trashCans []*objects.TrashCan, menu ui.Menu) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode2D(s.Camera)
	drawTiledBackground(chao, s.Camera, s.Width, s.Height)
	drawBuildings(buildings)

	for _, item := range items {
		if item.IsDropped {
			item.DrawAnimated(&item.Object)
		}
	}

	for _, box := range boxes {
		box.Draw()
	}

	for _, trash := range trashCans {
		trash.Draw()
	}

	em.Draw()
	p.Draw()

	rl.EndMode2D()

	if system.GameOverFlag {
		system.GameOver(&s)
	}

	ui.DrawLife(s, p)
	menu.Draw()

	rl.EndDrawing()
}

func drawTiledBackground(texture rl.Texture2D, camera rl.Camera2D, screenWidth, screenHeight int32) {
	texWidth := texture.Width
	texHeight := texture.Height

	visibleStartX := int32(camera.Target.X) - screenWidth/2 - texWidth
	visibleStartY := int32(camera.Target.Y) - screenHeight/2 - texHeight
	visibleEndX := int32(camera.Target.X) + screenWidth/2 + texWidth
	visibleEndY := int32(camera.Target.Y) + screenHeight/2 + texHeight

	for y := visibleStartY / texHeight; y <= visibleEndY/texHeight+1; y++ {
		for x := visibleStartX / texWidth; x <= visibleEndX/texWidth+1; x++ {
			rl.DrawTexture(texture, x*texWidth, y*texHeight, rl.White)
		}
	}
}

func drawBuildings(texture rl.Texture2D) {
	rl.DrawTexture(texture, 0, 0, rl.White)
}

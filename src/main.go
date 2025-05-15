package main

import (
	"otaviocosta2110/vincitorrado/src/audio"
	"otaviocosta2110/vincitorrado/src/enemy"
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/player"
	"otaviocosta2110/vincitorrado/src/props"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"
	"otaviocosta2110/vincitorrado/src/ui"
	"otaviocosta2110/vincitorrado/src/weapon"
	"time"

	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowWidth  int32  = 1280
	windowHeight int32  = 720
	windowTitle  string = "Vinci Torrado"
	playerScale  int32  = 4
	playerSizeX  int32  = 32
	playerSizeY  int32  = 32

	// feature flags
	oneHealthEnemies bool = true
	enableMusic      bool = true
	enableSoundFxs   bool = true
)

type GameState struct {
	Player       *player.Player
	EnemyManager *enemy.EnemyManager
	Screen       *screen.Screen
	Kickables    []physics.Kickable
	Items        []*equipment.Equipment
	Props        []*props.Prop
	Weapons      []*weapon.Weapon
	Menu         ui.Menu
	Music        *string
	Doors        []*props.Door
	MapManager   *system.MapManager
	Buildings    rl.Texture2D
	Chao         rl.Texture2D
	FloorPath    string
}

func main() {
	rl.InitWindow(windowWidth, windowHeight, windowTitle)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	rl.SetExitKey(0)
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	buildings := loadScaledTexture("assets/scenes/predio.png", playerScale)
	chao := loadScaledTexture("assets/scenes/chao.png", playerScale)

	mapManager := system.NewMapManager()
	mapManager.Maps["bar"] = &system.GameMap{
		Buildings:    "assets/scenes/continuacao_bar.jpg",
		Floor:        "assets/scenes/chao_bar.png",
		EnemiesPath:  "assets/enemies/bar_enemies.json", //novo json ou só algo q diga q ele é do bar
		PropsPath:    "assets/props/bar_props.json",     //tbm
		PlayerStartX: 100,
		PlayerStartY: 100,
	}

	if enableSoundFxs {
		audio.LoadSounds()
	}
	if enableMusic {
		audio.PlayMissionMusic()
	}
	defer audio.UnloadSounds()

	screen := screen.NewScreen(windowWidth, windowHeight, buildings.Width, buildings.Height, windowTitle)

	playerSprite := sprites.Sprite{
		SpriteWidth:  playerSizeX,
		SpriteHeight: playerSizeY,
		Texture:      rl.LoadTexture("assets/player/player.png"),
	}
	player := player.NewPlayer(screen.Width/2, screen.Height/2, playerSizeX, playerSizeY, 4, playerScale, playerSprite, screen)
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

	enemyManager := &enemy.EnemyManager{}
	for _, e := range enemies {
		if oneHealthEnemies {
			e.Health = 0
		}
		enemyManager.AddEnemy(e)
	}

	weapons, err := weapon.LoadWeaponsFromJSON("assets/weapons/1_00 weapon.json")
	if err != nil {
		panic("Failed to load weapons: " + err.Error())
	}

	props, doors, err := props.LoadPropsFromJSON("assets/props/props.json", items)
	if err != nil {
		panic("Failed to load props: " + err.Error())
	}

	var kickables []physics.Kickable
	for _, prop := range props {
		kickables = append(kickables, prop)
	}

	screen.InitCamera(player.Object.X, player.Object.Y)

	music := "mission1"
	gameState := GameState{
		Player:       player,
		EnemyManager: enemyManager,
		Screen:       screen,
		Kickables:    kickables,
		Items:        items,
		Props:        props,
		Weapons:      weapons,
		Menu:         *menu,
		Music:        &music,
		Buildings:    buildings,
		Chao:         chao,
		Doors:        doors,
		MapManager:   mapManager,
	}

	gameLoop(&gameState)
}

func gameLoop(gs *GameState) {
	for !rl.WindowShouldClose() {
		audio.UpdateMusic(*gs.Music)
		gs.Menu.Update()

		if !gs.Menu.IsVisible {
			update(gs)
		}
		draw(gs)
	}
}

func update(gs *GameState) {
	if system.GameOverFlag {
		return
	}
	if gs.Player.IsKicking && time.Since(gs.Player.LastKickTime) > 200*time.Millisecond {
		gs.Player.IsKicking = false
		gs.Player.Object.FrameY = 0
		gs.Player.Object.FrameX = 0
	}
	if gs.Player.IsAttacking && time.Since(gs.Player.LastAttackTime) > 200*time.Millisecond {
		gs.Player.IsAttacking = false
		gs.Player.Object.FrameX = 0
		gs.Player.Object.FrameY = 0
	}
	for i := range gs.Weapons {
		weapon := gs.Weapons[i]
		if weapon.IsDropped {
			weapon.DrawAnimated()
			dropBox := weapon.GetDropCollisionBox()
			if physics.CheckCollision(gs.Player.GetObject(), dropBox) {
				weapon.IsDropped = false
				weapon.IsEquipped = true
				gs.Player.PickUp(*weapon)
			}
		}
	}
	gs.Player.CheckKick(gs.Kickables, &gs.Items)

	for _, e := range gs.EnemyManager.Enemies {
		if e.Weapon != nil && e.Weapon.IsDropped {
			weapon := e.Weapon.Clone()
			gs.Weapons = append(gs.Weapons, weapon)
			e.Weapon = nil
		}
		if e.Object.Destroyed && e.Drop != nil && !e.DropCollected {
			dropBox := e.GetDropCollisionBox()
			e.Drop.IsDropped = true
			if physics.CheckCollision(gs.Player.GetObject(), dropBox) {
				rl.PlaySound(audio.CollectItemSound)
				gs.Player.AddToInventory(e.Drop)
				e.DropCollected = true
				e.Drop.IsDropped = false
			}
		}
	}

	for i := range gs.Items {
		item := gs.Items[i]
		if item.IsDropped {
			itemBox := item.GetObject()
			if physics.CheckCollision(gs.Player.GetObject(), *itemBox) {
				gs.Player.AddToInventory(item)
				item.IsDropped = false
				rl.PlaySound(audio.CollectItemSound)
				gs.Items = slices.Delete(gs.Items, i, i+1)
				break
			}
		}
	}

	gs.EnemyManager.Update(gs.Player, *gs.Screen, gs.Music)
	gs.Player.Update(gs.EnemyManager, *gs.Screen)
	canAdvance := len(gs.EnemyManager.ActiveEnemies) <= 0
	gs.Screen.UpdateCamera(gs.Player.Object.X, gs.Player.Object.Y, canAdvance)

	for _, door := range gs.Doors {
		if door.CheckTransition(gs.Player.GetObject(), canAdvance) {
			transitionMap(gs, door.NextMap)
			break
		}
	}
}

func draw(gs *GameState) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode2D(gs.Screen.Camera)
	drawTiledBackground(gs.Chao, gs.Screen.Camera, gs.Screen.Width, gs.Screen.Height)
	drawBuildings(gs.Buildings)

	for _, prop := range gs.Props {
		prop.Draw()
	}

	for _, door := range gs.Doors {
		door.Draw()
	}

	gs.EnemyManager.Draw()
	gs.Player.Draw()

	for _, item := range gs.Items {
		if item.IsDropped {
			item.DrawAnimated(&item.Object)
		}
	}

	for _, weapon := range gs.Weapons {
		if weapon.IsDropped {
			weapon.DrawAnimated()
		}
	}

	rl.EndMode2D()

	if system.GameOverFlag {
		system.GameOver(gs.Screen)
	}

	ui.DrawLife(*gs.Screen, gs.Player)
	gs.Menu.Draw()

	rl.EndDrawing()
}

func loadScaledTexture(path string, scale int32) rl.Texture2D {
	texture := rl.LoadTexture(path)
	texture.Width *= scale
	texture.Height *= scale
	return texture
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

func transitionMap(gs *GameState, mapName string) {
	rl.UnloadTexture(gs.Buildings)
	rl.UnloadTexture(gs.Chao)

	gs.Player.Object.FrameX = 0
	gs.Player.Object.FrameY = 0
	gs.Player.IsKicking = false
	gs.Player.LastKickTime = time.Now().Add(-time.Hour)

	newMap := gs.MapManager.Maps[mapName]

	gs.Buildings = loadScaledTexture(newMap.Buildings, playerScale)
	gs.Chao = loadScaledTexture(newMap.Floor, playerScale)

	enemies, _ := enemy.LoadEnemiesFromJSON(newMap.EnemiesPath, playerScale)
	gs.EnemyManager = &enemy.EnemyManager{}
	for _, e := range enemies {
		gs.EnemyManager.AddEnemy(e)
	}

	props, doors, _ := props.LoadPropsFromJSON(newMap.PropsPath, gs.Items)
	gs.Props = props
	gs.Doors = doors

	gs.Player.Object.X = newMap.PlayerStartX
	gs.Player.Object.Y = newMap.PlayerStartY
	gs.Screen.ResetCamera()

	gs.Kickables = nil
	for _, prop := range gs.Props {
		gs.Kickables = append(gs.Kickables, prop)
	}
}

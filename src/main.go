package main

import (
	"otaviocosta2110/vincitorrado/src/audio"
	"otaviocosta2110/vincitorrado/src/cutscene"
	"otaviocosta2110/vincitorrado/src/enemy"
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/girlfriend"
	"otaviocosta2110/vincitorrado/src/maps"
	"otaviocosta2110/vincitorrado/src/objects"
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
	oneHealthEnemies bool   = false
	enableMusic      bool   = false
	enableSoundFxs   bool   = true
	skipCutscenes    bool   = true
	startingMap      string = "bar"
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
	Cutscene     *cutscene.Cutscene
	Girlfriend   *girlfriend.Girlfriend
	Doors        []*props.Door
	MapManager   *maps.MapManager
	Buildings    *rl.Texture2D
	Chao         rl.Texture2D
	FloorPath    string
	CurrentMap   string
}

func main() {
	rl.InitWindow(windowWidth, windowHeight, windowTitle)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	rl.SetExitKey(0)
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	mapManager := maps.NewMapManager()
	mapManager.Maps["city"] = &maps.GameMap{
		Buildings:    "assets/scenes/predio.png",
		Floor:        "assets/scenes/chao.png",
		EnemiesPath:  "assets/enemies/enemyInfo/1_00 enemyInfo.json",
		PropsPath:    "assets/props/props.json",
		PlayerStartX: -100,
		PlayerStartY: windowHeight/2 + 50,
	}
	mapManager.Maps["bar"] = &maps.GameMap{
		Buildings:    "assets/scenes/bar.png",
		Floor:        "assets/scenes/chao_bar.png",
		EnemiesPath:  "assets/enemies/enemyInfo/2_00 enemyInfo.json",
		PropsPath:    "assets/props/bar_props.json",
		PlayerStartX: 100,
		PlayerStartY: 650,
	}

	currentMap := mapManager.Maps[startingMap]

	buildings := system.LoadScaledTexture(currentMap.Buildings, playerScale)
	chao := system.LoadScaledTexture(currentMap.Floor, playerScale)

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

	player := player.NewPlayer(currentMap.PlayerStartX, currentMap.PlayerStartY, playerSizeX, playerSizeY, 4, playerScale, playerSprite, screen)
	weaponSprite := sprites.Sprite{
		SpriteWidth:  playerSizeX,
		SpriteHeight: playerSizeY,
		Texture:      rl.LoadTexture("assets/weapons/flowers.png"),
	}

	stats := objects.Stats{
		Damage: 6,
	}

	playerWeapon := &weapon.Weapon{
		Object: &system.Object{
			X:      player.Object.X,
			Y:      player.Object.Y,
			Width:  32 * playerScale,
			Height: 32 * playerScale,
			Scale:  playerScale,
			Sprite: weaponSprite,
		},
		IsDropped: true,
		Stats:     stats,
		HitboxX:   30,
		HitboxY:   0,
		OffsetX:   9,
		OffsetY:   0,
		Health:    3,
	}
	player.PickUp(*playerWeapon)
	menu := ui.NewMenu(player, &playerSprite)

	items, err := equipment.LoadItemsFromJSON("assets/items/items.json")
	if err != nil {
		panic("Failed to load items: " + err.Error())
	}

	weapons, err := weapon.LoadWeaponsFromJSON("assets/weapons/1_00 weapon.json")
	if err != nil {
		panic("Failed to load weapons: " + err.Error())
	}

	screen.InitCamera(player.Object.X, player.Object.Y)

	gSprite := sprites.Sprite{
		SpriteWidth:  playerSizeX,
		SpriteHeight: playerSizeY,
		Texture:      rl.LoadTexture("assets/player/girlfriend.png"),
	}
	g := girlfriend.New(gSprite, 1000, player.Object.Y, 4)

	enemies, err := enemy.LoadEnemiesFromJSON(currentMap.EnemiesPath, playerScale)
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

	props, doors, err := props.LoadPropsFromJSON(currentMap.PropsPath, items)
	if err != nil {
		panic("Failed to load props: " + err.Error())
	}

	var kickables []physics.Kickable
	for _, prop := range props {
		kickables = append(kickables, prop)
	}

	gameState := GameState{
		Player:       player,
		EnemyManager: enemyManager,
		Screen:       screen,
		Kickables:    kickables,
		Items:        items,
		Props:        props,
		Weapons:      weapons,
		Menu:         *menu,
		Girlfriend:   g,
		Buildings:    &buildings,
		Chao:         chao,
		Doors:        doors,
		MapManager:   mapManager,
		CurrentMap:   startingMap,
	}

	transitionMap(&gameState, startingMap)
	gameLoop(&gameState)
}

func gameLoop(gs *GameState) {
	for !rl.WindowShouldClose() {
		audio.UpdateMusic(*gs.Music)
		gs.Menu.Update()

		if gs.Cutscene != nil && gs.Cutscene.IsPlaying() {
			gs.Cutscene.Update()
		} else if !gs.Menu.IsVisible {
			update(gs)
		}

		draw(gs)
	}
}

func update(gs *GameState) {
	if system.GameOverFlag || gs.Cutscene.IsPlaying() {
		return
	}
	if gs.Player.IsKicking && time.Since(gs.Player.LastKickTime) > 200*time.Millisecond {
		gs.Player.IsKicking = false
		gs.Player.Object.FrameY = 0
		gs.Player.Object.FrameX = 0
	}
	if gs.Player.IsAttacking && time.Since(gs.Player.LastAttackTime) > 400*time.Millisecond {
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

	gs.EnemyManager.Update(gs.Player, *gs.Screen, gs.Music, gs.Props, gs.Buildings)
	gs.Player.Update(gs.EnemyManager, *gs.Screen)

	activeEnemies := []*enemy.Enemy{}
	for _, enemy := range gs.EnemyManager.ActiveEnemies {
		if enemy.EnemyType != "mafia_boss" {
			activeEnemies = append(activeEnemies, enemy)
		}
	}

	canAdvance := len(activeEnemies) <= 0
	gs.Screen.UpdateCamera(gs.Player.Object.X, gs.Player.Object.Y, canAdvance)

	for _, door := range gs.Doors {
		if door.CheckTransition(gs.Player.GetObject(), canAdvance) {
			transitionMap(gs, door.NextMap)
			break
		}
	}

	for _, e := range gs.EnemyManager.Enemies {
		if e.EnemyType == "gf_monster" {
			e.UpdateGirlfriendHealth()
		}
	}
}

func draw(gs *GameState) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode2D(gs.Screen.Camera)
	drawTiledBackground(gs.Chao, gs.Screen.Camera, gs.Screen.Width, gs.Screen.Height)
	drawBuildings(*gs.Buildings)

	gs.EnemyManager.DrawDead()
	for _, prop := range gs.Props {
		prop.Draw()
	}
	gs.EnemyManager.Draw()
	gs.Player.Draw()

	for _, item := range gs.Items {
		if item.IsDropped {
			item.DrawAnimated(&item.Object)
		}
	}
	if gs.Girlfriend.IsActive() {
		gs.Girlfriend.Draw()
	}
	for _, door := range gs.Doors {
		door.Draw()
	}

	for _, weapon := range gs.Weapons {
		if weapon.IsDropped {
			weapon.DrawAnimated()
		}
	}

	var girlfriendBoss *enemy.Enemy
	for _, e := range gs.EnemyManager.Enemies {
		if e.EnemyType == "gf_monster" {
			girlfriendBoss = e
			break
		}
	}
	ui.DrawBossHealthBar(girlfriendBoss, gs.Screen.Width)

	rl.EndMode2D()

	if system.GameOverFlag {
		system.GameOver(gs.Screen)
	}

	ui.DrawLife(*gs.Screen, gs.Player)
	gs.Menu.Draw()

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

func transitionMap(gs *GameState, mapName string) {
	if gs.Buildings.ID != 0 {
		rl.UnloadTexture(*gs.Buildings)
	}
	if gs.Chao.ID != 0 {
		rl.UnloadTexture(gs.Chao)
	}
	rl.UnloadTexture(*gs.Buildings)
	rl.UnloadTexture(gs.Chao)

	gs.Player.Object.FrameX = 0
	gs.Player.Object.FrameY = 0
	gs.Player.IsKicking = false
	gs.Player.LastKickTime = time.Now().Add(-time.Hour)

	newMap := gs.MapManager.Maps[mapName]
	gs.CurrentMap = mapName

	*gs.Buildings = system.LoadScaledTexture(newMap.Buildings, playerScale)
	gs.Chao = system.LoadScaledTexture(newMap.Floor, playerScale)

	enemies, err := enemy.LoadEnemiesFromJSON(newMap.EnemiesPath, playerScale)
	if err != nil {
		panic("Failed to load enemies: " + err.Error())
	}

	gs.EnemyManager = &enemy.EnemyManager{
		CurrentMap: mapName,
	}
	for _, e := range enemies {
		if oneHealthEnemies {
			e.Health = 0
		}
		gs.EnemyManager.AddEnemy(e)
	}
	gs.Cutscene = cutscene.NewCutscene()

	switch gs.CurrentMap {
	case "city":
		music := "mission1"
		gs.Music = &music
		if !skipCutscenes {
			gs.Cutscene.IntroCutscenes(gs.Player, gs.Girlfriend, gs.EnemyManager)
			gs.Cutscene.Start()
		}
	case "bar":
		music := "mission2"
		if !skipCutscenes {
			gs.Cutscene.BarIntroCutscene(gs.Player, gs.Girlfriend, gs.EnemyManager)
			gs.Cutscene.Start()
		}
		gs.Music = &music
		audio.StopMusic()
		audio.PlayMission2Music()
	}

	props, doors, err := props.LoadPropsFromJSON(newMap.PropsPath, gs.Items)
	if err != nil {
		panic("Failed to load props: " + err.Error())
	}
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

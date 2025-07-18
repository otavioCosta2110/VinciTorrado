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
	"sort"
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
	MenuState           = iota
	EndingState
	GameOverState
	GameRunningState

	// feature flags
	playerInfiniteLife bool   = false
	oneHealthEnemies   bool   = false
	enableMusic        bool   = true
	enableSoundFxs     bool   = true
	skipCutscenes      bool   = true
	startingMap        string = "city" // "city", "bar", "transition", "lab", "gf_monster"
)

type GameState struct {
	CurrentState int
	NeedsRestart bool
	StartMenu    *ui.StartMenu
	Ending       *ui.StartMenu
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
	mapManager.Maps["transition"] = &maps.GameMap{
		Buildings:    "assets/scenes/predio_exploded.png",
		Floor:        "assets/scenes/chao.png",
		EnemiesPath:  "assets/enemies/enemyInfo/transition_enemies.json",
		PropsPath:    "assets/props/transition_props.json",
		PlayerStartX: 100,
		PlayerStartY: 350,
	}
	mapManager.Maps["lab"] = &maps.GameMap{
		Buildings:    "assets/scenes/crazy_lab.png",
		Floor:        "assets/scenes/chao_lab.png",
		EnemiesPath:  "assets/enemies/enemyInfo/3_00 enemyInfo.json",
		PropsPath:    "assets/props/lab_props.json",
		PlayerStartX: 100,
		PlayerStartY: 650,
	}
	mapManager.Maps["gf_monster"] = &maps.GameMap{
		Buildings:    "assets/scenes/gf_monster_scene.png",
		Floor:        "assets/scenes/chao_lab.png",
		EnemiesPath:  "assets/enemies/enemyInfo/gf_monster.json",
		PropsPath:    "assets/props/gf_monster_props.json",
		PlayerStartX: 100,
		PlayerStartY: 650,
	}
	mapManager.Maps["last"] = &maps.GameMap{
		Buildings:    "assets/scenes/last_scene.png",
		Floor:        "assets/scenes/chao_lab.png",
		EnemiesPath:  "assets/enemies/enemyInfo/last_enemy.json",
		PropsPath:    "assets/props/last_props.json",
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

	startMenu := ui.NewStartMenu()
	endingMenu := ui.NewEndingMenu()
	defer rl.UnloadTexture(startMenu.BgTexture)
	defer rl.UnloadTexture(endingMenu.BgTexture)

	screen := screen.NewScreen(windowWidth, windowHeight, buildings.Width, buildings.Height, windowTitle)

	playerSprite := sprites.Sprite{
		SpriteWidth:  playerSizeX,
		SpriteHeight: playerSizeY,
		Texture:      rl.LoadTexture("assets/player/player.png"),
	}

	var playerHealth int32

	if playerInfiniteLife {
		playerHealth = 9999
	} else {
		playerHealth = 5
	}

	player := player.NewPlayer(currentMap.PlayerStartX, currentMap.PlayerStartY, playerSizeX, playerSizeY, 4, playerHealth, playerScale, playerSprite, screen)
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
		StartMenu:    startMenu,
		Ending:       endingMenu,
	}

	transitionMap(&gameState, startingMap)

	gameLoop(&gameState)
}

func gameLoop(gs *GameState) {
	gs.CurrentState = MenuState

	for !rl.WindowShouldClose() {
		switch gs.CurrentState {
		case EndingState:
			audio.UpdateMusic(*gs.Music)
			gs.Ending.DrawEndingMenu(gs.Screen)
		case MenuState:
			if gs.StartMenu.DrawStartMenu(gs.Screen) {
				gs.CurrentState = GameRunningState
				if enableMusic {
					audio.PlayMissionMusic()
				}
			}

		case GameRunningState:
			audio.UpdateMusic(*gs.Music)
			gs.Menu.Update()

			if gs.Cutscene != nil && gs.Cutscene.IsPlaying() {
				gs.Cutscene.Update()
			}
			update(gs)

			draw(gs)

			if system.GameOverFlag {
				gs.CurrentState = GameOverState
			}

		case GameOverState:
			system.DrawGameOver(gs.Screen)
			if rl.IsKeyPressed(rl.KeyR) {
				gs.RestartGame()
			}
			if rl.IsKeyPressed(rl.KeyEscape) {
				rl.CloseWindow()
				break
			}

			if gs.Cutscene != nil && gs.Cutscene.IsPlaying() {
				gs.Cutscene.Update()
			}
			update(gs)

			draw(gs)
		}
	}
}

func update(gs *GameState) {
	if system.GameOverFlag || gs.Cutscene.IsPlaying() {
		return
	}

	gs.EnemyManager.Update(gs.Player, *gs.Screen, gs.Music, gs.Props, gs.Buildings, &gs.Menu.IsVisible)

	if gs.Menu.IsVisible {
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

}

func draw(gs *GameState) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode2D(gs.Screen.Camera)
	drawTiledBackground(gs.Chao, gs.Screen.Camera, gs.Screen.Width, gs.Screen.Height)
	drawBuildings(*gs.Buildings)

	renderables := collectRenderables(gs)

	sort.Slice(renderables, func(i, j int) bool {
		return renderables[i].Layer < renderables[j].Layer
	})

	gs.EnemyManager.DrawDead()

	for _, obj := range renderables {
		obj.Draw()
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
		gs.Player.Object.Destroyed = true
		system.DrawGameOver(gs.Screen)
	}

	ui.DrawLife(*gs.Screen, gs.Player)
	gs.Menu.Draw()
	if gs.Cutscene != nil && gs.Cutscene.IsPlaying() && gs.Cutscene.DrawBlackScreen {
		rl.DrawRectangle(0, 0, gs.Screen.Width, gs.Screen.Height, rl.Black)
	}

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

func (gs *GameState) RestartGame() {
	gs.Player.Reset()

	currentMap := gs.MapManager.Maps[gs.CurrentMap]
	enemies, err := enemy.LoadEnemiesFromJSON(currentMap.EnemiesPath, playerScale)
	if err != nil {
		panic("Failed to load enemies: " + err.Error())
	}

	gs.EnemyManager = &enemy.EnemyManager{}
	for _, e := range enemies {
		if oneHealthEnemies {
			e.Health = 0
		}
		gs.EnemyManager.AddEnemy(e)
	}

	for _, kik := range gs.Kickables {
		kik.Reset()
	}

	for _, weapon := range gs.Weapons {
		if weapon.IsDropped {
			weapon.IsDropped = false
		}
	}

	system.GameOverFlag = false
	gs.NeedsRestart = false

	gs.Screen.ResetCamera()
	gs.Player.Object.X = currentMap.PlayerStartX
	gs.Player.Object.Y = currentMap.PlayerStartY

	gs.CurrentState = GameRunningState
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
	gs.Player.RecordInitialEquipment()

	newMap := gs.MapManager.Maps[mapName]
	gs.CurrentMap = mapName

	*gs.Buildings = system.LoadScaledTexture(newMap.Buildings, playerScale)
	gs.Chao = system.LoadScaledTexture(newMap.Floor, playerScale)

	gs.Screen.ScenaryWidth = gs.Buildings.Width
	gs.Screen.ScenaryWidth = gs.Buildings.Width

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

	props, doors, err := props.LoadPropsFromJSON(newMap.PropsPath, gs.Items)
	if err != nil {
		panic("Failed to load props: " + err.Error())
	}
	gs.Props = props
	gs.Doors = doors

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
	case "transition":
		music := "mission1"
		if !skipCutscenes {
			gs.Cutscene.Transition(gs.Player, gs.Girlfriend, gs.EnemyManager)
			gs.Cutscene.Start()
		}
		gs.Music = &music
		audio.StopMusic()
		audio.PlayMission2Music()
	case "lab":
		music := "mission3"
		gs.Music = &music
		gs.Girlfriend.SetActive(false)
		audio.StopMusic()
		audio.PlayMission3Music()
	case "gf_monster":
		music := "gf_battle"
		if !skipCutscenes {
			gs.Cutscene.GfMonster(gs.Player, gs.Girlfriend, gs.EnemyManager, gs.Props)
			gs.Cutscene.Start()
		}
		gs.Music = &music
		gs.Girlfriend.SetActive(false)
		audio.StopMusic()
		audio.PlayGfBattleMusic()
	case "last":
		music := "ending"
		if !skipCutscenes {
			gs.Cutscene.Doctor(gs.Player, gs.Girlfriend, gs.EnemyManager, &gs.CurrentState)
			gs.Cutscene.Start()
		}
		gs.Music = &music
		gs.Girlfriend.SetActive(false)
		audio.StopMusic()
		audio.PlayEndingMusic()
	}

	gs.Player.Object.X = newMap.PlayerStartX
	gs.Player.Object.Y = newMap.PlayerStartY
	gs.Screen.ResetCamera()

	gs.Kickables = nil
	for _, prop := range gs.Props {
		gs.Kickables = append(gs.Kickables, prop)
	}
}

type renderableObject struct {
	Object *system.Object
	Draw   func()
	Layer  int32
}

func collectRenderables(gs *GameState) []renderableObject {
	var renderables []renderableObject

	for _, prop := range gs.Props {
		renderables = append(renderables, renderableObject{
			Object: &prop.Object,
			Draw:   prop.Draw,
			Layer:  prop.Layer,
		})
	}

	for _, door := range gs.Doors {
		renderables = append(renderables, renderableObject{
			Object: &door.Object,
			Draw:   door.Draw,
			Layer:  door.Layer,
		})
	}

	if !gs.Player.Object.Destroyed {
		renderables = append(renderables, renderableObject{
			Object: &gs.Player.Object,
			Draw:   gs.Player.Draw,
			Layer:  gs.Player.Object.Layer,
		})
	}

	for _, enemy := range gs.EnemyManager.Enemies {
		if !enemy.Object.Destroyed {
			renderables = append(renderables, renderableObject{
				Object: &enemy.Object,
				Draw:   enemy.Draw,
				Layer:  enemy.Object.Layer,
			})
		}
	}

	if gs.Girlfriend.IsActive() {
		renderables = append(renderables, renderableObject{
			Object: &gs.Girlfriend.Object,
			Draw:   gs.Girlfriend.Draw,
			Layer:  gs.Girlfriend.Object.Layer,
		})
	}

	for _, item := range gs.Items {
		if item.IsDropped {
			renderables = append(renderables, renderableObject{
				Object: item.GetObject(),
				Draw:   func() { item.DrawAnimated(item.GetObject()) },
				Layer:  item.Object.Layer,
			})
		}
	}

	for _, weapon := range gs.Weapons {
		if weapon.IsDropped {
			renderables = append(renderables, renderableObject{
				Object: weapon.Object,
				Draw:   weapon.DrawAnimated,
				Layer:  weapon.Object.Layer,
			})
		}
	}

	return renderables
}

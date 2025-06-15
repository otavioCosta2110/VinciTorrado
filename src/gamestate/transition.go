package gamestate

import (
	"otaviocosta2110/vincitorrado/src/audio"
	"otaviocosta2110/vincitorrado/src/cutscene"
	"otaviocosta2110/vincitorrado/src/enemy"
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/girlfriend"
	"otaviocosta2110/vincitorrado/src/maps"
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/player"
	"otaviocosta2110/vincitorrado/src/props"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/system"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	playerScale int32 = 4
)

type TransitionHandler struct {
	Player        *player.Player
	EnemyManager  *enemy.EnemyManager
	Screen        *screen.Screen
	Kickables     *[]physics.Kickable
	Items         *[]*equipment.Equipment
	Props         *[]*props.Prop
	Doors         *[]*props.Door
	MapManager    *maps.MapManager
	Buildings     **rl.Texture2D
	Chao          *rl.Texture2D
	CurrentMap    *string
	Music         **string
	Cutscene      **cutscene.Cutscene
	Girfriend     *girlfriend.Girlfriend
	SkipCutscenes bool
}

func NewTransitionHandler(
	player *player.Player,
	enemyManager *enemy.EnemyManager,
	screen *screen.Screen,
	kickables *[]physics.Kickable,
	items *[]*equipment.Equipment,
	props *[]*props.Prop,
	doors *[]*props.Door,
	mapManager *maps.MapManager,
	buildings **rl.Texture2D,
	chao *rl.Texture2D,
	currentMap *string,
	music **string,
	cutscene **cutscene.Cutscene,
	gf *girlfriend.Girlfriend,
	skipCutscenes bool,
) *TransitionHandler {
	return &TransitionHandler{
		Player:        player,
		EnemyManager:  enemyManager,
		Screen:        screen,
		Kickables:     kickables,
		Items:         items,
		Props:         props,
		Doors:         doors,
		MapManager:    mapManager,
		Buildings:     buildings,
		Chao:          chao,
		CurrentMap:    currentMap,
		Music:         music,
		Cutscene:      cutscene,
		Girfriend:     gf,
		SkipCutscenes: skipCutscenes,
	}
}

func (th *TransitionHandler) TransitionMap(mapName string) {
	if (*th.Buildings).ID != 0 {
		rl.UnloadTexture(**th.Buildings)
	}
	if th.Chao.ID != 0 {
		rl.UnloadTexture(*th.Chao)
	}

	th.Player.Object.FrameX = 0
	th.Player.Object.FrameY = 0
	th.Player.IsKicking = false
	th.Player.LastKickTime = time.Now().Add(-time.Hour)
	th.Player.RecordInitialEquipment()

	newMap := th.MapManager.Maps[mapName]
	*th.CurrentMap = mapName

	**th.Buildings = system.LoadScaledTexture(newMap.Buildings, playerScale)
	*th.Chao = system.LoadScaledTexture(newMap.Floor, playerScale)

	enemies, err := enemy.LoadEnemiesFromJSON(newMap.EnemiesPath, playerScale)
	if err != nil {
		panic("Failed to load enemies: " + err.Error())
	}

	th.EnemyManager = &enemy.EnemyManager{
		CurrentMap: mapName,
	}
	for _, e := range enemies {
		e.Health = 0 
		th.EnemyManager.AddEnemy(e)
	}
	*th.Cutscene = cutscene.NewCutscene()

	switch *th.CurrentMap {
	case "city":
		music := "mission1"
		*th.Music = &music
		if !th.SkipCutscenes {
			(*th.Cutscene).IntroCutscenes(th.Player, th.Girfriend, th.EnemyManager)
			(*th.Cutscene).Start()
		}
	case "bar":
		music := "mission2"
		if !th.SkipCutscenes {
			(*th.Cutscene).BarIntroCutscene(th.Player, th.Girfriend, th.EnemyManager)
			(*th.Cutscene).Start()
		}
		*th.Music = &music
		audio.StopMusic()
		audio.PlayMission2Music()
	}

	props, doors, err := props.LoadPropsFromJSON(newMap.PropsPath, *th.Items)
	if err != nil {
		panic("Failed to load props: " + err.Error())
	}
	*th.Props = props
	*th.Doors = doors

	th.Player.Object.X = newMap.PlayerStartX
	th.Player.Object.Y = newMap.PlayerStartY
	th.Screen.ResetCamera()

	*th.Kickables = nil
	for _, prop := range *th.Props {
		*th.Kickables = append(*th.Kickables, prop)
	}
}

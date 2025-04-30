package enemy

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"otaviocosta2110/vincitorrado/src/sprites"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemyConfig struct {
	Sprite     string `json:"sprite"`
	X          int32  `json:"X"`
	Y          int32  `json:"Y"`
	Width      int32  `json:"width"`
	Height     int32  `json:"height"`
	Health     int32  `json:"health"`
	Damage     int32  `json:"damage"`
	Speed      int32  `json:"speed"`
	WindUpTime int64  `json:"windUpTime"`
	Scale      int32  `json:"scale"`
}

func LoadEnemiesFromJSON(filename string, playerScale int32, handleDrop func(x, y int32)) ([]*Enemy, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var configs []EnemyConfig
	err = json.Unmarshal(file, &configs)
	if err != nil {
		return nil, err
	}

	var enemies []*Enemy
	for _, config := range configs {
		sprite := sprites.Sprite{
			SpriteWidth:  config.Width,
			SpriteHeight: config.Height,
			Texture:      rl.LoadTexture(config.Sprite),
		}

		enemyType := "normal"
		if strings.Contains(strings.ToLower(filepath.Base(config.Sprite)), "dwarf") {
			enemyType = "dwarf"
			println("TA PEGANDO ANÃO:", config.Sprite)
		}

		enemy := NewEnemy(
			config.X,
			config.Y,
			config.Speed,
			config.Width,
			config.Height,
			playerScale,
			sprite,
			config.WindUpTime,
			enemyType,
			handleDrop,
		)
		enemy.Health = config.Health
		enemy.MaxHealth = config.Health
		enemy.Damage = config.Damage

		enemies = append(enemies, enemy)
	}

	return enemies, nil
}

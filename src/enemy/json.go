package enemy

import (
	"encoding/json"
	"fmt"
	"os"

	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/sprites"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type DropStats struct {
	Health int32 `json:"health"`
	Speed  int32 `json:"speed"`
	Damage int32 `json:"damage"`
}

type Drop struct {
	Name   string    `json:"name"`
	Sprite string    `json:"sprite"`
	Stats  DropStats `json:"stats"`
}

type Drops struct {
	Equipment []Drop `json:"equipment"`
}

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
	Drops      *Drop  `json:"drops"`
}

func LoadEnemiesFromJSON(filename string, playerScale int32) ([]*Enemy, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read enemy file: %w", err)
	}

	var configs []EnemyConfig
	if err := json.Unmarshal(file, &configs); err != nil {
		return nil, fmt.Errorf("failed to parse enemy JSON: %w", err)
	}

	var enemies []*Enemy
	for _, config := range configs {
		sprite := sprites.Sprite{
			SpriteWidth:  config.Width,
			SpriteHeight: config.Height,
			Texture:      rl.LoadTexture(config.Sprite),
		}

		var drop *equipment.Equipment
		if config.Drops != nil {
			dropSprite := config.Drops.Sprite
			dropName := config.Drops.Name
			dropStatsJson := config.Drops.Stats
			dropStats := equipment.Stats{Life: dropStatsJson.Health, Damage: dropStatsJson.Damage, Speed: dropStatsJson.Speed}
			drop = equipment.New(dropName, dropSprite, dropStats)
		} else {
			drop = nil
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
			"normal",
			drop,
		)

		enemy.Health = config.Health
		enemy.MaxHealth = config.Health
		enemy.Damage = config.Damage

		enemies = append(enemies, enemy)
	}

	return enemies, nil
}

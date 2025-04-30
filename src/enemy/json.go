package enemy

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"otaviocosta2110/vincitorrado/src/sprites"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Drop struct {
	Name   string `json:"name"`
	Sprite string `json:"sprite"`
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
	Drops      *Drops `json:"drops"`
}

func LoadEnemiesFromJSON(filename string, playerScale int32, handleDrop func(x, y int32, itemName string)) ([]*Enemy, error) {
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

		enemyType := "normal"
		if strings.Contains(strings.ToLower(filepath.Base(config.Sprite)), "dwarf") {
			enemyType = "dwarf"
			fmt.Println("ANÃO A FRENTE:", config.Sprite)
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
			func(x, y int32) {
				if config.Drops != nil {
					for _, drop := range config.Drops.Equipment {
						fmt.Printf("Attempting to drop %s at (%d,%d)\n", drop.Name, x, y)
						handleDrop(x, y, drop.Name)
					}
				}
			},
		)
		enemy.Health = config.Health
		enemy.MaxHealth = config.Health
		enemy.Damage = config.Damage

		enemies = append(enemies, enemy)
	}

	return enemies, nil
}

package enemy

import (
	"encoding/json"
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
	"otaviocosta2110/vincitorrado/src/sprites"
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
}

func LoadEnemiesFromJSON(filename string, playerScale int32) ([]*Enemy, error) {
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

		enemy := NewEnemy(
			config.X,
			config.Y,
			config.Speed,
			config.Width,
			config.Height,
			playerScale,
			sprite,
			config.WindUpTime,
		)
		enemy.Health = config.Health
		enemy.Damage = config.Damage

		enemies = append(enemies, enemy)
	}

	return enemies, nil
}

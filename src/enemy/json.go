package enemy

import (
	"encoding/json"
	"fmt"
	"os"

	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/objects"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"
	"otaviocosta2110/vincitorrado/src/weapon"

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

type WeaponDrop struct {
	Sprite  string    `json:"sprite"`
	HitboxX int32     `json:"hitbox_X"`
	HitboxY int32     `json:"hitbox_Y"`
	OffsetX int32     `json:"offset_X"`
	OffsetY int32     `json:"offset_Y"`
	Width   int32     `json:"width"`
	Height  int32     `json:"height"`
	Stats   DropStats `json:"stats"`
	Health  int32     `json:"health"`
	Scale   int32     `json:"scale"`
}

type EnemyConfig struct {
	Sprite         string      `json:"sprite"`
	X              int32       `json:"X"`
	Y              int32       `json:"Y"`
	Activate_pos_X *int32      `json:"activate_pos_X"`
	Activate_pos_Y *int32      `json:"activate_pos_Y"`
	Width          int32       `json:"width"`
	Height         int32       `json:"height"`
	Health         int32       `json:"health"`
	Damage         int32       `json:"damage"`
	Speed          int32       `json:"speed"`
	WindUpTime     int64       `json:"windUpTime"`
	Scale          int32       `json:"scale"`
	Drops          *Drop       `json:"drops"`
	Weapon         *WeaponDrop `json:"weapon"`
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
			dropStats := objects.Stats{Life: dropStatsJson.Health, Damage: dropStatsJson.Damage, Speed: dropStatsJson.Speed}
			drop = equipment.New(dropName, dropSprite, dropStats)
		} else {
			drop = nil
		}

		var weaponFromJson *weapon.Weapon
		if config.Weapon != nil {
			// Create a weapon.Weapon from the WeaponDrop data.
			weaponSprite := sprites.Sprite{
				SpriteWidth:  config.Weapon.Width,
				SpriteHeight: config.Weapon.Height,
				Texture:      rl.LoadTexture(config.Weapon.Sprite),
			}
			weaponStatsJson := config.Weapon.Stats
			weaponStats := objects.Stats{Life: weaponStatsJson.Health, Damage: weaponStatsJson.Damage, Speed: weaponStatsJson.Speed}
			weaponFromJson = weapon.New(
				&system.Object{
					Sprite: weaponSprite,
					Scale: config.Scale,
					Width: config.Weapon.Width * config.Weapon.Scale,
					Height: config.Weapon.Height * config.Weapon.Scale,
				},
				config.Weapon.OffsetX,
				config.Weapon.OffsetY,
				config.Weapon.HitboxX,
				config.Weapon.HitboxY,
				weaponStats,
				config.Weapon.Health,
				true,
				false,
			)
		}

		if config.Activate_pos_X == nil {
			config.Activate_pos_X = &config.X
		}
		if config.Activate_pos_Y == nil {
			config.Activate_pos_Y = &config.Y
		}

		enemy := NewEnemy(
			config.X,
			config.Y,
			*config.Activate_pos_X,
			*config.Activate_pos_Y,
			config.Speed,
			config.Width,
			config.Height,
			playerScale,
			sprite,
			config.WindUpTime,
			"normal",
			drop,
			weaponFromJson,
		)

		enemy.Health = config.Health
		enemy.MaxHealth = config.Health
		enemy.Damage = config.Damage

		enemies = append(enemies, enemy)
	}

	return enemies, nil
}

package weapon

import (
	"encoding/json"
	"fmt"
	"os"

	"otaviocosta2110/vincitorrado/src/objects"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type WeaponConfig struct {
	Sprite  string        `json:"sprite"`
	X       int32         `json:"X"`
	Y       int32         `json:"Y"`
	HitboxX int32         `json:"hitbox_X"`
	HitboxY int32         `json:"hitbox_Y"`
	OffsetX int32         `json:"offset_X"`
	OffsetY int32         `json:"offset_Y"`
	Width   int32         `json:"width"`
	Height  int32         `json:"height"`
	Stats   objects.Stats `json:"stats"`
	Health  int32         `json:"health"`
	Scale   int32         `json:"scale"`
	IsGun   bool          `json:"is_gun"`
	Ammo    int32         `json:"ammo"`
	MaxAmmo int32         `json:"maxAmmo"`
}

func LoadWeaponsFromJSON(filename string) ([]*Weapon, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read weapon file: %w", err)
	}

	var configs []WeaponConfig
	if err := json.Unmarshal(file, &configs); err != nil {
		return nil, fmt.Errorf("failed to parse weapon JSON: %w", err)
	}

	var weapons []*Weapon
	for _, config := range configs {
		stats := objects.Stats{
			Damage: config.Stats.Damage,
		}

		spritesheet := sprites.Sprite{
			SpriteWidth:  config.Width,
			SpriteHeight: config.Height,
			Texture:      rl.LoadTexture(config.Sprite),
		}

		weapon := &Weapon{
			Object: &system.Object{
				X:      config.X,
				Y:      config.Y,
				Width:  config.Width * config.Scale,
				Height: config.Height * config.Scale,
				Scale:  config.Scale,
				Sprite: spritesheet,
			},
			IsDropped: true,
			Stats:     stats,
			HitboxX:   config.HitboxX,
			HitboxY:   config.HitboxY,
			OffsetX:   config.OffsetX,
			OffsetY:   config.OffsetY,
			Health:    config.Health,
			IsGun:     config.IsGun,
			Ammo:      config.Ammo,
			MaxAmmo:   config.MaxAmmo,
		}

		weapons = append(weapons, weapon)
	}

	return weapons, nil
}

package items

import (
	"encoding/json"
	"fmt"
	"os"

	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/objects"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ItemConfig struct {
	Name   string          `json:"name"`
	Sprite string          `json:"sprite"`
	X      int32           `json:"X"`
	Y      int32           `json:"Y"`
	Width  int32           `json:"width"`
	Height int32           `json:"height"`
	Scale  int32           `json:"scale"`
	Type   string          `json:"type"`
	Stats  objects.Stats `json:"stats"`
}

func LoadItemsFromJSON(filename string, playerScale int32) ([]*equipment.Equipment, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read item file: %w", err)
	}

	var configs []ItemConfig
	if err := json.Unmarshal(file, &configs); err != nil {
		return nil, fmt.Errorf("failed to parse item JSON: %w", err)
	}

	var items []*equipment.Equipment
	for _, config := range configs {
		item := equipment.New(
			config.Name,
			config.Sprite,
			config.Stats,
		)

		item.IsDropped = true
		item.Type = config.Type
		item.Object.X = config.X
		item.Object.Y = config.Y
		item.Object.Width = config.Width * config.Scale
		item.Object.Height = config.Height * config.Scale
		item.Object.Scale = config.Scale
		item.Object.Sprite.Texture = rl.LoadTexture(config.Sprite)

		items = append(items, item)
	}

	return items, nil
}

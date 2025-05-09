package equipment

import (
	"encoding/json"
	"os"
	"math"
	"otaviocosta2110/vincitorrado/src/objects"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Equipment struct {
	Name       string
	IsEquipped bool
	IsDropped  bool
	Type       string
	Stats      objects.Stats
	OffsetX    int32
	OffsetY    int32
	Object     system.Object
}

func New(name string, texturePath string, stats objects.Stats) *Equipment {
	spritesheet := sprites.Sprite{
		SpriteWidth:  32,
		SpriteHeight: 32,
		Texture:      rl.LoadTexture(texturePath),
	}
	return &Equipment{
		Name:    name,
		Stats:   stats,
		OffsetX: 0,
		OffsetY: 0,
		Object: system.Object{
			Sprite: spritesheet,
		},
	}
}

func (e *Equipment) DrawAnimated(obj *system.Object) {
	if !e.IsDropped {
		return
	}

	frameWidth := float32(e.Object.Sprite.SpriteWidth)
	frameHeight := float32(e.Object.Sprite.SpriteHeight)

	currentTime := float32(rl.GetTime())
	floatOffset := float32(math.Sin(float64(currentTime*2)) * 5)

	source := rl.NewRectangle(
		float32(obj.FrameX)*frameWidth,
		float32(obj.FrameY)*frameHeight,
		frameWidth,
		frameHeight,
	)

	dest := rl.NewRectangle(
		float32(obj.X)+float32(e.OffsetX),
		float32(obj.Y)+float32(e.OffsetY)+floatOffset-20,
		frameWidth*float32(obj.Scale),
		frameHeight*float32(obj.Scale),
	)

	rl.DrawTexturePro(
		e.Object.Sprite.Texture,
		source,
		dest,
		rl.NewVector2(dest.Width/2, dest.Height/2),
		0,
		rl.White,
	)
}

func NewConsumable(name, spritePath string, stats objects.Stats) *Equipment {
	return &Equipment{
		Name:  name,
		Type:  "consumable",
		Stats: stats,
		Object: system.Object{
			Width:  32,
			Height: 32,
			Scale:  4,
			Sprite: sprites.Sprite{
				Texture:      rl.LoadTexture(spritePath),
				SpriteWidth:  32,
				SpriteHeight: 32,
			},
		},
	}
}

func LoadItemsFromJSON(path string) ([]*Equipment, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var jsonItems []struct {
		Name   string `json:"name"`
		Sprite string `json:"sprite"`
		Type   string `json:"type"`
		Stats  objects.Stats  `json:"stats"`
		Scale  int32  `json:"scale"`
	}

	if err := json.Unmarshal(data, &jsonItems); err != nil {
		return nil, err
	}

	var items []*Equipment
	for _, item := range jsonItems {
		equip := NewConsumable(
			item.Name,
			item.Sprite,
			item.Stats,
		)
		equip.Object.Scale = item.Scale
		items = append(items, equip)
	}

	return items, nil
}
func (e *Equipment) GetObject() *system.Object {
	return &e.Object
}

package props

import (
	"encoding/json"
	"math/rand"
	"os"
	"otaviocosta2110/vincitorrado/src/audio"
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PropType string

const (
	PropTypeTrash   PropType = "trash"
	PropTypeDoor    PropType = "door"
	PropTypeTable   PropType = "table"
	PropTypeJukebox PropType = "jukebox"
)

type PropConfig struct {
	X             int32    `json:"X"`
	Y             int32    `json:"Y"`
	Width         int32    `json:"width"`
	Height        int32    `json:"height"`
	Scale         int32    `json:"scale"`
	NormalTexture string   `json:"normal_texture"`
	KickedTexture string   `json:"kicked_texture"`
	Loot          []string `json:"loot"`
	Type          PropType `json:"type"`
	NextMap       string   `json:"next_map,omitempty"`
}

type Prop struct {
	system.Object
	LootTable      []*equipment.Equipment
	Kicked         bool
	NormalTexture  rl.Texture2D
	KickedTexture  rl.Texture2D
	HitboxOffset   float32
	Type           PropType
	OriginalWidth  int32
	OriginalHeight int32
	OriginalY      int32
}

func LoadPropsFromJSON(path string, items []*equipment.Equipment) ([]*Prop, []*Door, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	var configs []PropConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		return nil, nil, err
	}

	var props []*Prop
	var doors []*Door

	for _, cfg := range configs {
		if cfg.Type == PropTypeDoor {
			door := NewDoor(
				cfg.X,
				cfg.Y,
				cfg.Width,
				cfg.Height,
				cfg.Scale,
				cfg.NormalTexture,
				cfg.NextMap,
			)
			doors = append(doors, door)
		} else {
			var loot []*equipment.Equipment
			for _, itemName := range cfg.Loot {
				for _, item := range items {
					if item.Name == itemName {
						loot = append(loot, item)
						break
					}
				}
			}

			prop := NewProp(
				cfg.X,
				cfg.Y,
				cfg.Scale,
				cfg.Width,
				cfg.Height,
				cfg.NormalTexture,
				cfg.KickedTexture,
				loot,
			)
			prop.OriginalWidth = cfg.Width
			prop.OriginalHeight = cfg.Height
			props = append(props, prop)
			prop.Type = cfg.Type
		}
	}
	return props, doors, nil
}

func NewProp(x, y, scale, width, height int32, normalTexPath, kickedTexPath string, loot []*equipment.Equipment) *Prop {
	normalTex := rl.LoadTexture(normalTexPath)
	kickedTex := rl.LoadTexture(kickedTexPath)

	return &Prop{
		Object: system.Object{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
			Scale:  scale,
			Sprite: sprites.Sprite{
				Texture:      normalTex,
				SpriteWidth:  width / scale,
				SpriteHeight: height / scale,
			},
		},
		LootTable:     loot,
		NormalTexture: normalTex,
		KickedTexture: kickedTex,
		OriginalY:     y,
	}
}

func (t *Prop) Draw() {
	rl.DrawTexturePro(
		t.Sprite.Texture,
		rl.NewRectangle(0, 0, 32, 32),
		rl.NewRectangle(
			float32(t.X),
			float32(t.Y),
			float32(t.Width)*float32(t.Scale),
			float32(t.Height)*float32(t.Scale),
		),
		rl.Vector2{},
		0,
		rl.White,
	)
}

func (t *Prop) IsKicked() bool {
	return t.Kicked
}

func (t *Prop) HandleKick(items *[]*equipment.Equipment, kicker system.Object) {
	t.Kicked = true
	t.Object.Sprite.Texture = t.KickedTexture

	switch t.Type {
	case PropTypeTable:
		return
	case PropTypeJukebox:
		audio.StopMusic()
		return
	}

	proto := t.LootTable[rand.Intn(len(t.LootTable))]
	item := &equipment.Equipment{
		Name:      proto.Name,
		Type:      proto.Type,
		Stats:     proto.Stats,
		IsDropped: true,
		Object: system.Object{
			X:      t.Object.X + (t.Object.Width*t.Object.Scale)/2,
			Y:      t.Object.Y,
			Width:  proto.Object.Width / 2,
			Height: proto.Object.Height,
			Scale:  proto.Object.Scale,
			Sprite: sprites.Sprite{
				Texture:      proto.Object.Sprite.Texture,
				SpriteWidth:  proto.Object.Sprite.SpriteWidth,
				SpriteHeight: proto.Object.Sprite.SpriteHeight,
			},
		},
	}
	*items = append(*items, item)
}

func (t *Prop) GetObject() system.Object {
	obj := t.Object

	if t.Kicked && t.Type == PropTypeTable {
		obj.Height = t.OriginalHeight + (t.OriginalHeight / 2)

		obj.Y = t.OriginalY + (t.OriginalHeight * t.Scale / 4)
	}

	return obj
}

func (t *Prop) Reset() {
	t.Kicked = false
	t.Object.Sprite.Texture = t.NormalTexture
	t.Object.Y = t.OriginalY
	t.Object.Height = t.OriginalHeight
	t.Object.Width = t.OriginalWidth
}

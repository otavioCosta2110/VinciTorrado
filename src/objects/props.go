package objects

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TrashConfig struct {
	X             int32    `json:"X"`
	Y             int32    `json:"Y"`
	Width         int32    `json:"width"`
	Height        int32    `json:"height"`
	Scale         int32    `json:"scale"`
	NormalTexture string   `json:"normal_texture"`
	KickedTexture string   `json:"kicked_texture"`
	Loot          []string `json:"loot"`
}

type TrashCan struct {
	system.Object
	LootTable     []*equipment.Equipment
	Kicked        bool
	NormalTexture rl.Texture2D
	KickedTexture rl.Texture2D
	HitboxOffset  float32
}

func LoadTrashCansFromJSON(path string, items []*equipment.Equipment) ([]*TrashCan, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var configs []TrashConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		return nil, err
	}

	var trashCans []*TrashCan
	for _, cfg := range configs {
		var loot []*equipment.Equipment
		for _, itemName := range cfg.Loot {
			for _, item := range items {
				if item.Name == itemName {
					loot = append(loot, item)
					break
				}
			}
		}

		trash := NewTrashCan(
			cfg.X,
			cfg.Y,
			cfg.Scale,
			cfg.Width,
			cfg.Height,
			cfg.NormalTexture,
			cfg.KickedTexture,
			loot,
		)
		trashCans = append(trashCans, trash)
	}

	return trashCans, nil
}

func NewTrashCan(x, y, scale, width, height int32, normalTexPath, kickedTexPath string, loot []*equipment.Equipment) *TrashCan {
	normalTex := rl.LoadTexture(normalTexPath)
	kickedTex := rl.LoadTexture(kickedTexPath)

	return &TrashCan{
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
	}
}

func (t *TrashCan) Draw() {
	rl.DrawTexturePro(
		t.Sprite.Texture,
		rl.NewRectangle(0, 0, 32, 32),
		rl.NewRectangle(
			float32(t.X),
			float32(t.Y),
			float32(t.Width),
			float32(t.Height),
		),
		rl.Vector2{},
		0,
		rl.White,
	)
}

func (t *TrashCan) HandleKick(kickHitbox system.Object, items *[]*equipment.Equipment, isFlipped bool, kickPower int32) bool {
	if !t.Kicked && physics.CheckCollision(kickHitbox, t.Object) {
		t.Kicked = true
		t.Object.Sprite.Texture = t.KickedTexture

		proto := t.LootTable[rand.Intn(len(t.LootTable))]
		item := &equipment.Equipment{
			Name:      proto.Name,
			Type:      proto.Type,
			Stats:     proto.Stats,
			IsDropped: true,
			Object: system.Object{
				X:      t.Object.X,
				Y:      t.Object.Y,
				Width:  proto.Object.Width,
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
		return true
	}
	return false
}

func (t *TrashCan) GetObject() system.Object {
	return t.Object
}

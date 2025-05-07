package objects

import (
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TrashCan struct {
	system.Object
	LootTable     []*equipment.Equipment
	Kicked        bool
	NormalTexture rl.Texture2D
	KickedTexture rl.Texture2D
	HitboxOffset  float32
}

func NewTrashCan(x, y, scale int32, loot []*equipment.Equipment) *TrashCan {
	normalTex := rl.LoadTexture("assets/props/Lixo.png")
	kickedTex := rl.LoadTexture("assets/props/Lixeira_Tombada.png")

	return &TrashCan{
		Object: system.Object{
			X:      x,
			Y:      y,
			Width:  32 * scale,
			Height: 32 * scale,
			Scale:  scale,
			Sprite: sprites.Sprite{
				Texture:      normalTex,
				SpriteWidth:  32,
				SpriteHeight: 32,
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

func (t *TrashCan) Unload() {
	rl.UnloadTexture(t.NormalTexture)
	rl.UnloadTexture(t.KickedTexture)
}

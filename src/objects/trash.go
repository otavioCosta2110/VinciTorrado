package objects

import (
	"math/rand"
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TrashCan struct {
	system.Object
	LootTable []*equipment.Equipment
	Kicked    bool
}

func NewTrashCan(x, y, scale int32, loot []*equipment.Equipment) *TrashCan {
	return &TrashCan{
		Object: system.Object{
			X:      x,
			Y:      y,
			Width:  32 * scale,
			Height: 32 * scale,
			Scale:  scale,
			Sprite: sprites.Sprite{
				Texture:      rl.LoadTexture("assets/props/Lixo.png"),
				SpriteWidth:  32,
				SpriteHeight: 32,
			},
		},
		LootTable: loot,
	}
}

func (t *TrashCan) Update(player system.Object, items *[]*equipment.Equipment) {
	if !t.Kicked && physics.CheckCollision(player, t.Object) && player.IsKicking {
		t.Kicked = true
		t.spawnLoot(items)
	}
}

func (t *TrashCan) spawnLoot(items *[]*equipment.Equipment) {
	if len(t.LootTable) > 0 {
		item := *t.LootTable[rand.Intn(len(t.LootTable))]
		item.Object.X = t.X
		item.Object.Y = t.Y
		item.IsDropped = true
		*items = append(*items, &item)
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

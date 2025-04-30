package objects

import (
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EquipmentPickup struct {
	Object    system.Object
	Equipment *equipment.Equipment
}

func NewTurbantePickup(x, y int32) *EquipmentPickup {
	texture := rl.LoadTexture("assets/player/Turbante.png")
	return &EquipmentPickup{
		Object: system.Object{
			X:      x,
			Y:      y,
			Width:  32,
			Height: 32,
		},
		Equipment: equipment.New("assets/player/Turbante.png", sprites.Sprite{
			SpriteWidth:  32,
			SpriteHeight: 32,
			Texture:      texture,
		}),
	}
}

func (ep *EquipmentPickup) Draw() {
	rl.DrawTexturePro(
		ep.Equipment.Texture,
		rl.NewRectangle(0, 0, float32(ep.Equipment.Texture.Width), float32(ep.Equipment.Texture.Height)),
		rl.NewRectangle(
			float32(ep.Object.X),
			float32(ep.Object.Y),
			float32(ep.Object.Width),
			float32(ep.Object.Height),
		),
		rl.NewVector2(float32(ep.Object.Width)/2, float32(ep.Object.Height)/2),
		0,
		rl.White,
	)
}

package props

import (
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Door struct {
	system.Object
	NextMap  string
	IsActive bool
	Texture  rl.Texture2D
}

func NewDoor(x, y, width, height, scale int32, texturePath, nextMap string) *Door {
	return &Door{
		Object: system.Object{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
			Scale:  scale,
			Sprite: sprites.Sprite{
				Texture:      rl.LoadTexture(texturePath),
				SpriteWidth:  width,
				SpriteHeight: height,
			},
		},
		NextMap: nextMap,
		Texture: rl.LoadTexture(texturePath),
	}
}

func (d *Door) CheckTransition(player system.Object, enemiesCleared bool) bool {
	return enemiesCleared && physics.CheckCollision(player, d.Object)
}

func (d *Door) Draw() {
	rl.DrawTexturePro(
		d.Texture,
		rl.NewRectangle(0, 0, float32(d.Object.Width), float32(d.Object.Height)),
		rl.NewRectangle(
			float32(d.Object.X),
			float32(d.Object.Y),
			float32(d.Object.Width)*float32(d.Object.Scale),
			float32(d.Object.Height)*float32(d.Object.Scale),
		),
		rl.Vector2{},
		0,
		rl.White,
	)
}

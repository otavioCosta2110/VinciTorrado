package girlfriend

import (
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Girlfriend struct {
	Object system.Object
	active bool
}

func (g *Girlfriend) IsActive() bool {
	return g.active
}

func (g *Girlfriend) SetActive(active bool) {
	g.active = active
}

func New(sprite sprites.Sprite, x, y int32, scale int32) *Girlfriend {
	return &Girlfriend{
		Object: system.Object{
			Sprite: sprite,
			X:      x,
			Y:      y,
			Scale:  scale,
		},
		active: true,
	}
}

func (g *Girlfriend) Update() {

}

func (g *Girlfriend) Draw() {
	var gfWidth float32 = float32(g.Object.Sprite.SpriteWidth)
	if g.Object.Flipped {
		gfWidth = -gfWidth
	}

	sourceRec := rl.NewRectangle(
		float32(g.Object.FrameX)*float32(g.Object.Sprite.SpriteWidth),
		float32(g.Object.FrameY)*float32(g.Object.Sprite.SpriteHeight),
		gfWidth,
		float32(g.Object.Sprite.SpriteHeight),
	)

	destinationRec := rl.NewRectangle(
		float32(g.Object.X),
		float32(g.Object.Y),
		float32(g.Object.Sprite.SpriteWidth)*float32(g.Object.Scale),
		float32(g.Object.Sprite.SpriteHeight)*float32(g.Object.Scale),
	)

	origin := rl.NewVector2(
		destinationRec.Width/2,
		destinationRec.Height/2,
	)

	rl.DrawTexturePro(
		g.Object.Sprite.Texture,
		sourceRec,
		destinationRec,
		origin,
		0.0,
		rl.White,
	)
}

func (g *Girlfriend) UpdateAnimation(animationName string) {
	switch animationName {
	case "walk":
		g.runAnimation(300, []int{0, 1}, []int{0, 0})
	}
}

func (g *Girlfriend) runAnimation(animationDelay int, framesX, framesY []int) {
	g.Object.UpdateAnimation(animationDelay, framesX, framesY)
}

func (g *Girlfriend) SetObject(obj system.Object) {
	g.Object = obj
}
func (g *Girlfriend) GetObject() system.Object {
	return g.Object
}
func (g *Girlfriend) TakeDamage(int32, system.Object) {}

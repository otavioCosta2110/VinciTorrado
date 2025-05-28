package weapon

import (
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Projectile struct {
	Object    *system.Object
	Speed     float32
	Direction rl.Vector2
	IsActive  bool
	Damage    int32
	Lifetime  float32
}

func (w *Weapon) Shoot(startX, startY float32, direction rl.Vector2) *Projectile {
	if w.Ammo <= 0 {
		return nil
	}
	if w.HasActiveBullets(){
		return nil
	}

	w.Ammo--

	proj := &Projectile{
		Object: &system.Object{
			X:      int32(startX),
			Y:      int32(startY),
			Width:  8 * w.Object.Scale,
			Height: 8 * w.Object.Scale,
			Scale:  w.Object.Scale,
			Sprite: sprites.Sprite{
				SpriteWidth:  8,
				SpriteHeight: 8,
				Texture:      rl.LoadTexture("assets/weapons/bullet.png"),
			},
		},
		Speed:     10.0,
		Direction: direction,
		IsActive:  true,
		Damage:    w.Stats.Damage,
		Lifetime:  2.0,
	}
	println("Projectile created at position:", proj.Damage)
	w.Projectile = proj

	return proj
}

func (p *Projectile) Update() {
	if !p.IsActive {
		return
	}

	p.Object.X += int32(p.Direction.X * p.Speed)
	p.Object.Y += int32(p.Direction.Y * p.Speed)

	p.Lifetime -= rl.GetFrameTime()
	if p.Lifetime <= 0 {
		p.IsActive = false
	}
}

func (p *Projectile) Draw() {
	if !p.IsActive {
		return
	}

	source := rl.NewRectangle(
		0, 0,
		float32(p.Object.Sprite.SpriteWidth),
		float32(p.Object.Sprite.SpriteHeight),
	)

	dest := rl.NewRectangle(
		float32(p.Object.X),
		float32(p.Object.Y),
		float32(p.Object.Sprite.SpriteWidth)*float32(p.Object.Scale),
		float32(p.Object.Sprite.SpriteHeight)*float32(p.Object.Scale),
	)

	origin := rl.NewVector2(dest.Width/2, dest.Height/2)

	rl.DrawTexturePro(
		p.Object.Sprite.Texture,
		source,
		dest,
		origin,
		0.0,
		rl.White,
	)
}

package weapon

import (
	"math"
	"otaviocosta2110/vincitorrado/src/objects"
	"otaviocosta2110/vincitorrado/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Weapon struct {
	Object     *system.Object
	IsDropped  bool
	IsEquipped bool
	OffsetX    int32
	OffsetY    int32
	HitboxX    int32
	HitboxY    int32
	Stats      objects.Stats
	Health     int32
	IsGun      bool
	Ammo       int32       
	MaxAmmo    int32       
	Projectile *Projectile 
}

func New(obj *system.Object, offsetX, offsetY int32, hitboxX, hitboxY int32, stats objects.Stats, health int32, isEquipped bool, isDropped bool) *Weapon {
	return &Weapon{
		Object:     obj,
		IsDropped:  isDropped,
		IsEquipped: isEquipped,
		OffsetX:    offsetX,
		OffsetY:    offsetY,
		HitboxX:    hitboxX,
		HitboxY:    hitboxY,
		Stats:      stats,
		Health:     health,
	}
}

func (w *Weapon) DrawAnimated() {
	frameWidth := float32(w.Object.Sprite.SpriteWidth)
	frameHeight := float32(w.Object.Sprite.SpriteHeight)

	currentTime := float32(rl.GetTime())
	floatOffset := float32(math.Sin(float64(currentTime*2)) * 5)

	source := rl.NewRectangle(
		float32(w.Object.FrameX)*frameWidth,
		float32(w.Object.FrameY)*frameHeight,
		frameWidth,
		frameHeight,
	)

	dest := rl.NewRectangle(
		float32(w.Object.X)+float32(w.OffsetX),
		float32(w.Object.Y)+float32(w.OffsetY)+floatOffset-20,
		frameWidth*float32(w.Object.Scale),
		frameHeight*float32(w.Object.Scale),
	)

	origin := rl.NewVector2(dest.Width/2, dest.Height/2)

	outlineSize := float32(4.0)

	directions := []rl.Vector2{
		{X: -outlineSize, Y: -outlineSize},
		{X: 0, Y: -outlineSize},
		{X: outlineSize, Y: -outlineSize},
		{X: -outlineSize, Y: 0},
		{X: outlineSize, Y: 0},
		{X: -outlineSize, Y: outlineSize},
		{X: 0, Y: outlineSize},
		{X: outlineSize, Y: outlineSize},
	}

	for _, dir := range directions {
		outlineDest := rl.NewRectangle(
			dest.X+dir.X,
			dest.Y+dir.Y,
			dest.Width,
			dest.Height,
		)
		rl.DrawTexturePro(
			w.Object.Sprite.Texture,
			source,
			outlineDest,
			origin,
			0,
			rl.Blue,
		)
	}

	rl.DrawTexturePro(
		w.Object.Sprite.Texture,
		source,
		dest,
		origin,
		0,
		rl.White,
	)
}
func (w *Weapon) DrawEquipped(obj *system.Object) {
	frameWidth := float32(w.Object.Sprite.SpriteWidth)
	if obj.Flipped {
		frameWidth = -frameWidth
	}
	frameHeight := float32(w.Object.Sprite.SpriteHeight)

	weaponSource := rl.NewRectangle(
		float32(w.Object.FrameX)*float32(w.Object.Sprite.SpriteWidth),
		float32(w.Object.FrameY)*frameHeight,
		frameWidth,
		frameHeight,
	)

	weaponOffsetX := w.OffsetX
	if obj.Flipped {
		weaponOffsetX = (int32(w.Object.Sprite.SpriteWidth)) - ((w.Object.Width / w.Object.Scale) + weaponOffsetX)
	}
	weaponDestination := rl.Rectangle{
		X:      float32(obj.X) + float32(weaponOffsetX)*float32(obj.Scale),
		Y:      float32(obj.Y),
		Width:  float32(w.Object.Sprite.SpriteWidth) * float32(obj.Scale),
		Height: float32(w.Object.Sprite.SpriteHeight) * float32(obj.Scale),
	}
	weaponOrigin := rl.NewVector2(
		weaponDestination.Width/2,
		weaponDestination.Height/2,
	)
	rl.DrawTexturePro(
		w.Object.Sprite.Texture,
		weaponSource,
		weaponDestination,
		weaponOrigin,
		0.0,
		rl.White,
	)
}
func (w *Weapon) Clone() *Weapon {
	return &Weapon{
		Object:     w.Object,
		IsDropped:  w.IsDropped,
		IsEquipped: w.IsEquipped,
		OffsetX:    w.OffsetX,
		OffsetY:    w.OffsetY,
		HitboxX:    w.HitboxX,
		HitboxY:    w.HitboxY,
		Stats:      w.Stats,
		Health:     w.Health,
	}
}
func (w *Weapon) GetDropCollisionBox() system.Object {
	dropWidth := int32(32 * w.Object.Scale)
	dropHeight := int32(32 * w.Object.Scale)
	dropY := w.Object.Y - 20
	return system.Object{
		X:      w.Object.X - dropWidth/4,
		Y:      dropY,
		Width:  dropWidth / 4,
		Height: dropHeight / 4,
	}
}
func (w *Weapon) HasActiveBullets() bool {
	if w.Projectile == nil {
		return false
	}

	return w.Projectile.IsActive
}

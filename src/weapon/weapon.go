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
}

func New( obj *system.Object, offsetX, offsetY int32, hitboxX, hitboxY int32, stats objects.Stats, health int32, isEquipped bool, isDropped bool) *Weapon {
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
	if !w.IsDropped {
		return
	}

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

	rl.DrawTexturePro(
		w.Object.Sprite.Texture,
		source,
		dest,
		rl.NewVector2(dest.Width/2, dest.Height/2),
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
	// println(int32(w.Object.Sprite.SpriteWidth) - ((w.Object.Width/w.Object.Scale) + weaponOffsetX))
	println(w.IsEquipped)
	if obj.Flipped {
		weaponOffsetX = (int32(w.Object.Sprite.SpriteWidth)) - ((w.Object.Width/w.Object.Scale) + weaponOffsetX)
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

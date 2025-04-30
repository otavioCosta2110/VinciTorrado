package equipment

import (
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Equipment struct {
	Texture     rl.Texture2D
	IsEquipped  bool
	OffsetX     int32
	OffsetY     int32
	SpriteSheet sprites.Sprite
}

func New(texturePath string, spriteSheet sprites.Sprite) *Equipment {
	return &Equipment{
		Texture:     rl.LoadTexture(texturePath),
		SpriteSheet: spriteSheet,
		OffsetX:     0,
		OffsetY:     -10,
	}
}

func (e *Equipment) DrawAnimated(obj *system.Object) {
	frameWidth := float32(e.SpriteSheet.SpriteWidth)
	frameHeight := float32(e.SpriteSheet.SpriteHeight)

	source := rl.NewRectangle(
		float32(obj.FrameX)*frameWidth,
		float32(obj.FrameY)*frameHeight,
		frameWidth,
		frameHeight,
	)

	dest := rl.NewRectangle(
		float32(obj.X)+float32(e.OffsetX),
		float32(obj.Y)+float32(e.OffsetY),
		frameWidth*float32(obj.Scale),
		frameHeight*float32(obj.Scale),
	)

	rl.DrawTexturePro(
		e.SpriteSheet.Texture,
		source,
		dest,
		rl.NewVector2(dest.Width/2, dest.Height/2),
		0,
		rl.White,
	)
}

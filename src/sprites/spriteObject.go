package sprites

import rl "github.com/gen2brain/raylib-go/raylib"

type Sprite struct {
	SpriteWidth  int32
	SpriteHeight int32
	Texture      rl.Texture2D
}

func (sprite Sprite) GetSpriteByCoordinates(x, y, width, height int32) rl.Rectangle {
	return rl.NewRectangle(
		float32(x)*float32(sprite.SpriteWidth),
		float32(y)*float32(sprite.SpriteHeight),
		float32(width),
		float32(height),
	)
}

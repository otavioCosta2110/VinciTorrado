package equipment

import (
	"math"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Equipment struct {
	IsEquipped bool
	OffsetX    int32
	OffsetY    int32
	Object     system.Object
}

func New(texturePath string) *Equipment {
	spritesheet := sprites.Sprite{
		SpriteWidth:  32,
		SpriteHeight: 32,
		Texture:      rl.LoadTexture(texturePath),
	}
	return &Equipment{
		OffsetX: 0,
		OffsetY: 0,
		Object: system.Object{
			Sprite: spritesheet,
		},
	}
}

func (e *Equipment) DrawAnimated(obj *system.Object) {
    frameWidth := float32(e.Object.Sprite.SpriteWidth)
    frameHeight := float32(e.Object.Sprite.SpriteHeight)

    currentTime := float32(rl.GetTime())
    floatOffset := float32(math.Sin(float64(currentTime*2)) * 5 )

    source := rl.NewRectangle(
        float32(obj.FrameX)*frameWidth,
        float32(obj.FrameY)*frameHeight,
        frameWidth,
        frameHeight,
    )

    dest := rl.NewRectangle(
        float32(obj.X)+float32(e.OffsetX),
        float32(obj.Y)+float32(e.OffsetY)+floatOffset - 20, 
        frameWidth*float32(obj.Scale),
        frameHeight*float32(obj.Scale),
    )

    rl.DrawTexturePro(
        e.Object.Sprite.Texture,
        source,
        dest,
        rl.NewVector2(dest.Width/2, dest.Height/2),
        0,
        rl.White,
    )
}

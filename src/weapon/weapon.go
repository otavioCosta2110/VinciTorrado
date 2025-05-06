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
}

func (w *Weapon) DrawAnimated(obj *system.Object) {
    if !w.IsDropped {
        return
    }

    frameWidth := float32(w.Object.Sprite.SpriteWidth)
    frameHeight := float32(w.Object.Sprite.SpriteHeight)

    currentTime := float32(rl.GetTime())
    floatOffset := float32(math.Sin(float64(currentTime*2)) * 5)

    source := rl.NewRectangle(
        float32(obj.FrameX)*frameWidth,
        float32(obj.FrameY)*frameHeight,
        frameWidth,
        frameHeight,
    )

    dest := rl.NewRectangle(
        float32(obj.X)+float32(w.OffsetX),
        float32(obj.Y)+float32(w.OffsetY)+floatOffset-20,
        frameWidth*float32(obj.Scale),
        frameHeight*float32(obj.Scale),
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

func (w *Weapon) Animate(animationDelay int, framesX, framesY []int){
	w.Object.UpdateAnimation(50, []int{0, 1}, []int{1, 1}) 
}


package player

import (
	"otaviocosta2110/getTheBlueBlocks/src/sprites"
	"otaviocosta2110/getTheBlueBlocks/src/system"
	"time"
)

type Player struct {
	Object          system.Object
	Speed           int32
	Health          int32
	MaxHealth       int32
	Flipped         bool
	Scale           int32
	LastDamageTaken time.Time
}

func NewPlayer(x, y, width, height, speed, scale int32, sprite sprites.Sprite) *Player {
	return &Player{
		Object: system.Object{
			X:             x,
			Y:             y,
			Width:         width * scale / 2,
			Height:        height * scale,
			KnockbackX:    0,
			KnockbackY:    0,
			FrameY:        0,
			FrameX:        0,
			LastFrameTime: time.Now(),
      Sprite: sprites.Sprite{
        SpriteWidth:  width,
        SpriteHeight: height,
        Texture:      sprite.Texture,
      },
		},
		Speed:  speed,
		Flipped:         false,
		Scale:           scale,
		MaxHealth:       3,
		Health:          3,
		LastDamageTaken: time.Now(),
	}
}

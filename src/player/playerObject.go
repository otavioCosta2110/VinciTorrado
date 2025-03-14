package player

import (
	"otaviocosta2110/getTheBlueBlocks/src/system"
	"time"
)

type Player struct {
	Object          system.Object
	Points          int32
	Speed           int32
	Health          int32
	MaxHealth       int32
	Sprite          system.Sprite
	Flipped         bool
	Scale           int32
	FrameX          int32
	FrameY          int32
	LastFrameTime   time.Time
	LastDamageTaken time.Time
}

func NewPlayer(x, y, width, height, points, speed, scale int32, sprite system.Sprite) *Player {
	return &Player{
		Object: system.Object{
			X:      x,
			Y:      y,
			Width:  width * scale / 2,
			Height: height * scale,
		},
		Points:          points,
		Speed:           speed,
		Sprite:          system.Sprite{
      SpriteWidth: width,
      SpriteHeight: height,
      Texture: sprite.Texture,
    },
		Flipped:         false,
		Scale:           scale,
		FrameY:          0,
		FrameX:          0,
		MaxHealth:       3,
		Health:          3,
		LastFrameTime:   time.Now(),
		LastDamageTaken: time.Now(),
	}
}

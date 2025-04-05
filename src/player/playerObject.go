package player

import (
	"otaviocosta2110/getTheBlueBlocks/src/sprites"
	"otaviocosta2110/getTheBlueBlocks/src/system"
	"time"
)

type Player struct {
	system.LiveObject
	IsKicking bool
}

func NewPlayer(x, y, width, height, speed, scale int32, sprite sprites.Sprite) *Player {
	return &Player{
		LiveObject: system.LiveObject{
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
				Scale: scale,
			},
			Speed:           speed,
			Flipped:         false,
			MaxHealth:       5,
			Health:          5,
			LastDamageTaken: time.Now(),
		},
		IsKicking:       false,
	}
}

func (p *Player) GetObject() system.Object {
	return p.Object
}

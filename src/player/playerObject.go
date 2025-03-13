package player

import (
	"otaviocosta2110/getTheBlueBlocks/src/system"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Object          system.Object
	Points          int32
	Speed           int32
	Health          int32
	MaxHealth       int32
	Sprite          rl.Texture2D
	Flipped         bool
	Scale           int32
	FrameX          int32
	FrameY          int32
	LastFrameTime   time.Time
	LastDamageTaken time.Time
}

func NewPlayer(x, y, width, height, points, speed, scale int32, sprite rl.Texture2D) *Player {
	return &Player{
		Object: system.Object{
			X:      x,
			Y:      y,
			Width:  width * scale,
			Height: height * scale,
		},
		Points:          points,
		Speed:           speed,
		Sprite:          sprite,
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

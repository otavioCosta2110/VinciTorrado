package player

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	X               int32
	Y               int32
	Width           int32
	Height          int32
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
		X:               x,
		Y:               y,
		Width:           width,
		Height:          height,
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

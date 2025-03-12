package player

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	X       int32
	Y       int32
	Width   int32
	Height  int32
	Points  int32
	Speed   int32
	Sprite  rl.Texture2D
	Flipped bool
	Scale   int32
	FrameX  int32
	FrameY  int32
  LastFrameTime time.Time
}

func NewPlayer(x, y, width, height, points, speed, scale int32, sprite rl.Texture2D) *Player {
	return &Player{
		X:       x,
		Y:       y,
		Width:   width,
		Height:  height,
		Points:  points,
		Speed:   speed,
		Sprite:  sprite,
		Flipped: false,
		Scale:   scale,
    FrameY: 0,
    FrameX: 0,
    LastFrameTime: time.Now(),
	}
}

func (p *Player) DrawPlayer() {
    var width float32 = 32
    if p.Flipped {
        width = -float32(width) 
    }

    sourceRec := rl.NewRectangle(
        float32(p.FrameX)*32,  
        float32(p.FrameY)*32,  
        width,           
        float32(32),
    )

    destinationRec := rl.NewRectangle(
        float32(p.X),
        float32(p.Y),
        float32(p.Width)*float32(p.Scale),
        float32(p.Height)*float32(p.Scale),
    )

    origin := rl.NewVector2(
        destinationRec.Width/2, 
        destinationRec.Height/2,
    )

    rl.DrawTexturePro(p.Sprite, sourceRec, destinationRec, origin, 0.0, rl.White)

    // faz a caixa vermelha pra ver colisao
    rl.DrawRectangleLines(
        int32(destinationRec.X - origin.X /2), 
        int32(destinationRec.Y - origin.Y), 
        int32(destinationRec.Width/2),
        int32(destinationRec.Height),
        rl.Red,
    )
}


package player
import(
  rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct{
  X int32
  Y int32
  Width int32
  Height int32
  Points int32
  Speed int32
}

func NewPlayer(x, y, width, height, points, speed int32) *Player {
	return &Player {
    X: x,
    Y: y,
    Width: width,
    Height: height,
    Points: points,
    Speed: speed,
  }
}

func (p *Player) DrawPlayer() {
	rl.DrawRectangle(p.X, p.Y, p.Width, p.Height, rl.Black)
}

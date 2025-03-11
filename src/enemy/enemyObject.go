package enemy

import(
  rl "github.com/gen2brain/raylib-go/raylib"
)

type Enemy struct{
  X int32
  Y int32
  Speed int32
  Width int32
  Height int32
}

func NewEnemy(x, y, speed, width, height int32) *Enemy {
	return &Enemy {
    X: x,
    Y: y,
    Width: width,
    Speed: speed,
    Height: height,
  }
}

func (e *Enemy) DrawEnemy(){
	rl.DrawRectangle(e.X, e.Y, e.Width, e.Height, rl.Red)
}

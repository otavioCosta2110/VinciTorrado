package points

import (
	"math/rand"
	"otaviocosta2110/ray/src/screen"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Point struct{
  X int32
  Y int32
}

func NewPoint(s screen.Screen) (Point) {
  x := rand.Intn(int(s.Width))
  y := rand.Intn(int(s.Height))
  return Point{
    X: int32(x),
    Y: int32(y),
  }
}

func (p Point) DrawPoint(){
	rl.DrawRectangle(p.X, p.Y, 10, 10, rl.Blue)
}

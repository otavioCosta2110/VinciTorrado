package screen

import rl "github.com/gen2brain/raylib-go/raylib"

type Screen struct {
	Width  int32
	Height int32
	Title  string
	Camera rl.Camera2D
}

func NewScreen(width, height int32, title string) *Screen {
	return &Screen{
		Width:  width,
		Height: height,
		Title:  title,
		Camera: rl.Camera2D{
			Offset:   rl.NewVector2(float32(width)/2, float32(height)/2),
			Target:   rl.NewVector2(0, 0),
			Rotation: 0,
			Zoom:     1,
		},
	}
}

func (s *Screen) UpdateCamera(targetX, targetY int32) {
	s.Camera.Target = rl.NewVector2(float32(targetX), float32(targetY))
}

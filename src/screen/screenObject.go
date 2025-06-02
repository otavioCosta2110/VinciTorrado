package screen

import rl "github.com/gen2brain/raylib-go/raylib"

type Screen struct {
	Width         int32
	Height        int32
	Title         string
	ScenaryWidth  int32
	ScenaryHeight int32
	Camera        rl.Camera2D
}

func NewScreen(width, height, scenaryWidth, scenaryHeight int32, title string) *Screen {
	return &Screen{
		Width:         width,
		Height:        height,
		Title:         title,
		ScenaryWidth:  scenaryWidth,
		ScenaryHeight: scenaryHeight,
		Camera: rl.Camera2D{
			Offset:   rl.NewVector2(float32(width)/2, float32(height)/2),
			Target:   rl.NewVector2(0, 0),
			Rotation: 0,
			Zoom:     1,
		},
	}
}

func (s *Screen) InitCamera(targetX, targetY int32) {
	s.Camera.Target = rl.NewVector2(float32(s.Width/2), float32(s.Height)/2)
}

func (s *Screen) UpdateCamera(targetX, targetY int32, canAdvance bool) {
	if canAdvance {
		buildingMinX := int32(0)
		buildingMaxX := int32(s.ScenaryWidth)

		camX := float32(targetX)

		halfWidth := float32(s.Width) / 2

		lerpFactor := float32(0.1)
		if camX < float32(buildingMinX)+halfWidth {
			camX = float32(buildingMinX) + halfWidth
		} else if camX > float32(buildingMaxX)-halfWidth {
			camX = float32(buildingMaxX) - halfWidth
		}

		currentX := s.Camera.Target.X
		s.Camera.Target.X = currentX + (camX-currentX)*lerpFactor
	}
}

func (s *Screen) ResetCamera() {
	s.Camera.Target = rl.NewVector2(float32(s.Width/2), float32(s.Height/2))
	s.Camera.Offset = rl.NewVector2(float32(s.Width/2), float32(s.Height/2))
}

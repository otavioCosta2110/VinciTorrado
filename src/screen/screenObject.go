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
	// Define building boundaries (adjust these to your actual building size)
	buildingMinX := int32(0)
	buildingMaxX := int32(640) * 4 // Building width

	// Calculate desired camera position (centered on player)
	camX := float32(targetX)

	// Clamp camera to building boundaries
	halfWidth := float32(s.Width) / 2

	// Clamp X position (left/right edges)
	if camX < float32(buildingMinX)+halfWidth {
		camX = float32(buildingMinX) + halfWidth 
	} else if camX > float32(buildingMaxX)-halfWidth {
		camX = float32(buildingMaxX) - halfWidth
	}


	println("targetX", targetX, camX)
	// Update camera target (WITHOUT overriding Y position)
	s.Camera.Target = rl.NewVector2(camX, float32(s.Height)/2)
}

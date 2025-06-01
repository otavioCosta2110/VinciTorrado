package system

import rl "github.com/gen2brain/raylib-go/raylib"

func LoadScaledTexture(path string, scale int32) rl.Texture2D {
	texture := rl.LoadTexture(path)
	texture.Width *= scale
	texture.Height *= scale
	return texture
}


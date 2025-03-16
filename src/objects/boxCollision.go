package objects

import (
	"otaviocosta2110/getTheBlueBlocks/src/physics"
	"otaviocosta2110/getTheBlueBlocks/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func CheckPlayerBoxCollision(playerObj *system.Object, box *Box) {
	boxObj := system.Object{
		X:      box.X,
		Y:      box.Y,
		Width:  box.Width,
		Height: box.Height,
	}

	// Verifica colisão entre o jogador e a caixa
	if physics.CheckCollision(*playerObj, boxObj) {
		resolveCollision(playerObj, &boxObj)
	}
}

// resolveCollision ajusta a posição do jogador para fora da caixa
func resolveCollision(playerObj, boxObj *system.Object) {
	// Calcula a sobreposição entre o jogador e a caixa
	overlapX := (playerObj.Width+boxObj.Width)/2 - abs(playerObj.X-boxObj.X)
	overlapY := (playerObj.Height+boxObj.Height)/2 - abs(playerObj.Y-boxObj.Y)

	if overlapX < overlapY {
		// Colisão horizontal
		if playerObj.X < boxObj.X {
			playerObj.X -= overlapX // Move o jogador para a esquerda
		} else {
			playerObj.X += overlapX // Move o jogador para a direita
		}
	} else {
		// Colisão vertical
		if playerObj.Y < boxObj.Y {
			playerObj.Y -= overlapY
		} else {
			playerObj.Y += overlapY
		}
	}

}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func DrawDebugCollision(playerObj *system.Object, box *Box) {
	rl.DrawRectangleLines(playerObj.X, playerObj.Y, playerObj.Width, playerObj.Height, rl.Blue)

	rl.DrawRectangleLines(box.X, box.Y, box.Width, box.Height, rl.Red)
}

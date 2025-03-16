package physics

import "otaviocosta2110/getTheBlueBlocks/src/system"

func CheckCollision(obj1, obj2 system.Object) bool {
	return obj1.X < obj2.X+obj2.Width && obj1.X+obj1.Width > obj2.X &&
		obj1.Y < obj2.Y+obj2.Height && obj1.Y+obj1.Height > obj2.Y
}

// ResolveCollision move o objeto b com base na colisão com o objeto a
func ResolveCollision(a, b *system.Object) {
	overlapX := (a.Width+b.Width)/2 - abs(a.X-b.X)
	overlapY := (a.Height+b.Height)/2 - abs(a.Y-b.Y)

	if overlapX < overlapY {
		// Colisão horizontal
		if a.X < b.X {
			b.X += overlapX
		} else {
			b.X -= overlapX
		}
	} else {
		if a.Y < b.Y {
			b.Y += overlapY
		} else {
			b.Y -= overlapY
		}
	}
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

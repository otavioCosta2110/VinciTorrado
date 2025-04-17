package physics

import "otaviocosta2110/getTheBlueBlocks/src/system"

func CheckCollision(obj1, obj2 system.Object) bool {
	return obj1.X-obj1.Width/2 < obj2.X+obj2.Width/2 &&
		obj1.X+obj1.Width/2 > obj2.X-obj2.Width/2 &&
		obj1.Y-obj1.Height/2 < obj2.Y+obj2.Height/2 &&
		obj1.Y+obj1.Height/2 > obj2.Y-obj2.Height/2
}

func ResolveCollision(a, b *system.Object) {
	overlapX := (a.Width/2 + b.Width/2) - abs(a.X-b.X)
	overlapY := (a.Height/2 + b.Height/2) - abs(a.Y-b.Y)

	if overlapX <= 0 || overlapY <= 0 {
		return
	}

	if overlapX < overlapY {
		if a.X < b.X {
			a.X -= overlapX / 4
			b.X += overlapX / 4
		} else {
			a.X += overlapX / 4
			b.X -= overlapX / 4
		}
	} else {
		if a.Y < b.Y {
			a.Y -= overlapY / 4
			b.Y += overlapY / 4
		} else {
			a.Y += overlapY / 4
			b.Y -= overlapY / 4
		}
	}
}
func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

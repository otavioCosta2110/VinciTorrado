package physics

import (
	"math"
	"otaviocosta2110/vincitorrado/src/system"
)

func CheckCollision(obj1, obj2 system.Object) bool {
	return obj1.X-obj1.Width/2 < obj2.X+obj2.Width/2 &&
		obj1.X+obj1.Width/2 > obj2.X-obj2.Width/2 &&
		obj1.Y-obj1.Height/2 < obj2.Y+obj2.Height/2 &&
		obj1.Y+obj1.Height/2 > obj2.Y-obj2.Height/2
}

func ResolveCollision(a, b *system.Object) {
	overlapX := float64(a.Width/2 + b.Width/2) - math.Abs(float64(a.X-b.X))
	overlapY := float64(a.Height/2 + b.Height/2) - math.Abs(float64(a.Y-b.Y))

	if overlapX <= 0 || overlapY <= 0 {
		return
	}

	if overlapX < overlapY {
		if a.X < b.X {
			a.X -= int32(overlapX)
			b.X += int32(overlapX)
		} else {
			a.X += int32(overlapX)
			b.X -= int32(overlapX)
		}
	} else {
		if a.Y < b.Y {
			a.Y -= int32(overlapY)
			b.Y += int32(overlapY)
		} else {
			a.Y += int32(overlapY)
			b.Y -= int32(overlapY)
		}
	}
}

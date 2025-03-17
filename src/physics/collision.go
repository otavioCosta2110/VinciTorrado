package physics

import "otaviocosta2110/getTheBlueBlocks/src/system"

func CheckCollision(obj1, obj2 system.Object) bool {
	return obj1.X < obj2.X+obj1.Width && obj1.X+obj1.Width > obj2.X &&
		obj1.Y < obj2.Y+obj1.Height && obj1.Y+obj1.Height > obj2.Y
}

func ResolveCollision(a, b *system.Object) {
	overlapX := (a.Width/2 + b.Width/2) - abs(a.X-b.X)
	overlapY := (a.Height/2 + b.Height/2) - abs(a.Y-b.Y)

	if overlapX > 0 && overlapY > 0 {
		if overlapY > overlapX {
			if a.X < b.X {
				a.X -= overlapX
			} else {
				a.X += overlapX
			}
		} else {
      if a.Y < b.Y {
        a.Y -= overlapY
      }else{
        a.Y += overlapY
      }
		}
	}
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

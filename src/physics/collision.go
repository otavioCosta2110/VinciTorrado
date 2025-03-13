package physics

import "otaviocosta2110/getTheBlueBlocks/src/system"

func CheckCollision(obj1, obj2 system.Object) bool {
	return obj1.X < obj2.X+obj1.Width && obj1.X+obj1.Width > obj2.X &&
		obj1.Y < obj2.Y+obj1.Height && obj1.Y+obj1.Height > obj2.Y
}

// essa func faz o obj a mexer o b, da pra usar pra empurrar alguma coisa
func ResolveCollision(a, b *system.Object) {
	
	overlapX := (a.Width + b.Width)/2 - abs(a.X-b.X)
	overlapY := (a.Height + b.Height)/2 - abs(a.Y-b.Y)

	
	if overlapX < overlapY {
		
		if a.X < b.X {
			a.X -= overlapX 
		} else {
			a.X += overlapX 
		}
	} else {
		
		if a.Y < b.Y {
			a.Y -= overlapY 
		} else {
			a.Y += overlapY 
		}
	}
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

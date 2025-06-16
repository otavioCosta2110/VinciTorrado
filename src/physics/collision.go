package physics

import (
	"otaviocosta2110/vincitorrado/src/system"
)

func CheckCollision(obj1, obj2 system.Object) bool {
	return obj1.X-obj1.Width/2 < obj2.X+obj2.Width/2 &&
		obj1.X+obj1.Width/2 > obj2.X-obj2.Width/2 &&
		obj1.Y-obj1.Height/2 < obj2.Y+obj2.Height/2 &&
		obj1.Y+obj1.Height/2 > obj2.Y-obj2.Height/2
}

package physics

import (
	"otaviocosta2110/getTheBlueBlocks/src/system"
)

func TakeKnockback(obj *system.Object) {
	obj.X += obj.KnockbackX
	obj.Y += obj.KnockbackY

	if obj.KnockbackX > 0 {
		obj.KnockbackX -= 1
		if obj.KnockbackX < 0 {
			obj.KnockbackX = 0
		}
	} else if obj.KnockbackX < 0 {
		obj.KnockbackX += 1
		if obj.KnockbackX > 0 {
			obj.KnockbackX = 0
		}
	}

	if obj.KnockbackY > 0 {
		obj.KnockbackY -= 1
		if obj.KnockbackY < 0 {
			obj.KnockbackY = 0
		}
	} else if obj.KnockbackY < 0 {
		obj.KnockbackY += 1
		if obj.KnockbackY > 0 {
			obj.KnockbackY = 0
		}
	}
}

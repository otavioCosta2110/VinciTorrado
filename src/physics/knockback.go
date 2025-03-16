package physics

import (
	"otaviocosta2110/getTheBlueBlocks/src/system"
)


func TakeKnockback(obj *system.Object) {
	obj.X += obj.KnockbackX
	obj.Y += obj.KnockbackY

	if obj.KnockbackX > 0 {
		obj.KnockbackX--
	} else if obj.KnockbackX < 0 {
		obj.KnockbackX++
	}

	if obj.KnockbackY > 0 {
		obj.KnockbackY--
	} else if obj.KnockbackY < 0 {
		obj.KnockbackY++
	}
}

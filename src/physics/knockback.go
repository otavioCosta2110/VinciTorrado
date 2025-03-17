package physics

import (
	"otaviocosta2110/getTheBlueBlocks/src/system"
)


func TakeKnockback(obj *system.Object) {
	// Apply knockback movement to object
	obj.X += obj.KnockbackX
	obj.Y += obj.KnockbackY

	// Reduce knockback strength gradually
	// Reduce knockbackX (horizontally) to zero, if not already stopped
	if obj.KnockbackX > 0 {
		obj.KnockbackX -= 1 // Gradually decrease in positive direction
		if obj.KnockbackX < 0 { // Ensure it doesn't overshoot
			obj.KnockbackX = 0
		}
	} else if obj.KnockbackX < 0 {
		obj.KnockbackX += 1 // Gradually decrease in negative direction
		if obj.KnockbackX > 0 { // Ensure it doesn't overshoot
			obj.KnockbackX = 0
		}
	}

	// Reduce knockbackY (vertically) to zero, if not already stopped
	if obj.KnockbackY > 0 {
		obj.KnockbackY -= 1 // Gradually decrease in positive direction
		if obj.KnockbackY < 0 { // Ensure it doesn't overshoot
			obj.KnockbackY = 0
		}
	} else if obj.KnockbackY < 0 {
		obj.KnockbackY += 1 // Gradually decrease in negative direction
		if obj.KnockbackY > 0 { // Ensure it doesn't overshoot
			obj.KnockbackY = 0
		}
	}
}


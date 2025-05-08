package physics

import (
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/system"
)

type Kickable interface {
	HandleKick(
		kickHitbox system.Object,
		items *[]*equipment.Equipment,
		isFlipped bool,
		kickPower int32,
	) bool
	GetObject() system.Object
}

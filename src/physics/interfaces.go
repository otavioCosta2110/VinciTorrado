package physics

import (
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/system"
)

type Kickable interface {
	IsKicked() bool
	HandleKick(items *[]*equipment.Equipment, playerObject system.Object)
	GetObject() system.Object
}

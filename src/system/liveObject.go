package system

import (
	"time"
)

type LiveObject struct {
	Object Object
	Speed           int32
	Health          int32
	MaxHealth       int32
	LastDamageTaken time.Time
	Flipped         bool
}

type Live interface  {
	Draw()
	TakeDamage(damage int32, eX int32, eY int32)
	GetObject() Object
	SetObject(Object)
}


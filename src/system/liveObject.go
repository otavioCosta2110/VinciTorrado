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
	TakeDamage(damage int32, eX, eY int32)
	GetObject() Object
	SetObject(Object)
}

type Player interface {
    GetObject() Object
    TakeDamage(damage int32, eobj Object)
}

type Enemy interface {
    GetObject() Object
    TakeDamage(damage int32, pX int32, pY int32)
}

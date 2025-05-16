package system

import (
	"time"
)

type LiveObject struct {
	Object          Object
	Speed           int32
	Health          int32
	MaxHealth       int32
	Damage          int32
	LastDamageTaken time.Time
}

type Live interface {
	Draw()
	TakeDamage(damage int32, obj Object)
	UpdateAnimation(animationName string)
	GetObject() Object
	SetObject(Object)
	IsActive() bool
	SetActive(active bool) 
}

type Player interface {
	GetObject() Object
	TakeDamage(damage int32, eobj Object)
}

type Enemy interface {
	GetObject() Object
	TakeDamage(damage int32, pX int32, pY int32)
}

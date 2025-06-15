package system

import (
	"otaviocosta2110/vincitorrado/src/sprites"
	"time"
)

type Object struct {
	X, Y           int32
	Width          int32
	Height         int32
	KnockbackX     int32
	KnockbackY     int32
	FrameX         int32
	FrameY         int32
	LastFrameTime  time.Time
	LastAttackTime time.Time
	Sprite         sprites.Sprite
	Scale          int32
	Destroyed      bool
	Flipped        bool
}

func (obj *Object) UpdateAnimation(animationDelay int32, framesX, framesY []int32) {
	if len(framesX) == 0 || len(framesY) == 0 || len(framesX) != len(framesY) {
		return // invalid frame data
	}

	if time.Since(obj.LastFrameTime) < time.Duration(animationDelay)*time.Millisecond {
		return
	}

	// Find current position in animation sequence
	var found bool
	for i := range framesX {
		if obj.FrameX == framesX[i] && obj.FrameY == framesY[i] {
			next := (i + 1) % len(framesX)
			obj.FrameX = framesX[next]
			obj.FrameY = framesY[next]
			found = true
			break
		}
	}

	if !found {
		obj.FrameX = framesX[0]
		obj.FrameY = framesY[0]
	}

	obj.LastFrameTime = time.Now()
}

func (obj *Object) SetKnockback(attackingObj Object) {
	knockbackStrengthX := int32(10)
	knockbackStrengthY := int32(10)

	if obj.X < attackingObj.X {
		obj.KnockbackX = -knockbackStrengthX
	} else {
		obj.KnockbackX = knockbackStrengthX
	}

	if obj.Y < attackingObj.Y {
		obj.KnockbackY = -knockbackStrengthY
	} else {
		obj.KnockbackY = knockbackStrengthY
	}

}

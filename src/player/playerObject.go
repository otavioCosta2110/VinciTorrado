package player

import (
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	system.LiveObject
	IsKicking     bool
	LastKickTime  time.Time
	KickCooldown  time.Duration
	KickPower     int32
	Equipment     *equipment.Equipment
	HatSprite     sprites.Sprite // Stores just the hat spritesheet
	OriginalSpeed int32          // Store base speed for power-ups
}

func NewPlayer(x, y, width, height, speed, scale int32, sprite sprites.Sprite) *Player {
	// Load hat spritesheet (transparent PNG with just the turbante)
	hatSprite := sprites.Sprite{
		SpriteWidth:  width,
		SpriteHeight: height,
		Texture:      rl.LoadTexture("assets/player/playerTurbante.png"),
	}

	return &Player{
		LiveObject: system.LiveObject{
			Object: system.Object{
				X:             x,
				Y:             y,
				Width:         width * scale / 2,
				Height:        height * scale,
				KnockbackX:    0,
				KnockbackY:    0,
				FrameY:        0,
				FrameX:        0,
				LastFrameTime: time.Now(),
				Sprite:        sprite,
				Scale:         scale,
			},
			Speed:           speed,
			Flipped:         false,
			MaxHealth:       5,
			Health:          5,
			LastDamageTaken: time.Now(),
			Damage:          1,
		},
		IsKicking:     false,
		LastKickTime:  time.Now(),
		KickCooldown:  500 * time.Millisecond,
		KickPower:     15,
		HatSprite:     hatSprite,
		OriginalSpeed: speed, // Store base speed
	}
}

func (p *Player) GetObject() system.Object {
	return p.Object
}

func (p *Player) Equip(item *equipment.Equipment) {
	if p.Equipment != nil {
		p.Unequip()
	}
	p.Equipment = item
	p.Equipment.IsEquipped = true
}

func (p *Player) Unequip() {
	if p.Equipment != nil {
		// Restore original stats
		p.Speed = p.OriginalSpeed
		p.Equipment.IsEquipped = false
		p.Equipment = nil
	}
}

func (p *Player) HasEquipment() bool {
	return p.Equipment != nil
}

func (p *Player) Cleanup() {
	rl.UnloadTexture(p.HatSprite.Texture)
}

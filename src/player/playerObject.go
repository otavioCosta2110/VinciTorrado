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
	IsKicking         bool
	LastKickTime      time.Time
	KickCooldown      time.Duration
	KickPower         int32
	Equipped          *equipment.Equipment
	Equipment         []*equipment.Equipment
	HatSprite         sprites.Sprite
	OriginalSpeed     int32
	OriginalMaxHealth int32
	OriginalDamage    int32
}

func NewPlayer(x, y, width, height, speed, scale int32, sprite sprites.Sprite) *Player {

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
		IsKicking:         false,
		LastKickTime:      time.Now(),
		KickCooldown:      500 * time.Millisecond,
		KickPower:         15,
		OriginalSpeed:     speed,
		OriginalMaxHealth: 5,
		OriginalDamage:    1,
	}
}

func (p *Player) GetObject() system.Object {
	return p.Object
}

func (p *Player) AddToInventory(item *equipment.Equipment) {
	p.Equipment = append(p.Equipment, item)
}

func (p *Player) Equip(item *equipment.Equipment) {
	if p.Equipped != nil {
		p.Unequip()
	}
	p.Equipped = item
	p.Equipped.IsEquipped = true

	p.MaxHealth += p.Equipped.Stats.Life
	p.Damage += p.Equipped.Stats.Damage
	p.Speed += p.Equipped.Stats.Speed

	println(p.MaxHealth)

	p.HatSprite = item.Object.Sprite
}

func (p *Player) Unequip() {
	if p.Equipped != nil {
		p.MaxHealth -= p.Equipped.Stats.Life
		p.Damage -= p.Equipped.Stats.Damage
		p.Speed -= p.Equipped.Stats.Speed

		if p.Health > p.MaxHealth {
			p.Health = p.MaxHealth
		}

		p.Equipped.IsEquipped = false
		p.Equipped = nil
	}
}
func (p *Player) HasEquipment() bool {
	return p.Equipped != nil
}

func (p *Player) Cleanup() {
	rl.UnloadTexture(p.HatSprite.Texture)
}

func (p *Player) UseConsumable(index int) {
	if index >= 0 && index < len(p.Equipment) {
		item := p.Equipment[index]
		if item.Type == "consumable" {
			// Apply healing
			p.Health = p.Health + item.Stats.Heal
			if p.Health > p.MaxHealth {
				p.Health = p.MaxHealth
			}

			// Remove from inventory
			p.Equipment = append(p.Equipment[:index], p.Equipment[index+1:]...)
		}
	}
}

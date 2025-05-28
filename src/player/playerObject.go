package player

import (
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"
	"otaviocosta2110/vincitorrado/src/weapon"
	"time"

	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	system.LiveObject
	IsKicking      bool
	IsAttacking    bool
	LastKickTime   time.Time
	KickCooldown   time.Duration
	KickPower      int32
	Equipped       *equipment.Equipment
	Equipment      []*equipment.Equipment
	Consumables    []*equipment.Equipment
	HatSprite      sprites.Sprite
	Screen         *screen.Screen
	Weapon         *weapon.Weapon
	LastAttackTime time.Time
	Projectiles []*weapon.Projectile
}

func NewPlayer(x, y, width, height, speed, scale int32, sprite sprites.Sprite, s *screen.Screen) *Player {

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
				Flipped:       false,
			},
			Speed:           speed,
			MaxHealth:       5,
			Health:          5,
			LastDamageTaken: time.Now(),
			Damage:          1,
		},
		IsKicking:    false,
		IsAttacking:  false,
		LastKickTime: time.Now(),
		KickCooldown: 500 * time.Millisecond,
		KickPower:    15,
		Screen:       s,
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
			p.Health = min(p.Health+item.Stats.Heal, p.MaxHealth)

			p.Equipment = slices.Delete(p.Equipment, index, index+1)
		}
	}
}

func (p *Player) PickUp(w weapon.Weapon) {
	if p.Weapon != nil {
		p.DropWeapon()
	}

	p.Weapon = &w
	p.Weapon.IsDropped = false
	p.Weapon.IsEquipped = true

	p.Damage += w.Stats.Damage
}

func (p *Player) DropWeapon() {
	if p.Weapon != nil {
		p.Damage -= p.Weapon.Stats.Damage

		p.Weapon.Object.X = p.Object.X
		p.Weapon.Object.Y = p.Object.Y
		p.Weapon.IsDropped = true
		p.Weapon.IsEquipped = false

		p.Weapon = nil
	}
}

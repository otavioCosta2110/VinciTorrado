package ui

import (
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/player"
	"otaviocosta2110/vincitorrado/src/sprites"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MenuState int

const (
	MenuMain MenuState = iota
	MenuEquipment
	MenuStats
)

type EquipmentSlot struct {
	Rect     rl.Rectangle
	IsActive bool
	IconPos  rl.Rectangle
	Item     *equipment.Equipment
}

type Menu struct {
	State           MenuState
	IsVisible       bool
	PlayerSprite    *sprites.Sprite
	PlayerReference *player.Player
	EquipmentSlots  []EquipmentSlot
	SelectedSlot    int
	Columns         int
	IconSheet       rl.Texture2D
	ConsumableSheet rl.Texture2D
	SlotWidth       float32
	SlotHeight      float32
	SlotSpacing     float32
}

func NewMenu(player *player.Player, sprite *sprites.Sprite) *Menu {
	menu := &Menu{
		State:           MenuMain,
		IsVisible:       false,
		PlayerSprite:    sprite,
		PlayerReference: player,
		SelectedSlot:    0,
		Columns:         2,
		IconSheet:       rl.LoadTexture("assets/ui/equipamentos.png"),
		ConsumableSheet: rl.LoadTexture("assets/ui/items.png"),
		SlotWidth:       120,
		SlotHeight:      120,
		SlotSpacing:     50,
	}

	menu.initEquipmentSlots()
	return menu
}

func (m *Menu) initEquipmentSlots() {
	startX := float32(rl.GetScreenWidth())/2 - (float32(m.Columns)*m.SlotWidth)/2 + 50
	startY := float32(rl.GetScreenHeight() / 5)

	inventorySize := len(m.PlayerReference.Equipment)

	if inventorySize == 0 {
		inventorySize = 0
	}

	m.EquipmentSlots = make([]EquipmentSlot, 0, inventorySize)

	for i := range inventorySize {
		row := i / m.Columns
		col := i % m.Columns

		var item *equipment.Equipment
		var iconPos rl.Rectangle

		if i < len(m.PlayerReference.Equipment) {
			item = m.PlayerReference.Equipment[i]
			iconPos = m.getItemIconPos(item)
		} else {
			iconPos = rl.NewRectangle(0, 0, 32, 32)
		}

		m.EquipmentSlots = append(m.EquipmentSlots, EquipmentSlot{
			Rect: rl.NewRectangle(
				startX+float32(col)*(m.SlotWidth+m.SlotSpacing),
				startY+float32(row)*(m.SlotHeight+m.SlotSpacing),
				m.SlotWidth,
				m.SlotHeight,
			),
			IsActive: false,
			IconPos:  iconPos,
			Item:     item,
		})
	}

	m.SelectedSlot = -1
	for i, slot := range m.EquipmentSlots {
		if slot.Item != nil {
			m.SelectedSlot = i
			break
		}
	}
}

func (m *Menu) getItemIconPos(item *equipment.Equipment) rl.Rectangle {
	if item == nil {
		return rl.NewRectangle(0, 0, 32, 32)
	}

	println("Item Name:", item.Name, "Type:", item.Type)
	switch {
	case item.Name == "Turbante":
		return rl.NewRectangle(32, 32, 32, 32)
	case item.Type == "consumable":
		switch item.Name {
		case "Hamburgão":
			return rl.NewRectangle(0, 0, 32, 32)
		case "Saunduiche":
			return rl.NewRectangle(32, 0, 32, 32)
		case "Cachaca":
			println("Cachaca icon")
			return rl.NewRectangle(0, 32, 32, 32)
		case "guarana":
			println("Guarana icon")
			return rl.NewRectangle(0, 64, 32, 32)
		default:
			return rl.NewRectangle(64, 0, 32, 32)
		}
	default:
		return rl.NewRectangle(0, 0, 32, 32)
	}
}

func (m *Menu) Draw() {
	if !m.IsVisible {
		return
	}

	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Fade(rl.Black, 0.5))

	menuWidth := float32(rl.GetScreenWidth()) * 0.8
	menuHeight := float32(rl.GetScreenHeight()) * 0.8
	menuX := (float32(rl.GetScreenWidth()) - menuWidth) / 2
	menuY := (float32(rl.GetScreenHeight()) - menuHeight) / 2

	rl.DrawRectangleRounded(
		rl.NewRectangle(menuX, menuY, menuWidth, menuHeight),
		0.05, 10, rl.DarkGray,
	)

	playerPreviewX := menuX
	playerPreviewY := menuY * 2
	sourceRec := rl.NewRectangle(0, 0, float32(m.PlayerSprite.SpriteWidth), float32(m.PlayerSprite.SpriteHeight))

	destinationRec := rl.NewRectangle(
		playerPreviewX,
		playerPreviewY,
		float32(m.PlayerSprite.SpriteWidth)*float32(m.PlayerReference.Object.Scale*3),
		float32(m.PlayerSprite.SpriteHeight)*float32(m.PlayerReference.Object.Scale*3),
	)

	origin := rl.NewVector2(0, 0)

	rl.DrawTexturePro(
		m.PlayerSprite.Texture,
		sourceRec,
		destinationRec,
		origin,
		0,
		rl.White,
	)

	if m.PlayerReference.HasEquipment() {
		rl.DrawTexturePro(
			m.PlayerReference.HatSprite.Texture,
			sourceRec,
			destinationRec,
			origin,
			0.0,
			rl.White,
		)
	}

	for i, slot := range m.EquipmentSlots {
		color := rl.Gray
		textColor := rl.White

		if i == m.SelectedSlot {
			color = rl.White
			textColor = rl.White
		}

		if slot.Item == nil {
			color = rl.Fade(rl.Gray, 0.3)
			textColor = rl.Fade(rl.White, 0.5)
		}

		rl.DrawRectangleRounded(slot.Rect, 0.1, 5, color)

		if slot.Item != nil {
			var texture rl.Texture2D
			if slot.Item.Type == "consumable" {
				texture = m.ConsumableSheet
			} else {
				texture = m.IconSheet
			}

			rl.DrawTexturePro(
				texture,
				slot.IconPos,
				rl.NewRectangle(
					slot.Rect.X+slot.Rect.Width/2-42.5,
					slot.Rect.Y+slot.Rect.Height/2-42.5,
					85,
					85,
				),
				rl.NewVector2(0, 0),
				0,
				rl.White,
			)

			rl.DrawText(slot.Item.Name, int32(slot.Rect.X+2), int32(slot.Rect.Y-20), 20, textColor)
		} else {
			rl.DrawText("Empty", int32(slot.Rect.X+2), int32(slot.Rect.Y-20), 20, textColor)
		}
	}

	if m.PlayerReference.HasEquipment() {
		rl.DrawText("Press U to unequip", int32(menuWidth)/5, int32(menuHeight), 20, rl.White)
	}
	if m.SelectedSlot >= 0 && m.SelectedSlot < len(m.EquipmentSlots) && m.EquipmentSlots[m.SelectedSlot].Item != nil {
		item := m.EquipmentSlots[m.SelectedSlot].Item
		statsX := menuX + menuWidth - 250
		statsY := menuY + 50

		rl.DrawRectangleRounded(
			rl.NewRectangle(statsX, statsY, 200, 150),
			0.1, 5, rl.Fade(rl.Black, 0.7),
		)

		rl.DrawText(item.Name, int32(statsX+10), int32(statsY+35), 20, rl.Gold)

		yOffset := int32(60)
		if item.Stats.Life != 0 {
			rl.DrawText("Health: ", int32(statsX+10), int32(int32(statsY)+yOffset), 18, rl.White)
			rl.DrawText(formatStat(int(item.Stats.Life)), int32(statsX+100), int32(int32(statsY)+yOffset), 18, getStatColor(int(item.Stats.Life)))
			yOffset += 25
		}

		if item.Stats.Damage != 0 {
			rl.DrawText("Damage: ", int32(statsX+10), int32(int32(statsY)+yOffset), 18, rl.White)
			rl.DrawText(formatStat(int(item.Stats.Damage)), int32(statsX+100), int32(int32(statsY)+yOffset), 18, getStatColor(int(item.Stats.Damage)))
			yOffset += 25
		}

		if item.Stats.Speed != 0 {
			rl.DrawText("Speed: ", int32(statsX+10), int32(int32(statsY)+yOffset), 18, rl.White)
			rl.DrawText(formatStat(int(item.Stats.Speed)), int32(statsX+100), int32(int32(statsY)+yOffset), 18, getStatColor(int(item.Stats.Speed)))
			yOffset += 25
		}

		if item.Type == "consumable" {
			rl.DrawText("Heal: "+strconv.Itoa(int(item.Stats.Heal)),
				int32(statsX+10), int32(statsY)+yOffset+50, 20, rl.Green)
		}
	}
}

func formatStat(value int) string {
	if value > 0 {
		return "+" + strconv.Itoa(value)
	}
	return strconv.Itoa(value)
}

func getStatColor(value int) rl.Color {
	if value > 0 {
		return rl.Green
	} else if value < 0 {
		return rl.Red
	}
	return rl.White
}

func (m *Menu) Update() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		m.IsVisible = !m.IsVisible
		m.Refresh()
	}

	if !m.IsVisible {
		return
	}

	prevSelected := m.SelectedSlot

	if rl.IsKeyPressed(rl.KeyRight) {
		m.findNextValidSlot(1)
	} else if rl.IsKeyPressed(rl.KeyLeft) {
		m.findNextValidSlot(-1)
	} else if rl.IsKeyPressed(rl.KeyDown) {
		m.findNextValidSlot(m.Columns)
	} else if rl.IsKeyPressed(rl.KeyUp) {
		m.findNextValidSlot(-m.Columns)
	}

	if prevSelected != m.SelectedSlot && m.SelectedSlot >= 0 {
		menu_move_sound := rl.LoadSound("assets/sounds/menu_move.mp3")
		rl.PlaySound(menu_move_sound)
	}

	if rl.IsKeyPressed(rl.KeyEnter) && m.SelectedSlot >= 0 {
		item := m.EquipmentSlots[m.SelectedSlot].Item
		if item != nil {
			if item.Type == "consumable" {
				m.PlayerReference.UseConsumable(m.SelectedSlot)
				m.Refresh()
				menu_select_sound := rl.LoadSound("assets/sounds/menu_selected.mp3")
				rl.PlaySound(menu_select_sound)
			} else {
				if m.PlayerReference.Equipped != m.EquipmentSlots[m.SelectedSlot].Item {
					m.PlayerReference.Unequip()
				}
				if m.PlayerReference.Equipped != m.EquipmentSlots[m.SelectedSlot].Item {
					m.PlayerReference.Equip(item)
					menu_select_sound := rl.LoadSound("assets/sounds/menu_selected.mp3")
					rl.PlaySound(menu_select_sound)
				}
			}
		}
	}

	if rl.IsKeyPressed(rl.KeyU) && m.PlayerReference.HasEquipment() {
		m.PlayerReference.Unequip()
		menu_select_sound := rl.LoadSound("assets/sounds/menu_selected.mp3")
		rl.PlaySound(menu_select_sound)
	}
}

func (m *Menu) findNextValidSlot(step int) {
	if len(m.EquipmentSlots) == 0 {
		m.SelectedSlot = -1
		return
	}

	start := max(m.SelectedSlot, 0)

	for i := 1; i <= len(m.EquipmentSlots); i++ {
		next := (start + i*step) % len(m.EquipmentSlots)
		if next < 0 {
			next += len(m.EquipmentSlots)
		}

		if m.EquipmentSlots[next].Item != nil {
			m.SelectedSlot = next
			return
		}
	}

	if m.SelectedSlot >= len(m.EquipmentSlots) || m.EquipmentSlots[m.SelectedSlot].Item == nil {
		m.SelectedSlot = -1
	}
}

func (m *Menu) Refresh() {
	m.initEquipmentSlots()
}

package gear

import "fmt"

// ItemInfo holds details about a specific gear item
type ItemInfo struct {
	Name        string
	Description string
	Rarity      string
}

// GearData maps [gearType][subType] to item information
// First byte is gear type, second byte is the specific variant/rarity
var GearData = map[uint8]map[uint8]ItemInfo{
	0x00: { // Swords
		0x00: {Name: "Iron Sword", Description: "A basic iron blade", Rarity: "Common"},
		0x01: {Name: "Steel Sword", Description: "Sharp steel weapon", Rarity: "Uncommon"},
		0x02: {Name: "Mythril Sword", Description: "Legendary elven blade", Rarity: "Legendary"},
	},
	0x01: { // Bows
		0x00: {Name: "Short Bow", Description: "Quick and nimble", Rarity: "Common"},
		0x01: {Name: "Long Bow", Description: "Long range weapon", Rarity: "Uncommon"},
		0x02: {Name: "Elven Bow", Description: "Blessed by nature", Rarity: "Legendary"},
	},
	0x02: { // Shields
		0x00: {Name: "Wooden Shield", Description: "Basic protection", Rarity: "Common"},
		0x01: {Name: "Iron Shield", Description: "Sturdy defense", Rarity: "Uncommon"},
	},
}

type Gear struct {
	ItemInfo
	XP    uint32
	Flags []string
}

func DecodeGear(data []byte) (*Gear, error) {
	if len(data) != 7 {
		return nil, fmt.Errorf("expected 7 bytes, got %d", len(data))
	}

	gearType := data[0]
	subType := data[1]

	// Look up the item info from nested map
	typeMap, exists := GearData[gearType]
	if !exists {
		return nil, fmt.Errorf("unknown gear type: 0x%02x", gearType)
	}

	itemInfo, exists := typeMap[subType]
	if !exists {
		return nil, fmt.Errorf("unknown sub-type 0x%02x for gear type 0x%02x", subType, gearType)
	}

	// XP is 3 bytes: data[2], data[3], data[4]
	xp := uint32(data[2])<<16 | uint32(data[3])<<8 | uint32(data[4])

	// Flags from remaining bytes (if needed)
	var flags []string
	// ... flag parsing logic similar to player

	return &Gear{
		ItemInfo: itemInfo,
		XP:       xp,
		Flags:    flags,
	}, nil
}

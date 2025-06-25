package main

import (
	"fmt"

	"github.com/mikegio27/nfc-game/gear"
	"github.com/mikegio27/nfc-game/player"
)

func main() {
	// Example packed header: ID=1, Name=0, Level=5, XP=0x0003E8 (1000), Flags=0b00000110 (Poisoned, Burning)
	header := []byte{0x01, 0x00, 0x05, 0x00, 0x03, 0xE8, 0x06}

	player, err := player.DecodePlayerHeader(header)
	if err != nil {
		panic(err)
	}

	fmt.Printf("ID: %d\n", player.ID)
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("XP: %d\n", player.XP)
	fmt.Printf("Flags: %v\n", player.Flags)

	gearData := []byte{0x00, 0x02, 0x00, 0x03, 0xE8, 0x00, 0x00} // Gear type 0 (Swords), sub-type 2 (Mythril Sword), XP=1000
	gearItem, err := gear.DecodeGear(gearData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Gear Name: %s\n", gearItem.Name)
	fmt.Printf("Description: %s\n", gearItem.Description)
	fmt.Printf("Rarity: %s\n", gearItem.Rarity)
	fmt.Printf("XP: %d\n", gearItem.XP)
	fmt.Printf("Flags: %v\n", gearItem.Flags)
}

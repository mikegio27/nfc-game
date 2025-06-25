package main

import (
	"fmt"

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
}

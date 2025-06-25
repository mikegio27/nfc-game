package player

import (
	"fmt"
)

// PlayerHeader represents the human-readable decoded player header
type PlayerHeader struct {
	ID    uint8
	Name  string
	Level uint8
	XP    uint32
	Flags []string
}

// Example name index map
var nameTable = map[uint8]string{
	0x00: "Sparrow",
	0x01: "Talon",
	0x02: "Vex",
}

// Example status flags
var flagBits = map[uint8]string{
	0: "Dead",
	1: "Poisoned",
	2: "Burning",
	3: "Frozen",
	4: "Invisible",
	5: "Stunned",
	6: "Silenced",
	7: "Feared",
}

func DecodePlayerHeader(data []byte) (*PlayerHeader, error) {
	if len(data) != 7 {
		return nil, fmt.Errorf("expected 7 bytes, got %d", len(data))
	}

	id := data[0]
	nameIndex := data[1]
	level := data[2]

	// XP is 3 bytes: data[3], data[4], data[5]
	xp := uint32(data[3])<<16 | uint32(data[4])<<8 | uint32(data[5])

	// Flags
	flagByte := data[6]
	var flags []string
	for i := uint8(0); i < 8; i++ {
		if (flagByte>>i)&1 == 1 {
			flags = append(flags, flagBits[i])
		}
	}

	name, ok := nameTable[nameIndex]
	if !ok {
		name = fmt.Sprintf("Unknown (index %d)", nameIndex)
	}

	return &PlayerHeader{
		ID:    id,
		Name:  name,
		Level: level,
		XP:    xp,
		Flags: flags,
	}, nil
}

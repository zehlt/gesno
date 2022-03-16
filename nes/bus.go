package nes

import (
	"fmt"

	"github.com/zehlt/go6502"
)

const (
	RAM     = 0x0000
	RAM_END = 0x1FFF
	RAM_BUS = 0b00000111_11111111

	PPU     = 0x2000
	PPU_END = 0x3FFF
	PPU_BUS = 0b00100000_00000111
)

type NesBus struct {
	m *go6502.Mem
}

func (b NesBus) Read(addr uint16) uint8 {
	if addr >= RAM && addr <= RAM_END {
		shortAddr := addr & RAM_BUS
		return b.Read(shortAddr)
	} else if addr >= PPU && addr <= PPU {
		shortAddr := addr & PPU_BUS
		fmt.Println("PPU is not supported yet!!")
		return b.Read(shortAddr)
	}

	return b.m.Read(addr)
}

func (b NesBus) Write(addr uint16, data uint8) {
	b.m.Write(addr, data)
}

func (b NesBus) ReadWord(addr uint16) uint16 {
	return b.m.ReadWord(addr)
}

func (b NesBus) WriteWord(addr uint16, data uint16) {
	b.m.WriteWord(addr, data)
}

func (b NesBus) WriteBytes(addr uint16, data []uint8) {
	b.m.WriteBytes(addr, data)
}

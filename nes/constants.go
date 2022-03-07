package nes

const (
	BRK_IMP = 0x00
	LDA_IMM = 0xA9
	TAX_IMP = 0xAA
)

const (
	Carry     = 0b0000_0001
	Zero      = 0b0000_0010
	Interrupt = 0b0000_0100
	Decimal   = 0b0000_1000
	Break     = 0b0001_0000
	Verflow   = 0b0100_0000
	Negative  = 0b1000_0000
)

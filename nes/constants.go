package nes

const (
	Implied = iota
	Immediate
	ZeroPage
	ZeroPageX
	Absolute
	AbsoluteX
	AbsoluteY
	IndirectX
	IndirectY
)

type Opcode struct {
	Code     uint8
	ByteSize int
	Cycles   int
	Mode     int
}

const (
	LDA_IMM = 0xA9
	LDA_ZER = 0xA5
	BRK_IMP = 0x00
	TAX_IMP = 0xAA
	INX_IMP = 0xE8
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

package nes

const (
	Implied = iota
	Immediate
	ZeroPage
	ZeroPageX
	ZeroPageY
	Absolute
	AbsoluteX
	AbsoluteY
	IndirectX
	IndirectY
)

const (
	LDA_IMM = 0xA9
	LDA_ZER = 0xA5
	LDA_ZRX = 0xB5
	LDA_ABS = 0xAD
	LDA_ABX = 0xBD
	LDA_ABY = 0xB9
	LDA_IDX = 0xA1
	LDA_IDY = 0xB1

	LDX_IMM = 0xA2
	LDX_ZER = 0xA6
	LDX_ZRY = 0xB6
	LDX_ABS = 0xAE
	LDX_ABY = 0xBE

	BRK_IMP = 0x00
	TAX_IMP = 0xAA
	INX_IMP = 0xE8
)

type Opcode struct {
	Code     uint8
	ByteSize int
	Cycles   int
	Mode     int
}

var Opcodes = map[uint8]Opcode{
	LDA_IMM: {Code: LDA_IMM, ByteSize: 2, Cycles: 2, Mode: Immediate},
	LDA_ZER: {Code: LDA_ZER, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	LDA_ZRX: {Code: LDA_ZRX, ByteSize: 2, Cycles: 4, Mode: ZeroPageX},
	LDA_ABS: {Code: LDA_ABS, ByteSize: 3, Cycles: 4, Mode: Absolute},
	LDA_ABX: {Code: LDA_ABX, ByteSize: 3, Cycles: 4 /*+1 crossed*/, Mode: AbsoluteX},
	LDA_ABY: {Code: LDA_ABY, ByteSize: 3, Cycles: 4 /*+1 crossed*/, Mode: AbsoluteY},
	LDA_IDX: {Code: LDA_IDX, ByteSize: 2, Cycles: 6, Mode: IndirectX},
	LDA_IDY: {Code: LDA_IDY, ByteSize: 2, Cycles: 5 /*+1 crossed*/, Mode: IndirectY},

	LDX_IMM: {Code: LDX_IMM, ByteSize: 2, Cycles: 2, Mode: Immediate},
	LDX_ZER: {Code: LDX_ZER, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	LDX_ZRY: {Code: LDX_ZRY, ByteSize: 2, Cycles: 4, Mode: ZeroPageY},
	LDX_ABS: {Code: LDX_ABS, ByteSize: 3, Cycles: 4, Mode: Absolute},
	LDX_ABY: {Code: LDX_ABY, ByteSize: 3, Cycles: 4 /*+1 crossed*/, Mode: AbsoluteY},

	BRK_IMP: {Code: BRK_IMP, ByteSize: 1, Cycles: 7, Mode: Implied},
}

const (
	Carry     = 0b0000_0001
	Zero      = 0b0000_0010
	Interrupt = 0b0000_0100
	Decimal   = 0b0000_1000
	Break     = 0b0001_0000
	Verflow   = 0b0100_0000
	Negative  = 0b1000_0000
)

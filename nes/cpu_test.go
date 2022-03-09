package nes

import (
	"testing"

	"github.com/zehlt/gesno/asrt"
)

func TestLdaImmediatePositiveValue(t *testing.T) {
	memory := Memory{
		LDA_IMM, 0x10, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x10))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestLdaImmediateZeroValue(t *testing.T) {
	memory := Memory{
		LDA_IMM, 0x00, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x00))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestLdaImmediateNegativeValue(t *testing.T) {
	memory := Memory{
		LDA_IMM, 0x80, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x80))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaZeroPagePositiveValue(t *testing.T) {
	memory := Memory{
		LDA_ZER, 0xAA, BRK_IMP,
	}
	memory[0xAA] = 0x67

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x67))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaZeroPageNegativeValue(t *testing.T) {
	memory := Memory{
		LDA_ZER, 0xAA, BRK_IMP,
	}
	memory[0xAA] = 0xDF

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0xDF))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaZeroPageZeroValue(t *testing.T) {
	memory := Memory{
		LDA_ZER, 0xAA, BRK_IMP,
	}
	memory[0xAA] = 0x00

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdaZeroPageXPositiveValue(t *testing.T) {
	memory := Memory{
		LDA_ZRX, 0x20, BRK_IMP,
	}
	memory[0x30] = 0x79

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x79))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaZeroPageXNegativeValue(t *testing.T) {
	memory := Memory{
		LDA_ZRX, 0x20, BRK_IMP,
	}
	memory[0x30] = 0xEF

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0xEF))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaZeroPageXZeroValue(t *testing.T) {
	memory := Memory{
		LDA_ZRX, 0x20, BRK_IMP,
	}
	memory[0x30] = 0x00

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdaZeroPageXWrappingValue(t *testing.T) {
	memory := Memory{
		LDA_ZRX, 0x80, BRK_IMP,
	}
	memory[0x7F] = 0x22

	cpu := Cpu{}
	cpu.XIndex = 0xFF
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x22))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaAbsolutePositiveValue(t *testing.T) {
	memory := Memory{
		LDA_ABS, 0xFE, 0x01, BRK_IMP,
	}
	memory[0x01fe] = 0x33

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x33))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaAbsoluteNegativeValue(t *testing.T) {
	memory := Memory{
		LDA_ABS, 0xFE, 0x01, BRK_IMP,
	}
	memory[0x01fe] = 0xAA

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0xAA))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaAbsoluteZeroValue(t *testing.T) {
	memory := Memory{
		LDA_ABS, 0xFE, 0x01, BRK_IMP,
	}
	memory[0x01fe] = 0x00

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdaAbsoluteXPositiveValueCrossedPage(t *testing.T) {
	memory := Memory{
		LDA_ABX, 0xFF, 0x01, BRK_IMP,
	}
	// 0x01ff+0x0010 = 0x020f
	memory[0x020f] = 0x66

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x66))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABX].Cycles+Opcodes[BRK_IMP].Cycles+1)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaAbsoluteXPositiveValueNotCrossedPage(t *testing.T) {
	memory := Memory{
		LDA_ABX, 0x00, 0x01, BRK_IMP,
	}
	memory[0x0110] = 0x77

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x77))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaAbsoluteXNegativeValue(t *testing.T) {
	memory := Memory{
		LDA_ABX, 0x20, 0x01, BRK_IMP,
	}
	memory[0x0130] = 0x90

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x90))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaAbsoluteXZeroValue(t *testing.T) {
	memory := Memory{
		LDA_ABX, 0x00, 0x01, BRK_IMP,
	}
	memory[0x0110] = 0x00

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdaAbsoluteYPositiveValueNotCrossed(t *testing.T) {
	memory := Memory{
		LDA_ABY, 0x10, 0x0A, BRK_IMP,
	}
	memory[0x0A30] = 0x45

	cpu := Cpu{}
	cpu.YIndex = 0x20
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x45))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaAbsoluteYNegativeValue(t *testing.T) {
	memory := Memory{
		LDA_ABY, 0x10, 0x0A, BRK_IMP,
	}
	memory[0x0A30] = 0xA5

	cpu := Cpu{}
	cpu.YIndex = 0x20
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0xA5))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABY].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaIndirectXPositiveValue(t *testing.T) {
	memory := Memory{
		LDA_IDX, 0x10, BRK_IMP,
	}
	memory[0x0015] = 0x07
	memory[0x0016] = 0x09

	memory[0x0907] = 0x79

	cpu := Cpu{}
	cpu.XIndex = 0x05
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x79))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_IDX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaIndirectXNegativeValue(t *testing.T) {
	memory := Memory{
		LDA_IDX, 0x10, BRK_IMP,
	}
	memory[0x0015] = 0x07
	memory[0x0016] = 0x09

	memory[0x0907] = 0xFF

	cpu := Cpu{}
	cpu.XIndex = 0x05
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0xFF))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_IDX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaIndirectXZeroValue(t *testing.T) {
	memory := Memory{
		LDA_IDX, 0x10, BRK_IMP,
	}
	memory[0x0015] = 0x07
	memory[0x0016] = 0x09

	memory[0x0907] = 0x00

	cpu := Cpu{}
	cpu.XIndex = 0x05
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_IDX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdaIndirectYPositiveValueNotCrossed(t *testing.T) {
	memory := Memory{
		LDA_IDY, 0x20, BRK_IMP,
	}
	memory[0x0020] = 0x03
	memory[0x0021] = 0x07

	memory[0x0704] = 0x55

	cpu := Cpu{}
	cpu.YIndex = 0x01
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x55))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_IDY].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaIndirectYPositiveValueCrossed(t *testing.T) {
	memory := Memory{
		LDA_IDY, 0x20, BRK_IMP,
	}
	memory[0x0020] = 0xFF
	memory[0x0021] = 0x07

	memory[0x080f] = 0x66

	cpu := Cpu{}
	cpu.YIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x66))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_IDY].Cycles+Opcodes[BRK_IMP].Cycles+1)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxImmediatePositiveValue(t *testing.T) {
	memory := Memory{
		LDX_IMM, 0x50, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x50))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxImmediateNegativeValue(t *testing.T) {
	memory := Memory{
		LDX_IMM, 0x81, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x81))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxImmediateZeroValue(t *testing.T) {
	memory := Memory{
		LDX_IMM, 0x00, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdxZeroPagePositiveValue(t *testing.T) {
	memory := Memory{
		LDX_ZER, 0x45, BRK_IMP,
	}

	memory[0x0045] = 0x22

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x22))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxZeroPageNegativeValue(t *testing.T) {
	memory := Memory{
		LDX_ZER, 0x45, BRK_IMP,
	}

	memory[0x0045] = 0xAE

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0xAE))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

// TODO: assure that wrap up doesn't cause any pb
func TestLdxZeroPageZeroValue(t *testing.T) {
	memory := Memory{
		LDX_ZER, 0x45, BRK_IMP,
	}

	memory[0x0045] = 0x00

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdxZeroPageYPositiveValue(t *testing.T) {
	memory := Memory{
		LDX_ZRY, 0x50, BRK_IMP,
	}

	memory[0x0060] = 0x25

	cpu := Cpu{}
	cpu.YIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x25))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ZRY].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxZeroPageYNegativeValue(t *testing.T) {
	memory := Memory{
		LDX_ZRY, 0x50, BRK_IMP,
	}

	memory[0x0060] = 0xDE

	cpu := Cpu{}
	cpu.YIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0xDE))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ZRY].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxAbsolutePositiveValue(t *testing.T) {
	memory := Memory{
		LDX_ABS, 0x50, 0x20, BRK_IMP,
	}

	memory[0x2050] = 0x10

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x10))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxAbsoluteNegativeValue(t *testing.T) {
	memory := Memory{
		LDX_ABS, 0x50, 0x20, BRK_IMP,
	}

	memory[0x2050] = 0xCE

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0xCE))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxAbsoluteZeroValue(t *testing.T) {
	memory := Memory{
		LDX_ABS, 0x50, 0x20, BRK_IMP,
	}

	memory[0x2050] = 0x00

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdxAbsoluteYPositiveValue(t *testing.T) {
	memory := Memory{
		LDX_ABY, 0x50, 0x20, BRK_IMP,
	}

	memory[0x2061] = 0x35

	cpu := Cpu{}
	cpu.YIndex = 0x11
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x35))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxAbsoluteYPositiveValueCrossedPage(t *testing.T) {
	memory := Memory{
		LDX_ABY, 0xFF, 0x20, BRK_IMP,
	}

	memory[0x210F] = 0x35

	cpu := Cpu{}
	cpu.YIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x35))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ABS].Cycles+Opcodes[BRK_IMP].Cycles+1)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxAbsoluteYPositiveValueCrossedPageAndNegative(t *testing.T) {
	memory := Memory{
		LDX_ABY, 0xFF, 0x20, BRK_IMP,
	}

	memory[0x210F] = 0xCC

	cpu := Cpu{}
	cpu.YIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0xCC))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ABS].Cycles+Opcodes[BRK_IMP].Cycles+1)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

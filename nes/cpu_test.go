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

/*
func TestLdaImmediateZeroFlag(t *testing.T) {
	memory := Memory{
		LDA_IMM, 0x00, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(memory)

	//assertEqual(t, int(cpu.Accumulator), 0x00)
}

// TODO: NEGATIVE
func TestLdaImmediateNegative(t *testing.T) {
	memory := Memory{
		LDA_IMM, 0x00, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(memory)

	//assertEqual(t, int(cpu.Accumulator), 0x00)
}
*/
//func Test

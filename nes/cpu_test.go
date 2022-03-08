package nes

import (
	"testing"

	"github.com/zehlt/gesno/asrt"
)

func TestLdaImmediateNormalValue(t *testing.T) {
	memory := Memory{
		LDA_IMM, 0x10, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x10))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
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
}

func TestLdaImmediateNegativeValue(t *testing.T) {
	memory := Memory{
		LDA_IMM, 0x80, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x80))
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
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

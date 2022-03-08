package nes

import (
	"testing"
)

func TestLdaImmediateNormalData(t *testing.T) {
	/*
		memory := Memory{
			LDA_IMM, 0x10, BRK_IMP,
		}
		cpu := Cpu{}

		cpu.Run(memory)
	*/

	//assertEqual(t, int(cpu.Accumulator), 0x10)
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

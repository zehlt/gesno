package nes

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Fatalf("Received: %v (%v), expected: %v (%v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

func assert(t *testing.T, tst bool) {
	if tst {
		return
	}
	t.Fatalf("Error!")
}

func TestLdaImmediateNormalData(t *testing.T) {
	memory := Memory{
		LDA_IMM, 0x10, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(memory)

	assertEqual(t, int(cpu.Accumulator), 0x10)
	assert(t, !cpu.Status.Has(Zero))
	assert(t, !cpu.Status.Has(Negative))
}

func TestLdaImmediateZeroFlag(t *testing.T) {
	memory := Memory{
		LDA_IMM, 0x00, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(memory)

	assertEqual(t, int(cpu.Accumulator), 0x00)
	assert(t, cpu.Status.Has(Zero))
	assert(t, !cpu.Status.Has(Negative))
}

// TODO: NEGATIVE
func TestLdaImmediateNegative(t *testing.T) {
	memory := Memory{
		LDA_IMM, 0x00, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(memory)

	assertEqual(t, int(cpu.Accumulator), 0x00)
	assert(t, cpu.Status.Has(Zero))
	assert(t, !cpu.Status.Has(Negative))
}

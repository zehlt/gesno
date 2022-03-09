package nes

type Cpu struct {
	Cycle int
	Registers
}

var Opcodes = map[uint8]Opcode{
	LDA_IMM: {Code: LDA_IMM, ByteSize: 2, Cycles: 2, Mode: Immediate},
	LDA_ZER: {Code: LDA_ZER, ByteSize: 2, Cycles: 3, Mode: ZeroPage},

	BRK_IMP: {Code: BRK_IMP, ByteSize: 1, Cycles: 7, Mode: Implied},
}

/*
func (c *Cpu) getNextByte(memory Memory) uint8 {
	opcode := memory.readByte(uint16(c.ProgramCounter))
	c.ProgramCounter++
	return opcode
}
*/

/*
func (c *Cpu) ldaImm(memory Memory) {
	operand := c.getNextByte(memory)
	c.Accumulator = Register8(operand)

	c.updateZeroAndNegativeFlags(c.Accumulator)

}

func (c *Cpu) taxImp() {
	c.XIndex = c.Accumulator

	c.updateZeroAndNegativeFlags(c.XIndex)
}

func (c *Cpu) inxImp() {
	c.XIndex += 1

	c.updateZeroAndNegativeFlags(c.XIndex)
}
*/

func (c *Cpu) updateZeroAndNegativeFlags(value Register8) {
	if value == 0 {
		c.Status.Add(Zero)
	} else {
		c.Status.Remove(Zero)
	}

	if value.IsNegative() {
		c.Status.Add(Negative)
	} else {
		c.Status.Remove(Negative)
	}
}

func (c *Cpu) getOperandAddress(mem *Memory, mode int) uint16 {
	switch mode {
	case Implied:
		panic("Implied mode not supported")
	case Immediate:
		return uint16(c.ProgramCounter)
	case ZeroPage:
		return uint16(mem.readByte(uint16(c.ProgramCounter)))
	default:
		panic("Addressing mode not implemented")
	}
}

func (c *Cpu) lda(mem *Memory, mode int) {
	operand := mem.readByte(c.getOperandAddress(mem, mode))
	c.Accumulator = Register8(operand)
	c.updateZeroAndNegativeFlags(Register8(operand))
}

func (c *Cpu) Reset(mem *Memory) {
	c.Accumulator = 0
	c.XIndex = 0
	c.YIndex = 0
	c.StackPointer = 0
	c.Status = 0

	c.ProgramCounter = Register16(mem.readWord(0xFFFC))
}

func (c *Cpu) interpret(opcode uint8, memory *Memory) bool {
	opc := Opcodes[opcode]

	switch opcode {
	case LDA_IMM, LDA_ZER:
		c.lda(memory, opc.Mode)
		c.Cycle += opc.Cycles
		c.ProgramCounter += Register16(opc.ByteSize - 1)
	case TAX_IMP:
		// c.taxImp()
	case INX_IMP:
		// c.inxImp()
	case BRK_IMP:
		c.Cycle += 7
		c.ProgramCounter += Register16(opc.ByteSize - 1)
		return true
	default:
		panic("Unknown opcode!")
	}

	return false
}

func (c *Cpu) Run(memory *Memory) {
	for {
		opcode := memory.readByte(uint16(c.ProgramCounter))
		c.ProgramCounter++

		if c.interpret(opcode, memory) {
			break
		}
	}
}

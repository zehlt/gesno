package nes

type Cpu struct {
	Cycle int
	Registers
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

func (c *Cpu) interpret(opcode uint8, memory *Memory) bool {
	switch opcode {
	case LDA_IMM:
		c.ldaImm(memory)
	case LDA_ZER:
		c.ldaZer(memory)
	case TAX_IMP:
		// c.taxImp()
	case INX_IMP:
		// c.inxImp()
	case BRK_IMP:
		c.Cycle += 7
		return true
	default:
		panic("Unknown opcode!")
	}

	return false
}

func (c *Cpu) ldaImm(mem *Memory) {
	// lda immediate
	data := mem.readByte(uint16(c.ProgramCounter))
	c.ProgramCounter++
	c.Accumulator = Register8(data)

	c.updateZeroAndNegativeFlags(Register8(data))
	c.Cycle += 2
}

func (c *Cpu) ldaZer(mem *Memory) {

}

func (c *Cpu) Reset(mem *Memory) {
	c.Accumulator = 0
	c.XIndex = 0
	c.YIndex = 0
	c.StackPointer = 0
	c.Status = 0

	c.ProgramCounter = Register16(mem.readWord(0xFFFC))
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

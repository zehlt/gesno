package nes

type Cpu struct {
	Registers
}

func (c *Cpu) getNextByte(memory Memory) uint8 {
	opcode := memory[c.ProgramCounter]
	c.ProgramCounter++
	return opcode
}

func (c *Cpu) ldaImmediate(memory Memory) {
	operand := c.getNextByte(memory)
	c.Accumulator = Register8(operand)

	c.updateZeroAndNegativeFlags(c.Accumulator)

}

func (c *Cpu) taxImplied() {
	c.XIndex = c.Accumulator

	c.updateZeroAndNegativeFlags(c.XIndex)
}

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

func (c *Cpu) interpret(opcode uint8, memory Memory) bool {
	switch opcode {
	case LDA_IMM:
		c.ldaImmediate(memory)
	case TAX_IMP:
		c.taxImplied()
	case BRK_IMP:
		return true
	default:
		panic("Unknown opcode!")
	}

	return false
}

func (c *Cpu) Run(memory Memory) {
	c.Accumulator = 0

	for {
		opcode := c.getNextByte(memory)
		if c.interpret(opcode, memory) {
			break
		}
	}
}

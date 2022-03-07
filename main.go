package main

import (
	"fmt"

	"github.com/zehlt/gesno/nes"
)

func main() {
	memory := nes.Memory{
		0xa9, 0x10, 0x00,
	}

	cpu := nes.Cpu{}
	cpu.Run(memory)
	fmt.Printf("status: %08b\n", cpu.Status)
}

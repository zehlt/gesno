package nes

import "github.com/zehlt/go6502"

type Nes struct {
	Cpu go6502.Cpu
	Mem go6502.Mem
	Bus NesBus
}

func (n *Nes) Init() {
	n.Bus.m = &n.Mem
}

func (n *Nes) Load() {

}

func (n *Nes) Start() {
}

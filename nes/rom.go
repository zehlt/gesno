package nes

const (
	MIRRORING_VERTICAL = iota
	MIRRORING_HORIZONTAL
	MIRRORING_FOUR_SCREEN
)

type Rom struct {
	Prg    []uint8
	Chr    []uint8
	Mapper uint8
	Mirror int
}

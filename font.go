package main

import (
	_ "embed"

	"github.com/zehlt/gelly"
)

//go:embed robmed.ttf
var robmed []byte

var FONT_ROBMED = gelly.FontAsset{
	Format: gelly.TTF,
	Data:   robmed,
	Path:   "robmed.ttf",
}

package main

import (
	"github.com/zehlt/gelly"
	"github.com/zehlt/gelly/cmd"
)

func main() {
	client := cmd.NewEbitenClient(&GesnoEmul{}, gelly.Configuration{
		Width:  1280,
		Height: 720,
		Title:  "Gesno - Go NES Emulator",
	})
	client.Run()
}

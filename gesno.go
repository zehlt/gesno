package main

import (
	"fmt"

	"github.com/zehlt/gesno/go65"
)

func main() {
	fmt.Println("Print from main")
	go65.PrintCpuName()
}

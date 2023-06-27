package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-isatty"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Mode for Stdin: %v\n", fi.Mode())
	fmt.Printf("Stdin is a character device: %v\n", fi.Mode()&os.ModeCharDevice != 0)
	fmt.Printf("Stdin is a named pipe: %v\n", fi.Mode()&os.ModeNamedPipe != 0)
	if isatty.IsTerminal(os.Stdin.Fd()) {
		fmt.Println("Stdout is a terminal")
	}
}

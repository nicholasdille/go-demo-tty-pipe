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
	fmt.Printf("Mode for Stdin: %032b\n", fi.Mode())
	fmt.Printf("                %032b\n", os.ModeCharDevice)
	fmt.Printf("                %032b\n", os.ModeNamedPipe)
	if fi.Mode()&os.ModeCharDevice == os.ModeCharDevice {
		fmt.Println("Stdin is a character device")
	}
	if fi.Mode()&os.ModeNamedPipe == os.ModeNamedPipe {
		fmt.Println("Stdin is a named pipe")
	}
	if isatty.IsTerminal(os.Stdin.Fd()) {
		fmt.Println("Stdin is a terminal")
	}

	fo, err := os.Stdout.Stat()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Mode for Stdout: %032b\n", fo.Mode())
	fmt.Printf("                %032b\n", os.ModeCharDevice)
	fmt.Printf("                %032b\n", os.ModeNamedPipe)
	if fo.Mode()&os.ModeCharDevice == os.ModeCharDevice {
		fmt.Println("Stdout is a character device")
	}
	if fo.Mode()&os.ModeNamedPipe == os.ModeNamedPipe {
		fmt.Println("Stdout is a named pipe")
	}
	if isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Println("Stdout is a terminal")
	}
}

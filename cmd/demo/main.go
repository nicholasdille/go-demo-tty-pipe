package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-isatty"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}
	fmt.Fprintf(os.Stderr, "Mode for Stdin: %032b\n", fi.Mode())
	fmt.Fprintf(os.Stderr, "                %032b\n", os.ModeCharDevice)
	fmt.Fprintf(os.Stderr, "                %032b\n", os.ModeNamedPipe)
	if fi.Mode()&os.ModeCharDevice == os.ModeCharDevice {
		fmt.Fprintf(os.Stderr, "Stdin is a character device")
	} else {
		fmt.Fprintf(os.Stderr, "Stdin is not a character device")
	}
	if fi.Mode()&os.ModeNamedPipe == os.ModeNamedPipe {
		fmt.Fprintf(os.Stderr, "Stdin is a named pipe")
	} else {
		fmt.Fprintf(os.Stderr, "Stdin is not a named pipe")
	}
	if isatty.IsTerminal(os.Stdin.Fd()) {
		fmt.Fprintf(os.Stderr, "Stdin is a terminal")
	} else {
		fmt.Fprintf(os.Stderr, "Stdin is not a terminal")
	}

	fo, err := os.Stdout.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}
	fmt.Fprintf(os.Stderr, "Mode for Stdout: %032b\n", fo.Mode())
	fmt.Fprintf(os.Stderr, "                %032b\n", os.ModeCharDevice)
	fmt.Fprintf(os.Stderr, "                %032b\n", os.ModeNamedPipe)
	if fo.Mode()&os.ModeCharDevice == os.ModeCharDevice {
		fmt.Fprintf(os.Stderr, "Stdout is a character device")
	} else {
		fmt.Fprintf(os.Stderr, "Stdout is not a character device")
	}
	if fo.Mode()&os.ModeNamedPipe == os.ModeNamedPipe {
		fmt.Fprintf(os.Stderr, "Stdout is a named pipe")
	} else {
		fmt.Fprintf(os.Stderr, "Stdout is not a named pipe")
	}
	if isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Fprintf(os.Stderr, "Stdout is a terminal")
	} else {
		fmt.Fprintf(os.Stderr, "Stdout is not a terminal")
	}
}

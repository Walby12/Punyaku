package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func check_extension(file_path string) {
	ext := filepath.Ext(file_path)

	if ext != ".pun" {
		fmt.Println("Error: file extension must be .pun")
		os.Exit(-1)
	}
}
func usage() int {
	fmt.Println("USAGE: Punyaku <SUBCOMMAND> [ARGS]")
	fmt.Println("SUBCOMMANDS:")
	fmt.Println("    sim <file_name>    Simulate the program")
	fmt.Println("    com <file_name>    Compile the program")
	os.Exit(1)
	return 0
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: insufficient arguments")
		usage()
	}

	subcommand := os.Args[1]
	filePath := os.Args[2]

	check_extension(filePath)

	program := parse_program(filePath)

	switch subcommand {
	case "sim":
		sim_prog(program)
	case "com":
		asmFilePath := strings.TrimSuffix(filePath, filepath.Ext(filePath)) + ".asm"
		gen_asm_file(asmFilePath, program)
	default:
		fmt.Println("Error: unrecognized subcommand:", subcommand)
		usage()
	}
}

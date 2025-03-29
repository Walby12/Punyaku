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
	fmt.Println("    sim [optional flags] <file_name>    Simulate the program")
	fmt.Println("    com [optional flags] <file_name>    Compile the program")
	os.Exit(1)
	return 0
}

func main() {
	if len(os.Args) < 2 || len(os.Args) == 1 {
		fmt.Println("Error: insufficient arguments")
		usage()
	}

	runAfterCompile := false
	argIndex := 1
	if os.Args[1] == "-r" {
		runAfterCompile = true
		argIndex++
	}
	if os.Args[1] == "help" {
		usage()
		os.Exit(0)
	}
	if os.Args[1] == "-h" {
		usage()
		os.Exit(0)
	}

	if len(os.Args) <= argIndex {
		fmt.Println("Error: missing file name")
		usage()
	}

	subcommand := os.Args[argIndex]
	filePath := os.Args[argIndex+1]

	check_extension(filePath)

	program := parse_program(filePath)

	switch subcommand {
	case "sim":
		sim_prog(program)
	case "com":
		asmFilePath := strings.TrimSuffix(filePath, filepath.Ext(filePath)) + ".asm"
		gen_asm_file(asmFilePath, program)
		executeFilePath := strings.TrimSuffix(asmFilePath, filepath.Ext(asmFilePath)) + ""
		if runAfterCompile {
			fmt.Println(" ")
			fmt.Println("--- !!! Running file !!! ---")
			fmt.Println(" ")
			run_asm_file(executeFilePath)
		}
	default:
		fmt.Println("Error: unrecognized subcommand:", subcommand)
		usage()
	}
}

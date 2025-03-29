package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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
		fmt.Println("USAGE: Punyaku <SUBCOMMAND> <file_name>")
		os.Exit(1)
	}

	subcommand := os.Args[1]
	filePath := os.Args[2]

	check_extension(filePath)

	// Parse the program from the .pun file
	program := parse_program(filePath)

	switch subcommand {
	case "sim":
		sim_prog(program)
	case "com":
		asmFilePath := strings.TrimSuffix(filePath, filepath.Ext(filePath)) + ".asm"
		gen_asm_file(asmFilePath, program)
	default:
		fmt.Println("Error: unrecognized subcommand:", subcommand)
		os.Exit(1)
	}
}

type Instruction struct {
	op_code int
	value   any
}

const (
	OP_PUSH = iota
	OP_PLUS
	OP_MINUS
	OP_DUMP
	COUNT_OPS
	COUNT
)

func push(x any) Instruction {
	return Instruction{OP_PUSH, x}
}

func plus() Instruction {
	return Instruction{OP_PLUS, nil}
}

func dump() Instruction {
	return Instruction{OP_DUMP, nil}
}

func minus() Instruction {
	return Instruction{OP_MINUS, nil}
}

func parse_program(file_path string) []Instruction {
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(-1)
	}
	defer file.Close()

	var program []Instruction
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, ";;") { // Skip empty lines or comments
			continue
		}

		// Split the line into tokens
		tokens := strings.Fields(line)
		i := 0
		for i < len(tokens) {
			op := strings.ToUpper(tokens[i])
			switch op {
			case "=": // PUSH instruction
				if i+1 >= len(tokens) {
					fmt.Println("Error: PUSH requires a value")
					os.Exit(-1)
				}
				value, err := strconv.Atoi(tokens[i+1])
				if err != nil {
					fmt.Println("Error: invalid value for PUSH:", tokens[i+1])
					os.Exit(-1)
				}
				program = append(program, push(value))
				i += 2 // Move to the next instruction
			case "+": // PLUS instruction
				program = append(program, plus())
				i++
			case "-": // MINUS instruction
				program = append(program, minus())
				i++
			case ".": // DUMP instruction
				program = append(program, dump())
				i++
			default:
				fmt.Println("Error: unrecognized instruction:", op)
				os.Exit(-1)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(-1)
	}

	return program
}

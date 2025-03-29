package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	op_code int
	value   any
}

const (
	OP_PUSH = iota
	OP_PLUS
	OP_MINUS
	OP_DUMP
	OP_EQUALS
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

func equals() Instruction {
	return Instruction{OP_EQUALS, nil}
}

func parse_program(file_path string) []Instruction {
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(-1)
	}
	defer file.Close()

	var program []Instruction
	var stack []int
	scanner := bufio.NewScanner(file)

	num_line := 1

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, ";;") {
			continue
		}

		tokens := strings.Fields(line)
		i := 0
		for i < len(tokens) {
			op := tokens[i]
			switch op {
			case "int":
				if i+1 >= len(tokens) {
					fmt.Println("Error: PUSH requires a value", "on line:", num_line, "at index:", i)
					os.Exit(-1)
				}
				value, err := strconv.Atoi(tokens[i+1])
				if err != nil {
					fmt.Println("Error: invalid value for PUSH:", tokens[i+1], "on line:", num_line, "at index:", i)
					os.Exit(-1)
				}
				stack = append(stack, value)
				program = append(program, push(value))
				i += 2
			case "+":
				if len(stack) < 2 {
					fmt.Println("Error: not enough values on the stack for PLUS", "on line:", num_line, "at index:", i)
					os.Exit(-1)
				}
				a := stack[len(stack)-1]
				b := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, a+b)
				program = append(program, plus())
				i++
			case "-":
				if len(stack) < 2 {
					fmt.Println("Error: not enough values on the stack for MINUS", "on line:", num_line, "at index:", i)
					os.Exit(-1)
				}
				a := stack[len(stack)-1]
				b := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, b-a)
				program = append(program, minus())
				i++
			case ".":
				if len(stack) == 0 {
					fmt.Println("Error: stack is empty, cannot DUMP", "on line:", num_line, "at index:", i)
					os.Exit(-1)
				}
				stack = stack[:len(stack)-1]
				program = append(program, dump())
				i++
			case "=":
				if len(stack) < 2 {
					fmt.Println("Error: not enough values on the stack for EQUALS", "on line:", num_line, "at index:", i)
					os.Exit(-1)
				}
				a := stack[len(stack)-1]
				b := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				if a == b {
					stack = append(stack, 1)
				} else {
					stack = append(stack, 0)
				}
				program = append(program, equals())
				i++
			default:
				fmt.Println("Error: unrecognized instruction:", op, "with start at index:", i, "on line:", num_line)
				os.Exit(-1)
			}
		}
		num_line++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(-1)
	}

	return program
}

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
	line    int
}

const (
	OP_PUSH = iota
	OP_PLUS
	OP_MINUS
	OP_DUMP
	OP_EQUALS
	OP_IF
	OP_END
	COUNT_OPS
	COUNT
)

func push(x any) Instruction {
	return Instruction{OP_PUSH, x, -1}
}

func plus() Instruction {
	return Instruction{OP_PLUS, nil, -1}
}

func dump() Instruction {
	return Instruction{OP_DUMP, nil, -1}
}

func minus() Instruction {
	return Instruction{OP_MINUS, nil, -1}
}

func equals() Instruction {
	return Instruction{OP_EQUALS, nil, -1}
}

func if_op() Instruction {
	return Instruction{OP_IF, nil, -1}
}

func end_op() Instruction {
	return Instruction{OP_END, nil, -1}
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
			num_line++
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
					fmt.Println("Error: invalid int value for PUSH:", tokens[i+1], "on line:", num_line, "at index:", i)
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
			case "if":
				program = append(program, if_op())
				i++
			case "end":
				program = append(program, end_op())
				i++
			default:
				fmt.Println("Error: unrecognized instruction:", op, "on line:", num_line, "at index:", i)
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

func validate_if_end_balance(program []Instruction) {
	balance := 0

	for _, instr := range program {
		switch instr.op_code {
		case OP_IF:
			balance++
		case OP_END:
			balance--
			if balance < 0 {
				fmt.Printf("Error: Unmatched 'end' at instruction")
				os.Exit(1)
			}
		}
	}

	if balance != 0 {
		for _, instr := range program {
			if instr.op_code == OP_IF {
				fmt.Printf("Error: Unmatched 'if' at instruction")
				os.Exit(1)
			}
		}
	}
}

package main

import (
	"fmt"
	"os"
)

func usage() int {
	fmt.Println("USAGE: Punyaku <SUBCOMMAND> [ARGS]")
	fmt.Println("SUBCOMMANDS:")
	fmt.Println("    sim     Simulate the program")
	fmt.Println("    com     Compile the program")
	os.Exit(1)
	return 0
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage())
	}
	program := []Instruction{push(35), push(34), plus(), dump(), push(500), push(80), minus(), dump()}
	subcommand := os.Args[1]
	switch subcommand {
	case "sim":
		sim_prog(program)
		os.Exit(1)
	case "com":
		compile_assembly(os.Args[2], program)
	default:
		fmt.Println("Error: unrecognized sub command:", os.Args[1])
	}
}

func sim_prog(program []Instruction) {
	stack := []Instruction{}

	for _, instr := range program {
		switch instr.op_code {
		case OP_PUSH:
			stack = append(stack, instr)

		case OP_PLUS:
			if len(stack) < 2 {
				fmt.Println("Error: Not enough values for PLUS operation")
				os.Exit(1)
			}

			val1 := stack[len(stack)-1]
			val2 := stack[len(stack)-2]

			val1_int, ok1 := val1.value.(int)
			val2_int, ok2 := val2.value.(int)

			if !ok1 || !ok2 {
				fmt.Println("Error: PLUS operation requires two integers")
				os.Exit(1)
			}

			stack = stack[:len(stack)-2]

			stack = append(stack, Instruction{op_code: OP_PUSH, value: val1_int + val2_int})

		case OP_MINUS:
			if len(stack) < 2 {
				fmt.Println("Error: Not enough values for PLUS operation")
				os.Exit(1)
			}

			val1 := stack[len(stack)-1]
			val2 := stack[len(stack)-2]

			val1_int, ok1 := val1.value.(int)
			val2_int, ok2 := val2.value.(int)

			if !ok1 || !ok2 {
				fmt.Println("Error: MINUS operation requires two integers")
				os.Exit(1)
			}

			stack = stack[:len(stack)-2]

			stack = append(stack, Instruction{op_code: OP_PUSH, value: -val1_int + val2_int})

		case OP_DUMP:
			if len(stack) == 0 {
				fmt.Println("Error: Stack is empty")
				os.Exit(1)
			}

			top := stack[len(stack)-1]
			fmt.Println(top.value)

		default:
			fmt.Println("Unknown instruction:", instr.op_code)
			os.Exit(1)
		}
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

package main

import (
	"fmt"
	"os"
)

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

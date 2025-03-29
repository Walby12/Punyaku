package main

import (
	"fmt"
	"os"
)

func sim_prog(program []Instruction) {
	validate_if_end_balance(program)

	stack := []Instruction{}

	for i := 0; i < len(program); i++ {
		instr := program[i]

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
				fmt.Println("Error: Not enough values for MINUS operation")
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

			stack = append(stack, Instruction{op_code: OP_PUSH, value: val2_int - val1_int})

		case OP_DUMP:
			if len(stack) == 0 {
				fmt.Println("Error: Stack is empty")
				os.Exit(1)
			}

			top := stack[len(stack)-1]
			fmt.Println(top.value)

		case OP_EQUALS:
			if len(stack) < 2 {
				fmt.Println("Error: Not enough values for EQUALS operation")
				os.Exit(1)
			}

			val1 := stack[len(stack)-1]
			val2 := stack[len(stack)-2]

			if val1.value == val2.value {
				stack = append(stack, Instruction{op_code: OP_PUSH, value: 1})
			} else {
				stack = append(stack, Instruction{op_code: OP_PUSH, value: 0})
			}

		case OP_IF:
			if len(stack) == 0 {
				fmt.Println("Error: stack is empty, cannot evaluate IF condition")
				os.Exit(1)
			}

			condition := stack[len(stack)-1].value
			stack = stack[:len(stack)-1]

			conditionInt, ok := condition.(int)
			if !ok {
				fmt.Println("Error: IF condition must be an integer (0 or 1)")
				os.Exit(1)
			}

			if conditionInt == 0 {
				_, endIndex := parse_until_end(program, i)
				i = endIndex
			}
		case OP_END:
			continue

		default:
			fmt.Println("Unknown instruction:", instr.op_code)
			os.Exit(1)
		}
	}
}

func parse_until_end(program []Instruction, startIndex int) (int, int) {
	if startIndex < 0 || startIndex >= len(program) {
		fmt.Println("Error: Invalid start index")
		os.Exit(1)
	}

	nestedIfCount := 0

	for i := startIndex + 1; i < len(program); i++ {
		switch program[i].op_code {
		case OP_IF:

			nestedIfCount++
		case OP_END:
			if nestedIfCount == 0 {

				return startIndex, i
			}

			nestedIfCount--
		}
	}

	fmt.Println("Error: No matching 'end' found for 'if' starting at index", startIndex)
	os.Exit(1)
	return -1, -1
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func gen_code() {
	start := `
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}
	`
	err := os.WriteFile("out.go", []byte(start), 0644)
	if err != nil {
		fmt.Println("Failed to write the code:", err)
		os.Exit(-1)
	}
	cmd := exec.Command("go", "build", "out.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Compilation failed,", err)
		os.Exit(-1)
	} else {
		fmt.Println("Compilation phase succesful, exec created: out")
	}
	time.Sleep(2 * time.Second)
	cmd = exec.Command("cmd", "/C", "del", "out.go")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Clean up failed,", err)
		os.Exit(-1)
	} else {
		fmt.Println("Clean up phase succesful")
	}
	//stack := []Instruction{}
	/*
		for _, code := range code_to_gen {
			switch code.op_code {
			case OP_PUSH:

			}
		}
	*/
}

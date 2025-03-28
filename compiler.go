package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func compile_assembly(file_path string) {
	// Generate the .asm file
	gen_asm_file(file_path)

	// Use filepath.Ext to get the file extension
	path_ext := filepath.Ext(file_path)

	// Use strings.TrimSuffix to remove the extension and add ".exe"
	new_file_path_ext := strings.TrimSuffix(file_path, path_ext) + ".exe"

	// Call the Makefile to compile the assembly code
	cmd := exec.Command("make", fmt.Sprintf("ASM_FILE=%s", file_path), fmt.Sprintf("EXE_FILE=%s", new_file_path_ext))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Compilation failed,", err)
		os.Exit(-1)
	} else {
		fmt.Println("Compilation phase successful, file created:", new_file_path_ext)
	}
}

func gen_code() string {
	// This function would generate the assembly code and write it to the file

	start_asm_code := `
%define SYS_EXIT 60

segment .text
global main
main:

`
	var asmCodeBuilder strings.Builder
	asmCodeBuilder.WriteString(start_asm_code)

	asmCodeBuilder.WriteString("    mov rax, SYS_EXIT\n")
	asmCodeBuilder.WriteString("    mov rdi, 42\n")
	asmCodeBuilder.WriteString("    syscall\n")

	return asmCodeBuilder.String()
}

func gen_asm_file(out_file_path string) {
	// Generate the assembly code as a string
	asm_code := gen_code()

	// Write the assembly code to the specified file
	err := os.WriteFile(out_file_path, []byte(asm_code), 0644)
	if err != nil {
		fmt.Println("Error writing assembly file:", err)
		os.Exit(-1)
	}

	fmt.Println("Assembly file created:", out_file_path)
}

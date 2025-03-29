package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func gen_code(program []Instruction) string {
	// Start the assembly code with the necessary directives
	start_asm_code := `%define SYS_EXIT 60

section .text
global _start

dump:
    push    rbp
    mov     rbp, rsp
    sub     rsp, 64
    mov     DWORD  [rbp-52], edi
    mov     QWORD  [rbp-8], 1
    mov     eax, 32
    sub     rax, QWORD  [rbp-8]
    mov     BYTE  [rbp-48+rax], 10
.L2:
    mov     ecx, DWORD  [rbp-52]
    mov     edx, ecx
    mov     eax, 3435973837
    imul    rax, rdx
    shr     rax, 32
    mov     edx, eax
    shr     edx, 3
    mov     eax, edx
    sal     eax, 2
	add     eax, edx
    add     eax, eax
    sub     ecx, eax
    mov     edx, ecx
    mov     eax, edx
    lea     edx, [rax+48]
    mov     eax, 31
    sub     rax, QWORD  [rbp-8]
    mov     BYTE  [rbp-48+rax], dl
    mov     eax, DWORD  [rbp-52]
    mov     edx, eax
    mov     eax, 3435973837
    imul    rax, rdx
    shr     rax, 32
    shr     eax, 3
    mov     DWORD  [rbp-52], eax
    add     QWORD  [rbp-8], 1
    cmp     DWORD  [rbp-52], 0
    jne     .L2
    mov     eax, 32
    sub     rax, QWORD  [rbp-8]
    lea     rdx, [rbp-48]
    lea     rcx, [rdx+rax]
    mov     rax, QWORD  [rbp-8]
    mov     rdx, rax
    mov     rsi, rcx
    mov     edi, 1
    mov     eax, 0
	mov     rax, 1
    syscall
    nop
    leave
    ret

_start:
`
	var asmCodeBuilder strings.Builder
	asmCodeBuilder.WriteString(start_asm_code)

	for _, instruction := range program {
		// Convert each instruction to its assembly equivalent
		switch instruction.op_code {
		case OP_PUSH:
			asmCodeBuilder.WriteString(fmt.Sprintf("    ;; -- push %d --\n", instruction.value))
			asmCodeBuilder.WriteString(fmt.Sprintf("    push %d\n", instruction.value))
		case OP_PLUS:
			asmCodeBuilder.WriteString("    ;; -- plus --\n")
			asmCodeBuilder.WriteString("    pop rbx\n")
			asmCodeBuilder.WriteString("    pop rax\n")
			asmCodeBuilder.WriteString("    add rax, rbx\n")
			asmCodeBuilder.WriteString("    push rax\n")
		case OP_MINUS:
			asmCodeBuilder.WriteString("    ;; -- minus --\n")
			asmCodeBuilder.WriteString("    pop rbx\n")
			asmCodeBuilder.WriteString("    pop rax\n")
			asmCodeBuilder.WriteString("    sub rax, rbx\n")
			asmCodeBuilder.WriteString("    push rax\n")
		case OP_DUMP:
			asmCodeBuilder.WriteString("    ;; -- dump --\n")
			asmCodeBuilder.WriteString("    pop rdi\n")
			asmCodeBuilder.WriteString("    call dump\n")
		default:
			fmt.Println("Error: unrecognized operation code:", instruction.op_code)
			os.Exit(-1)
		}
	}

	asmCodeBuilder.WriteString("    ;; -- exit --\n")
	asmCodeBuilder.WriteString("    mov rax, SYS_EXIT\n")
	asmCodeBuilder.WriteString("    mov rdi, 0\n")
	asmCodeBuilder.WriteString("    syscall\n")

	return asmCodeBuilder.String()
}

func gen_asm_file(out_file_path string, program []Instruction) {
	// Generate the assembly code as a string
	asm_code := gen_code(program)

	// Create or overwrite the specified file
	file, err := os.Create(out_file_path)
	if err != nil {
		fmt.Println("Error creating assembly file:", err)
		os.Exit(-1)
	}
	defer file.Close() // Ensure the file is closed after writing

	// Write the assembly code to the file
	_, err = file.WriteString(asm_code)
	if err != nil {
		fmt.Println("Error writing to assembly file:", err)
		os.Exit(-1)
	}

	fmt.Println("Assembly file created:", out_file_path)
	// Use filepath.Ext to get the file extension
	path_ext := filepath.Ext(out_file_path)

	// Use strings.TrimSuffix to remove the extension and add ".exe"
	new_file_path_ext := strings.TrimSuffix(out_file_path, path_ext) + ""

	// Call the Makefile to compile the assembly code
	cmd := exec.Command("make", fmt.Sprintf("ASM_FILE=%s", out_file_path), fmt.Sprintf("EXE_FILE=%s", new_file_path_ext))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Compilation failed,", err)
		os.Exit(-1)
	} else {
		fmt.Println("Compilation phase successful, file created:", new_file_path_ext)
	}
}

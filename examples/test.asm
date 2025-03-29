%define SYS_EXIT 60

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
    ;; -- push 34 --
    push 34
    ;; -- push 35 --
    push 35
    ;; -- plus --
    pop rbx
    pop rax
    add rax, rbx
    push rax
    ;; -- dump --
    pop rdi
    call dump
    ;; -- push 32 --
    push 32
    ;; -- push 21 --
    push 21
    ;; -- minus --
    pop rbx
    pop rax
    sub rax, rbx
    push rax
    ;; -- dump --
    pop rdi
    call dump
    ;; -- exit --
    mov rax, SYS_EXIT
    mov rdi, 0
    syscall

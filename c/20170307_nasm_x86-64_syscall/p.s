; 参考地址：http://blog.csdn.net/mydo/article/details/45007805
section .text

; if use ld
global _start
_start:
; if use gcc
; global main
; main:

	mov rax, 1 ; write NO
	mov rdi, 1 ; fd
	mov rsi, msg ; addr of msg string
	mov rdx, msg_len ; length of msg string
	syscall

	mov rax, 60 ; exit NO
	mov rdi, 0 ; error code
	syscall

	msg: db "hello world\n"
	msg_len: equ $-msg

; 编译、链接和运行
; nasm -f elf64 p.s
; ld -o p p.o
; ./p


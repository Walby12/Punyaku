# Compiler and assembler
AS = nasm
CC = gcc

# Flags
ASFLAGS = -f win64
LDFLAGS = 

# Default file names (can be overridden by command-line variables)
ASM_FILE ?= output.asm
OBJ_FILE ?= output.obj
EXE_FILE ?= output.exe

# Build rules
all: $(EXE_FILE)

$(EXE_FILE): $(OBJ_FILE)
	$(CC) $(LDFLAGS) -o $(EXE_FILE) $(OBJ_FILE)
	del /f $(OBJ_FILE)

$(OBJ_FILE): $(ASM_FILE)
	$(AS) $(ASFLAGS) -o $(OBJ_FILE) $(ASM_FILE)
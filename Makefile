# Compiler and assembler
AS = nasm
LD = ld

# Flags
ASFLAGS = -f elf64
LDFLAGS = 

# Default file names (can be overridden by command-line variables)
ASM_FILE ?= output.asm
OBJ_FILE ?= output.obj
EXE_FILE ?= output.exe

# Build rules
all: $(EXE_FILE)

$(EXE_FILE): $(OBJ_FILE)
	$(LD) $(LDFLAGS) -o $(EXE_FILE) $(OBJ_FILE)
	rm -f $(OBJ_FILE)

$(OBJ_FILE): $(ASM_FILE)
	$(AS) $(ASFLAGS) -o $(OBJ_FILE) $(ASM_FILE)

clean:
	rm -f $(OBJ_FILE) $(EXE_FILE)
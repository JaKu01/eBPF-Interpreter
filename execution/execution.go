package execution

import (
	"eBPF-Interpreter/constants"
	"eBPF-Interpreter/parse"
	"eBPF-Interpreter/types"
	"encoding/binary"
	"fmt"
)

var applicationState types.ApplicationState

func ExecuteProgram(instructions []types.Instruction) {
	for instructionCounter := 0; instructionCounter < len(instructions); instructionCounter++ {
		instruction := instructions[instructionCounter]

		switch parse.GetInstructionClass(instruction.Opcode) {
		case constants.BPF_LD:
			handleNonStandardLoadOperations(instruction)
		case constants.BPF_LDX:
			handleLoadIntoRegisterOperations(instruction)
		case constants.BPF_ST:
			handleStoreFromImmediateOperations(instruction)
		case constants.BPF_STX:
			handleStoreFromRegisterOperations(instruction)
		case constants.BPF_ALU:
			handle32bitArithmeticOperations(instruction)
		case constants.BPF_JMP:
			handle64bitJumpInstructions(instruction)
		case constants.BPF_JMP32:
			handle32bitJumpInstructions(instruction)
		case constants.BPF_ALU64:
			handle64bitArithmeticOperations(instruction)
		}

	}
}

func handleNonStandardLoadOperations(instruction types.Instruction) {
	switch instruction.Opcode {
	case constants.BPF_LD | constants.BPF_W | constants.BPF_IMM:
		fmt.Println("BPF_LD | BPF_W | BPF_IMM")
	case constants.BPF_LD | constants.BPF_H | constants.BPF_IMM:
		fmt.Println("BPF_LD | BPF_H | BPF_IMM")
	case constants.BPF_LD | constants.BPF_B | constants.BPF_IMM:
		fmt.Println("BPF_LD | BPF_B | BPF_IMM")
	case constants.BPF_LD | constants.BPF_DW | constants.BPF_IMM:
		fmt.Println("BPF_LD | BPF_DW | BPF_IMM")
	case constants.BPF_LD | constants.BPF_W | constants.BPF_ABS:
		fmt.Println("BPF_LD | BPF_W | BPF_ABS")
	case constants.BPF_LD | constants.BPF_H | constants.BPF_ABS:
		fmt.Println("BPF_LD | BPF_H | BPF_ABS")
	case constants.BPF_LD | constants.BPF_B | constants.BPF_ABS:
		fmt.Println("BPF_LD | BPF_B | BPF_ABS")
	case constants.BPF_LD | constants.BPF_DW | constants.BPF_ABS:
		fmt.Println("BPF_LD | BPF_DW | BPF_ABS")
	case constants.BPF_LD | constants.BPF_W | constants.BPF_IND:
		fmt.Println("BPF_LD | BPF_W | BPF_IND")
	case constants.BPF_LD | constants.BPF_H | constants.BPF_IND:
		fmt.Println("BPF_LD | BPF_H | BPF_IND")
	case constants.BPF_LD | constants.BPF_B | constants.BPF_IND:
		fmt.Println("BPF_LD | BPF_B | BPF_IND")
	case constants.BPF_LD | constants.BPF_DW | constants.BPF_IND:
		fmt.Println("BPF_LD | BPF_DW | BPF_IND")
	case constants.BPF_LD | constants.BPF_W | constants.BPF_MEM:
		fmt.Println("BPF_LD | BPF_W | BPF_MEM")
	case constants.BPF_LD | constants.BPF_H | constants.BPF_MEM:
		fmt.Println("BPF_LD | BPF_H | BPF_MEM")
	case constants.BPF_LD | constants.BPF_B | constants.BPF_MEM:
		fmt.Println("BPF_LD | BPF_B | BPF_MEM")
	case constants.BPF_LD | constants.BPF_DW | constants.BPF_MEM:
		fmt.Println("BPF_LD | BPF_DW | BPF_MEM")
	case constants.BPF_LD | constants.BPF_W | constants.BPF_ATOMIC:
		fmt.Println("BPF_LD | BPF_W | BPF_ATOMIC")
	case constants.BPF_LD | constants.BPF_H | constants.BPF_ATOMIC:
		fmt.Println("BPF_LD | BPF_H | BPF_ATOMIC")
	case constants.BPF_LD | constants.BPF_B | constants.BPF_ATOMIC:
		fmt.Println("BPF_LD | BPF_B | BPF_ATOMIC")
	case constants.BPF_LD | constants.BPF_DW | constants.BPF_ATOMIC:
		fmt.Println("BPF_LD | BPF_DW | BPF_ATOMIC")
	}
}

func handleLoadIntoRegisterOperations(instruction types.Instruction) {
	switch instruction.Opcode {
	case constants.BPF_LDX | constants.BPF_W | constants.BPF_IMM:
		fmt.Println("BPF_LDX | BPF_W | BPF_IMM")
	case constants.BPF_LDX | constants.BPF_H | constants.BPF_IMM:
		fmt.Println("BPF_LDX | BPF_H | BPF_IMM")
	case constants.BPF_LDX | constants.BPF_B | constants.BPF_IMM:
		fmt.Println("BPF_LDX | BPF_B | BPF_IMM")
	case constants.BPF_LDX | constants.BPF_DW | constants.BPF_IMM:
		fmt.Println("BPF_LDX | BPF_DW | BPF_IMM")
	case constants.BPF_LDX | constants.BPF_W | constants.BPF_ABS:
		fmt.Println("BPF_LDX | BPF_W | BPF_ABS")
	case constants.BPF_LDX | constants.BPF_H | constants.BPF_ABS:
		fmt.Println("BPF_LDX | BPF_H | BPF_ABS")
	case constants.BPF_LDX | constants.BPF_B | constants.BPF_ABS:
		fmt.Println("BPF_LDX | BPF_B | BPF_ABS")
	case constants.BPF_LDX | constants.BPF_DW | constants.BPF_ABS:
		fmt.Println("BPF_LDX | BPF_DW | BPF_ABS")
	case constants.BPF_LDX | constants.BPF_W | constants.BPF_IND:
		fmt.Println("BPF_LDX | BPF_W | BPF_IND")
	case constants.BPF_LDX | constants.BPF_H | constants.BPF_IND:
		fmt.Println("BPF_LDX | BPF_H | BPF_IND")
	case constants.BPF_LDX | constants.BPF_B | constants.BPF_IND:
		fmt.Println("BPF_LDX | BPF_B | BPF_IND")
	case constants.BPF_LDX | constants.BPF_DW | constants.BPF_IND:
		fmt.Println("BPF_LDX | BPF_DW | BPF_IND")
	case constants.BPF_LDX | constants.BPF_W | constants.BPF_MEM:
		fmt.Println("BPF_LDX | BPF_W | BPF_MEM")
	case constants.BPF_LDX | constants.BPF_H | constants.BPF_MEM:
		fmt.Println("BPF_LDX | BPF_H | BPF_MEM")
	case constants.BPF_LDX | constants.BPF_B | constants.BPF_MEM:
		fmt.Println("BPF_LDX | BPF_B | BPF_MEM")
	case constants.BPF_LDX | constants.BPF_DW | constants.BPF_MEM:
		fmt.Println("BPF_LDX | BPF_DW | BPF_MEM")
	case constants.BPF_LDX | constants.BPF_W | constants.BPF_ATOMIC:
		fmt.Println("BPF_LDX | BPF_W | BPF_ATOMIC")
	case constants.BPF_LDX | constants.BPF_H | constants.BPF_ATOMIC:
		fmt.Println("BPF_LDX | BPF_H | BPF_ATOMIC")
	case constants.BPF_LDX | constants.BPF_B | constants.BPF_ATOMIC:
		fmt.Println("BPF_LDX | BPF_B | BPF_ATOMIC")
	case constants.BPF_LDX | constants.BPF_DW | constants.BPF_ATOMIC:
		fmt.Println("BPF_LDX | BPF_DW | BPF_ATOMIC")
	}
}

func handleStoreFromImmediateOperations(instruction types.Instruction) {
	switch instruction.Opcode {
	case constants.BPF_ST | constants.BPF_W | constants.BPF_IMM:
		fmt.Println("BPF_ST | BPF_W | BPF_IMM")
	case constants.BPF_ST | constants.BPF_H | constants.BPF_IMM:
		fmt.Println("BPF_ST | BPF_H | BPF_IMM")
	case constants.BPF_ST | constants.BPF_B | constants.BPF_IMM:
		fmt.Println("BPF_ST | BPF_B | BPF_IMM")
	case constants.BPF_ST | constants.BPF_DW | constants.BPF_IMM:
		fmt.Println("BPF_ST | BPF_DW | BPF_IMM")
	case constants.BPF_ST | constants.BPF_W | constants.BPF_ABS:
		fmt.Println("BPF_ST | BPF_W | BPF_ABS")
	case constants.BPF_ST | constants.BPF_H | constants.BPF_ABS:
		fmt.Println("BPF_ST | BPF_H | BPF_ABS")
	case constants.BPF_ST | constants.BPF_B | constants.BPF_ABS:
		fmt.Println("BPF_ST | BPF_B | BPF_ABS")
	case constants.BPF_ST | constants.BPF_DW | constants.BPF_ABS:
		fmt.Println("BPF_ST | BPF_DW | BPF_ABS")
	case constants.BPF_ST | constants.BPF_W | constants.BPF_IND:
		fmt.Println("BPF_ST | BPF_W | BPF_IND")
	case constants.BPF_ST | constants.BPF_H | constants.BPF_IND:
		fmt.Println("BPF_ST | BPF_H | BPF_IND")
	case constants.BPF_ST | constants.BPF_B | constants.BPF_IND:
		fmt.Println("BPF_ST | BPF_B | BPF_IND")
	case constants.BPF_ST | constants.BPF_DW | constants.BPF_IND:
		fmt.Println("BPF_ST | BPF_DW | BPF_IND")
	case constants.BPF_ST | constants.BPF_W | constants.BPF_MEM:
		fmt.Println("BPF_ST | BPF_W | BPF_MEM")
	case constants.BPF_ST | constants.BPF_H | constants.BPF_MEM:
		fmt.Println("BPF_ST | BPF_H | BPF_MEM")
	case constants.BPF_ST | constants.BPF_B | constants.BPF_MEM:
		fmt.Println("BPF_ST | BPF_B | BPF_MEM")
	case constants.BPF_ST | constants.BPF_DW | constants.BPF_MEM:
		fmt.Println("BPF_ST | BPF_DW | BPF_MEM")
	case constants.BPF_ST | constants.BPF_W | constants.BPF_ATOMIC:
		fmt.Println("BPF_ST | BPF_W | BPF_ATOMIC")
	case constants.BPF_ST | constants.BPF_H | constants.BPF_ATOMIC:
		fmt.Println("BPF_ST | BPF_H | BPF_ATOMIC")
	case constants.BPF_ST | constants.BPF_B | constants.BPF_ATOMIC:
		fmt.Println("BPF_ST | BPF_B | BPF_ATOMIC")
	case constants.BPF_ST | constants.BPF_DW | constants.BPF_ATOMIC:
		fmt.Println("BPF_ST | BPF_DW | BPF_ATOMIC")
	}
}

func handleStoreFromRegisterOperations(instruction types.Instruction) {
	switch instruction.Opcode {
	case constants.BPF_STX | constants.BPF_W | constants.BPF_IMM:
		fmt.Println("BPF_STX | BPF_W | BPF_IMM")
	case constants.BPF_STX | constants.BPF_H | constants.BPF_IMM:
		fmt.Println("BPF_STX | BPF_H | BPF_IMM")
	case constants.BPF_STX | constants.BPF_B | constants.BPF_IMM:
		fmt.Println("BPF_STX | BPF_B | BPF_IMM")
	case constants.BPF_STX | constants.BPF_DW | constants.BPF_IMM:
		fmt.Println("BPF_STX | BPF_DW | BPF_IMM")
	case constants.BPF_STX | constants.BPF_W | constants.BPF_ABS:
		fmt.Println("BPF_STX | BPF_W | BPF_ABS")
	case constants.BPF_STX | constants.BPF_H | constants.BPF_ABS:
		fmt.Println("BPF_STX | BPF_H | BPF_ABS")
	case constants.BPF_STX | constants.BPF_B | constants.BPF_ABS:
		fmt.Println("BPF_STX | BPF_B | BPF_ABS")
	case constants.BPF_STX | constants.BPF_DW | constants.BPF_ABS:
		fmt.Println("BPF_STX | BPF_DW | BPF_ABS")
	case constants.BPF_STX | constants.BPF_W | constants.BPF_IND:
		fmt.Println("BPF_STX | BPF_W | BPF_IND")
	case constants.BPF_STX | constants.BPF_H | constants.BPF_IND:
		fmt.Println("BPF_STX | BPF_H | BPF_IND")
	case constants.BPF_STX | constants.BPF_B | constants.BPF_IND:
		fmt.Println("BPF_STX | BPF_B | BPF_IND")
	case constants.BPF_STX | constants.BPF_DW | constants.BPF_IND:
		fmt.Println("BPF_STX | BPF_DW | BPF_IND")
	case constants.BPF_STX | constants.BPF_W | constants.BPF_MEM:
		fmt.Println("BPF_STX | BPF_W | BPF_MEM")
	case constants.BPF_STX | constants.BPF_H | constants.BPF_MEM:
		fmt.Println("BPF_STX | BPF_H | BPF_MEM")
	case constants.BPF_STX | constants.BPF_B | constants.BPF_MEM:
		fmt.Println("BPF_STX | BPF_B | BPF_MEM")
	case constants.BPF_STX | constants.BPF_DW | constants.BPF_MEM:
		fmt.Println("BPF_STX | BPF_DW | BPF_MEM")
	case constants.BPF_STX | constants.BPF_W | constants.BPF_ATOMIC:
		fmt.Println("BPF_STX | BPF_W | BPF_ATOMIC")
	case constants.BPF_STX | constants.BPF_H | constants.BPF_ATOMIC:
		fmt.Println("BPF_STX | BPF_H | BPF_ATOMIC")
	case constants.BPF_STX | constants.BPF_B | constants.BPF_ATOMIC:
		fmt.Println("BPF_STX | BPF_B | BPF_ATOMIC")
	case constants.BPF_STX | constants.BPF_DW | constants.BPF_ATOMIC:
		fmt.Println("BPF_STX | BPF_DW | BPF_ATOMIC")
	}
}

func handle32bitArithmeticOperations(instruction types.Instruction) {
	switch instruction.Opcode {
	// 32 immediate as source operand
	case constants.BPF_ALU | constants.BPF_K | constants.BPF_ADD:
		fmt.Println("BPF_ALU | BPF_K | BPF_ADD")
	case constants.BPF_ALU | constants.BPF_K | constants.BPF_SUB:
		fmt.Println("BPF_ALU | BPF_K | BPF_SUB")
	case constants.BPF_ALU | constants.BPF_K | constants.BPF_MUL:
		fmt.Println("BPF_ALU | BPF_K | BPF_MUL")
	case constants.BPF_ALU | constants.BPF_K | constants.BPF_DIV:
		fmt.Println("BPF_ALU | BPF_K | BPF_DIV")
	case constants.BPF_ALU | constants.BPF_K | constants.BPF_OR:
		fmt.Println("BPF_ALU | BPF_K | BPF_OR")
	case constants.BPF_ALU | constants.BPF_K | constants.BPF_AND:
		fmt.Println("BPF_ALU | BPF_K | BPF_AND")
	case constants.BPF_ALU | constants.BPF_K | constants.BPF_LSH:
		fmt.Println("BPF_ALU | BPF_K | BPF_LSH")
	case constants.BPF_ALU | constants.BPF_K | constants.BPF_RSH:
		fmt.Println("BPF_ALU | BPF_K | BPF_RSH")
	case constants.BPF_ALU | constants.BPF_K | constants.BPF_NEG:
		fmt.Println("BPF_ALU | BPF_K | BPF_NEG")
	case constants.BPF_ALU | constants.BPF_K | constants.BPF_MOD:
		fmt.Println("BPF_ALU | BPF_K | BPF_MOD")
	case constants.BPF_ALU | constants.BPF_K | constants.BPF_XOR:
		fmt.Println("BPF_ALU | BPF_K | BPF_XOR")
	case constants.BPF_ALU | constants.BPF_K | constants.BPF_MOV:
		fmt.Println("BPF_ALU | BPF_K | BPF_MOV")
	case constants.BPF_ALU | constants.BPF_K | constants.BPF_ARSH:
		fmt.Println("BPF_ALU | BPF_K | BPF_ARSH")
	case constants.BPF_ALU | constants.BPF_K | constants.BPF_END:
		fmt.Println("BPF_ALU | BPF_K | BPF_END")

	// register as source operand
	case constants.BPF_ALU | constants.BPF_X | constants.BPF_ADD:
		fmt.Println("BPF_ALU | BPF_X | BPF_ADD")
	case constants.BPF_ALU | constants.BPF_X | constants.BPF_SUB:
		fmt.Println("BPF_ALU | BPF_X | BPF_SUB")
	case constants.BPF_ALU | constants.BPF_X | constants.BPF_MUL:
		fmt.Println("BPF_ALU | BPF_X | BPF_MUL")
	case constants.BPF_ALU | constants.BPF_X | constants.BPF_DIV:
		fmt.Println("BPF_ALU | BPF_X | BPF_DIV")
	case constants.BPF_ALU | constants.BPF_X | constants.BPF_OR:
		fmt.Println("BPF_ALU | BPF_X | BPF_OR")
	case constants.BPF_ALU | constants.BPF_X | constants.BPF_AND:
		fmt.Println("BPF_ALU | BPF_X | BPF_AND")
	case constants.BPF_ALU | constants.BPF_X | constants.BPF_LSH:
		fmt.Println("BPF_ALU | BPF_X | BPF_LSH")
	case constants.BPF_ALU | constants.BPF_X | constants.BPF_RSH:
		fmt.Println("BPF_ALU | BPF_X | BPF_RSH")
	case constants.BPF_ALU | constants.BPF_X | constants.BPF_NEG:
		fmt.Println("BPF_ALU | BPF_X | BPF_NEG")
	case constants.BPF_ALU | constants.BPF_X | constants.BPF_MOD:
		fmt.Println("BPF_ALU | BPF_X | BPF_MOD")
	case constants.BPF_ALU | constants.BPF_X | constants.BPF_XOR:
		fmt.Println("BPF_ALU | BPF_X | BPF_XOR")
	case constants.BPF_ALU | constants.BPF_X | constants.BPF_MOV:
		fmt.Println("BPF_ALU | BPF_X | BPF_MOV")
	case constants.BPF_ALU | constants.BPF_X | constants.BPF_ARSH:
		fmt.Println("BPF_ALU | BPF_X | BPF_ARSH")
	case constants.BPF_ALU | constants.BPF_X | constants.BPF_END:
		fmt.Println("BPF_ALU | BPF_X | BPF_END")
	}

}

func handle64bitArithmeticOperations(instruction types.Instruction) {
	switch instruction.Opcode {
	// 32 immediate as source operand
	case constants.BPF_ALU64 | constants.BPF_K | constants.BPF_ADD:
		fmt.Println("BPF_ALU64 | BPF_K | BPF_ADD")
		applicationState.Registers[instruction.DestinationRegister] += uint64(instruction.Immediate)
	case constants.BPF_ALU64 | constants.BPF_K | constants.BPF_SUB:
		fmt.Println("BPF_ALU64 | BPF_K | BPF_SUB")
		applicationState.Registers[instruction.DestinationRegister] -= uint64(instruction.Immediate)
	case constants.BPF_ALU64 | constants.BPF_K | constants.BPF_MUL:
		fmt.Println("BPF_ALU64 | BPF_K | BPF_MUL")
		applicationState.Registers[instruction.DestinationRegister] *= uint64(instruction.Immediate)
	case constants.BPF_ALU64 | constants.BPF_K | constants.BPF_DIV:
		fmt.Println("BPF_ALU64 | BPF_K | BPF_DIV")
		applicationState.Registers[instruction.DestinationRegister] /= uint64(instruction.Immediate)
	case constants.BPF_ALU64 | constants.BPF_K | constants.BPF_OR:
		fmt.Println("BPF_ALU64 | BPF_K | BPF_OR")
		applicationState.Registers[instruction.DestinationRegister] |= uint64(instruction.Immediate)
	case constants.BPF_ALU64 | constants.BPF_K | constants.BPF_AND:
		fmt.Println("BPF_ALU64 | BPF_K | BPF_AND")
		applicationState.Registers[instruction.DestinationRegister] &= uint64(instruction.Immediate)
	case constants.BPF_ALU64 | constants.BPF_K | constants.BPF_LSH:
		fmt.Println("BPF_ALU64 | BPF_K | BPF_LSH")
		applicationState.Registers[instruction.DestinationRegister] <<= uint64(instruction.Immediate)
	case constants.BPF_ALU64 | constants.BPF_K | constants.BPF_RSH:
		fmt.Println("BPF_ALU64 | BPF_K | BPF_RSH")
		applicationState.Registers[instruction.DestinationRegister] >>= uint64(instruction.Immediate)
	case constants.BPF_ALU64 | constants.BPF_K | constants.BPF_NEG:
		fmt.Println("BPF_ALU64 | BPF_K | BPF_NEG")
		applicationState.Registers[instruction.DestinationRegister] = uint64(^instruction.Immediate)
	case constants.BPF_ALU64 | constants.BPF_K | constants.BPF_MOD:
		fmt.Println("BPF_ALU64 | BPF_K | BPF_MOD")
		applicationState.Registers[instruction.DestinationRegister] %= uint64(instruction.Immediate)
	case constants.BPF_ALU64 | constants.BPF_K | constants.BPF_XOR:
		fmt.Println("BPF_ALU64 | BPF_K | BPF_XOR")
		applicationState.Registers[instruction.DestinationRegister] ^= uint64(instruction.Immediate)
	case constants.BPF_ALU64 | constants.BPF_K | constants.BPF_MOV:
		fmt.Println("BPF_ALU64 | BPF_K | BPF_MOV")
		applicationState.Registers[instruction.DestinationRegister] = uint64(instruction.Immediate)
	case constants.BPF_ALU64 | constants.BPF_K | constants.BPF_ARSH:
		fmt.Println("BPF_ALU64 | BPF_K | BPF_ARSH")
		signedValueFromDestinationRegister := int64(applicationState.Registers[instruction.DestinationRegister])
		shiftedValue := signedValueFromDestinationRegister >> uint64(instruction.Immediate)
		applicationState.Registers[instruction.DestinationRegister] = uint64(shiftedValue)

	// register as source operand
	case constants.BPF_ALU64 | constants.BPF_X | constants.BPF_ADD:
		fmt.Println("BPF_ALU64 | BPF_X | BPF_ADD")
		applicationState.Registers[instruction.DestinationRegister] += applicationState.Registers[instruction.SourceRegister]
	case constants.BPF_ALU64 | constants.BPF_X | constants.BPF_SUB:
		fmt.Println("BPF_ALU64 | BPF_X | BPF_SUB")
		applicationState.Registers[instruction.DestinationRegister] -= applicationState.Registers[instruction.SourceRegister]
	case constants.BPF_ALU64 | constants.BPF_X | constants.BPF_MUL:
		fmt.Println("BPF_ALU64 | BPF_X | BPF_MUL")
		applicationState.Registers[instruction.DestinationRegister] *= applicationState.Registers[instruction.SourceRegister]
	case constants.BPF_ALU64 | constants.BPF_X | constants.BPF_DIV:
		fmt.Println("BPF_ALU64 | BPF_X | BPF_DIV")
		applicationState.Registers[instruction.DestinationRegister] /= applicationState.Registers[instruction.SourceRegister]
	case constants.BPF_ALU64 | constants.BPF_X | constants.BPF_OR:
		fmt.Println("BPF_ALU64 | BPF_X | BPF_OR")
		applicationState.Registers[instruction.DestinationRegister] |= applicationState.Registers[instruction.SourceRegister]
	case constants.BPF_ALU64 | constants.BPF_X | constants.BPF_AND:
		fmt.Println("BPF_ALU64 | BPF_X | BPF_AND")
		applicationState.Registers[instruction.DestinationRegister] &= applicationState.Registers[instruction.SourceRegister]
	case constants.BPF_ALU64 | constants.BPF_X | constants.BPF_LSH:
		fmt.Println("BPF_ALU64 | BPF_X | BPF_LSH")
		applicationState.Registers[instruction.DestinationRegister] <<= applicationState.Registers[instruction.SourceRegister]
	case constants.BPF_ALU64 | constants.BPF_X | constants.BPF_RSH:
		fmt.Println("BPF_ALU64 | BPF_X | BPF_RSH")
		applicationState.Registers[instruction.DestinationRegister] >>= applicationState.Registers[instruction.SourceRegister]
	case constants.BPF_ALU64 | constants.BPF_X | constants.BPF_NEG:
		fmt.Println("BPF_ALU64 | BPF_X | BPF_NEG")
		applicationState.Registers[instruction.DestinationRegister] = ^applicationState.Registers[instruction.SourceRegister]
	case constants.BPF_ALU64 | constants.BPF_X | constants.BPF_MOD:
		fmt.Println("BPF_ALU64 | BPF_X | BPF_MOD")
		applicationState.Registers[instruction.DestinationRegister] %= applicationState.Registers[instruction.SourceRegister]
	case constants.BPF_ALU64 | constants.BPF_X | constants.BPF_XOR:
		fmt.Println("BPF_ALU64 | BPF_X | BPF_XOR")
		applicationState.Registers[instruction.DestinationRegister] ^= applicationState.Registers[instruction.SourceRegister]
	case constants.BPF_ALU64 | constants.BPF_X | constants.BPF_MOV:
		fmt.Println("BPF_ALU64 | BPF_X | BPF_MOV")
		applicationState.Registers[instruction.DestinationRegister] = applicationState.Registers[instruction.SourceRegister]
	case constants.BPF_ALU64 | constants.BPF_X | constants.BPF_ARSH:
		fmt.Println("BPF_ALU64 | BPF_X | BPF_ARSH")
		signedValueFromDestinationRegister := int64(applicationState.Registers[instruction.DestinationRegister])
		shiftedValue := signedValueFromDestinationRegister >> applicationState.Registers[instruction.SourceRegister]
		applicationState.Registers[instruction.DestinationRegister] = uint64(shiftedValue)
	case constants.BPF_ALU64 | constants.BPF_TO_LE | constants.BPF_END:
		fmt.Println("BPF_ALU64 | BPF_X | BPF_END")
		tempConvertedValue := make([]byte, 8)
		switch instruction.Immediate {
		case 0x10:
			fmt.Println("16 bit")
			binary.NativeEndian.PutUint16(tempConvertedValue, uint16(applicationState.Registers[instruction.DestinationRegister]))
			applicationState.Registers[instruction.DestinationRegister] = binary.LittleEndian.Uint64(tempConvertedValue)
		case 0x20:
			fmt.Println("32 bit")
			binary.NativeEndian.PutUint32(tempConvertedValue, uint32(applicationState.Registers[instruction.DestinationRegister]))
			applicationState.Registers[instruction.DestinationRegister] = binary.LittleEndian.Uint64(tempConvertedValue)
		case 0x40:
			fmt.Println("64 bit")
			binary.NativeEndian.PutUint64(tempConvertedValue, applicationState.Registers[instruction.DestinationRegister])
			applicationState.Registers[instruction.DestinationRegister] = binary.LittleEndian.Uint64(tempConvertedValue)
		default:
			fmt.Println("Invalid size for byte swap width")
		}

	case constants.BPF_ALU64 | constants.BPF_TO_BE | constants.BPF_END:
		fmt.Println("BPF_ALU64 | BPF_X | BPF_END")
		tempConvertedValue := make([]byte, 8)
		switch instruction.Immediate {
		case 0x10:
			fmt.Println("16 bit")
			binary.NativeEndian.PutUint16(tempConvertedValue, uint16(applicationState.Registers[instruction.DestinationRegister]))
			applicationState.Registers[instruction.DestinationRegister] = binary.BigEndian.Uint64(tempConvertedValue)
		case 0x20:
			fmt.Println("32 bit")
			binary.NativeEndian.PutUint32(tempConvertedValue, uint32(applicationState.Registers[instruction.DestinationRegister]))
			applicationState.Registers[instruction.DestinationRegister] = binary.BigEndian.Uint64(tempConvertedValue)
		case 0x40:
			fmt.Println("64 bit")
			binary.NativeEndian.PutUint64(tempConvertedValue, applicationState.Registers[instruction.DestinationRegister])
			applicationState.Registers[instruction.DestinationRegister] = binary.BigEndian.Uint64(tempConvertedValue)
		default:
			fmt.Println("Invalid size for byte swap width")
		}
	default:
		fmt.Println("Invalid opcode for 64 bit arithmetic operations")
	}
}

func handle64bitJumpInstructions(instruction types.Instruction) {
	switch instruction.Opcode {
	// 32 bit immediate as source operand
	case constants.BPF_JMP | constants.BPF_K | constants.BPF_JA:
		fmt.Println("BPF_JMP | BPF_K | BPF_JA")
	case constants.BPF_JMP | constants.BPF_K | constants.BPF_JEQ:
		fmt.Println("BPF_JMP | BPF_K | BPF_JEQ")
	case constants.BPF_JMP | constants.BPF_K | constants.BPF_JGT:
		fmt.Println("BPF_JMP | BPF_K | BPF_JGT")
	case constants.BPF_JMP | constants.BPF_K | constants.BPF_JGE:
		fmt.Println("BPF_JMP | BPF_K | BPF_JGE")
	case constants.BPF_JMP | constants.BPF_K | constants.BPF_JSET:
		fmt.Println("BPF_JMP | BPF_K | BPF_JSET")
	case constants.BPF_JMP | constants.BPF_K | constants.BPF_JNE:
		fmt.Println("BPF_JMP | BPF_K | BPF_JNE")
	case constants.BPF_JMP | constants.BPF_K | constants.BPF_JSGT:
		fmt.Println("BPF_JMP | BPF_K | BPF_JSGT")
	case constants.BPF_JMP | constants.BPF_K | constants.BPF_JSGE:
		fmt.Println("BPF_JMP | BPF_K | BPF_JSGE")
	case constants.BPF_JMP | constants.BPF_K | constants.BPF_CALL:
		fmt.Println("BPF_JMP | BPF_K | BPF_CALL")
	case constants.BPF_JMP | constants.BPF_K | constants.BPF_EXIT:
		fmt.Println("BPF_JMP | BPF_K | BPF_EXIT")
	case constants.BPF_JMP | constants.BPF_K | constants.BPF_JLT:
		fmt.Println("BPF_JMP | BPF_K | BPF_JLT")
	case constants.BPF_JMP | constants.BPF_K | constants.BPF_JLE:
		fmt.Println("BPF_JMP | BPF_K | BPF_JLE")
	case constants.BPF_JMP | constants.BPF_K | constants.BPF_JSLT:
		fmt.Println("BPF_JMP | BPF_K | BPF_JSLT")
	case constants.BPF_JMP | constants.BPF_K | constants.BPF_JSLE:
		fmt.Println("BPF_JMP | BPF_K | BPF_JSLE")

	// register as source operand
	case constants.BPF_JMP | constants.BPF_X | constants.BPF_JA:
		fmt.Println("BPF_JMP | BPF_X | BPF_JA")
	case constants.BPF_JMP | constants.BPF_X | constants.BPF_JEQ:
		fmt.Println("BPF_JMP | BPF_X | BPF_JEQ")
	case constants.BPF_JMP | constants.BPF_X | constants.BPF_JGT:
		fmt.Println("BPF_JMP | BPF_X | BPF_JGT")
	case constants.BPF_JMP | constants.BPF_X | constants.BPF_JGE:
		fmt.Println("BPF_JMP | BPF_X | BPF_JGE")
	case constants.BPF_JMP | constants.BPF_X | constants.BPF_JSET:
		fmt.Println("BPF_JMP | BPF_X | BPF_JSET")
	case constants.BPF_JMP | constants.BPF_X | constants.BPF_JNE:
		fmt.Println("BPF_JMP | BPF_X | BPF_JNE")
	case constants.BPF_JMP | constants.BPF_X | constants.BPF_JSGT:
		fmt.Println("BPF_JMP | BPF_X | BPF_JSGT")
	case constants.BPF_JMP | constants.BPF_X | constants.BPF_JSGE:
		fmt.Println("BPF_JMP | BPF_X | BPF_JSGE")
	case constants.BPF_JMP | constants.BPF_X | constants.BPF_CALL:
		fmt.Println("BPF_JMP | BPF_X | BPF_CALL")
	case constants.BPF_JMP | constants.BPF_X | constants.BPF_EXIT:
		fmt.Println("BPF_JMP | BPF_X | BPF_EXIT")
	case constants.BPF_JMP | constants.BPF_X | constants.BPF_JLT:
		fmt.Println("BPF_JMP | BPF_X | BPF_JLT")
	case constants.BPF_JMP | constants.BPF_X | constants.BPF_JLE:
		fmt.Println("BPF_JMP | BPF_X | BPF_JLE")
	case constants.BPF_JMP | constants.BPF_X | constants.BPF_JSLT:
		fmt.Println("BPF_JMP | BPF_X | BPF_JSLT")
	case constants.BPF_JMP | constants.BPF_X | constants.BPF_JSLE:
		fmt.Println("BPF_JMP | BPF_X | BPF_JSLE")
	}
}

func handle32bitJumpInstructions(instruction types.Instruction) {
	switch instruction.Opcode {
	// 32 bit immediate as source operand
	case constants.BPF_JMP32 | constants.BPF_K | constants.BPF_JA:
		fmt.Println("BPF_JMP32 | BPF_K | BPF_JA")
	case constants.BPF_JMP32 | constants.BPF_K | constants.BPF_JEQ:
		fmt.Println("BPF_JMP32 | BPF_K | BPF_JEQ")
	case constants.BPF_JMP32 | constants.BPF_K | constants.BPF_JGT:
		fmt.Println("BPF_JMP32 | BPF_K | BPF_JGT")
	case constants.BPF_JMP32 | constants.BPF_K | constants.BPF_JGE:
		fmt.Println("BPF_JMP32 | BPF_K | BPF_JGE")
	case constants.BPF_JMP32 | constants.BPF_K | constants.BPF_JSET:
		fmt.Println("BPF_JMP32 | BPF_K | BPF_JSET")
	case constants.BPF_JMP32 | constants.BPF_K | constants.BPF_JNE:
		fmt.Println("BPF_JMP32 | BPF_K | BPF_JNE")
	case constants.BPF_JMP32 | constants.BPF_K | constants.BPF_JSGT:
		fmt.Println("BPF_JMP32 | BPF_K | BPF_JSGT")
	case constants.BPF_JMP32 | constants.BPF_K | constants.BPF_JSGE:
		fmt.Println("BPF_JMP32 | BPF_K | BPF_JSGE")
	case constants.BPF_JMP32 | constants.BPF_K | constants.BPF_CALL:
		fmt.Println("BPF_JMP32 | BPF_K | BPF_CALL")
	case constants.BPF_JMP32 | constants.BPF_K | constants.BPF_EXIT:
		fmt.Println("BPF_JMP32 | BPF_K | BPF_EXIT")
	case constants.BPF_JMP32 | constants.BPF_K | constants.BPF_JLT:
		fmt.Println("BPF_JMP32 | BPF_K | BPF_JLT")
	case constants.BPF_JMP32 | constants.BPF_K | constants.BPF_JLE:
		fmt.Println("BPF_JMP32 | BPF_K | BPF_JLE")
	case constants.BPF_JMP32 | constants.BPF_K | constants.BPF_JSLT:
		fmt.Println("BPF_JMP32 | BPF_K | BPF_JSLT")
	case constants.BPF_JMP32 | constants.BPF_K | constants.BPF_JSLE:
		fmt.Println("BPF_JMP32 | BPF_K | BPF_JSLE")

	// register as source operand
	case constants.BPF_JMP32 | constants.BPF_X | constants.BPF_JA:
		fmt.Println("BPF_JMP32 | BPF_X | BPF_JA")
	case constants.BPF_JMP32 | constants.BPF_X | constants.BPF_JEQ:
		fmt.Println("BPF_JMP32 | BPF_X | BPF_JEQ")
	case constants.BPF_JMP32 | constants.BPF_X | constants.BPF_JGT:
		fmt.Println("BPF_JMP32 | BPF_X | BPF_JGT")
	case constants.BPF_JMP32 | constants.BPF_X | constants.BPF_JGE:
		fmt.Println("BPF_JMP32 | BPF_X | BPF_JGE")
	case constants.BPF_JMP32 | constants.BPF_X | constants.BPF_JSET:
		fmt.Println("BPF_JMP32 | BPF_X | BPF_JSET")
	case constants.BPF_JMP32 | constants.BPF_X | constants.BPF_JNE:
		fmt.Println("BPF_JMP32 | BPF_X | BPF_JNE")
	case constants.BPF_JMP32 | constants.BPF_X | constants.BPF_JSGT:
		fmt.Println("BPF_JMP32 | BPF_X | BPF_JSGT")
	case constants.BPF_JMP32 | constants.BPF_X | constants.BPF_JSGE:
		fmt.Println("BPF_JMP32 | BPF_X | BPF_JSGE")
	case constants.BPF_JMP32 | constants.BPF_X | constants.BPF_CALL:
		fmt.Println("BPF_JMP32 | BPF_X | BPF_CALL")
	case constants.BPF_JMP32 | constants.BPF_X | constants.BPF_EXIT:
		fmt.Println("BPF_JMP32 | BPF_X | BPF_EXIT")
	case constants.BPF_JMP32 | constants.BPF_X | constants.BPF_JLT:
		fmt.Println("BPF_JMP32 | BPF_X | BPF_JLT")
	case constants.BPF_JMP32 | constants.BPF_X | constants.BPF_JLE:
		fmt.Println("BPF_JMP32 | BPF_X | BPF_JLE")
	case constants.BPF_JMP32 | constants.BPF_X | constants.BPF_JSLT:
		fmt.Println("BPF_JMP32 | BPF_X | BPF_JSLT")
	case constants.BPF_JMP32 | constants.BPF_X | constants.BPF_JSLE:
		fmt.Println("BPF_JMP32 | BPF_X | BPF_JSLE")
	}
}

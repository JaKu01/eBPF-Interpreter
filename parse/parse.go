package parse

import (
	"eBPF-Interpreter/types"
	"encoding/binary"
	"fmt"
)

// InstructionLength is the length of an eBPF instruction in bytes
const InstructionLength int = 8

// getDestinationRegister extracts the destination register
// The four least significant bits of the register byte field represent the destination register
// e.g. registerByteField = 0x26 -> 0x6 -> register 6
func getDestinationRegister(registerByteField byte) uint8 {
	return registerByteField & 0xF
}

// getSourceRegister extracts the source register
// The four most significant bits of the register byte field represent the source register
// e.g. registerByteField = 0x26 -> 0x2 -> register 2
func getSourceRegister(registerByteField byte) uint8 {
	return registerByteField >> 4
}

// GetInstructionClass extracts the instruction class
// The three least significant bits of the opcode represent the instruction class
// e.g. opcode = 0xbc = 0b10111100 -> 0b00000100 = 0x4 -> instruction class 4
func GetInstructionClass(opcode byte) uint8 {
	return opcode & 0x7
}

// CreateInstructionFromRawInstruction decodes a raw eBPF instruction into an Instruction struct
// The raw instruction is expected to be 8 bytes long
// The first byte is the opcode
// The second byte is the register byte field
// The third and fourth bytes are the offset field
// The fifth to eighth bytes are the immediate field
// error is returned if the raw instruction is not 8 bytes long
// Possible errors:
//   - length of instruction is not 8 bytes
//   - error when decoding offset and immediate fields
func CreateInstructionFromRawInstruction(rawInstruction []byte) (types.Instruction, error) {
	var instruction types.Instruction
	lengthOfInstruction := len(rawInstruction)

	if lengthOfInstruction != InstructionLength {
		return instruction, fmt.Errorf("instruction has to be %v bytes long instead of %v bytes",
			InstructionLength, lengthOfInstruction)
	}

	instruction.Opcode = rawInstruction[0]
	instruction.DestinationRegister = getDestinationRegister(rawInstruction[1])
	instruction.SourceRegister = getSourceRegister(rawInstruction[1])

	_, err := binary.Decode(rawInstruction[2:4], binary.LittleEndian, &instruction.Offset)
	if err != nil {
		return instruction, fmt.Errorf("error when decoding offset field: %v", err)
	}

	_, err = binary.Decode(rawInstruction[4:], binary.NativeEndian, &instruction.Immediate)
	if err != nil {
		return instruction, fmt.Errorf("error when decoding immediate field: %v", err)
	}
	return instruction, nil
}

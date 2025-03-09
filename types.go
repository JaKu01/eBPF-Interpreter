// Based on eBPF kernel documentation: https://www.kernel.org/doc/html/v5.17/bpf/instruction-set.html

package main

import "fmt"

// ApplicationState represents the state of the eBPF virtual machine
// It contains the registers, the instruction counter and the stack
// The registers are 64 bit wide, the stack's size is 512 bytes
type ApplicationState struct {
	Registers          [11]uint64
	InstructionCounter uint64
	Stack              [512]byte
}

// Instruction represents one 64 bit eBPF instruction
// It corresponds to 64 VirtualBits / 8 VirtualBytes
// 63              31          15     11    7       0
// +--------------+-----------+------+-----+--------+
// |   Immediate  |  Offset   |  Src | Dst | Opcode |
// +--------------+-----------+------+-----+--------+
type Instruction struct {
	Immediate           uint32
	Offset              uint16
	SourceRegister      uint8
	DestinationRegister uint8
	Opcode              uint8
}

func (inst Instruction) String() string {
	return fmt.Sprintf("Opcode: %#x\nDestination Register: %#x\nSource Register: %#x\nOffset: %#x\nImmediate: %#x\n",
		inst.Opcode, inst.DestinationRegister, inst.SourceRegister, inst.Offset, inst.Immediate)
}

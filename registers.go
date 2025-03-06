package main

type VirtualBit uint8

type VirtualByte [8]VirtualBit

type Register [64]VirtualBit

type ApplicationState struct {
	Registers          [11]Register
	InstructionCounter uint64
	Stack              [512]VirtualByte
}

type Instruction struct {
	Immediate           [32]VirtualBit
	Offset              [16]VirtualBit
	SourceRegister      [4]VirtualBit
	DestinationRegister [4]VirtualBit
	Opcode              [8]VirtualBit
}

package main

import "testing"

func TestDetDestinationRegister(t *testing.T) {
	registerByteField := byte(0x26)
	result := getDestinationRegister(registerByteField)
	if result != 0x6 {
		t.Errorf("Expected 0x6, got %x", result)
	}
}

func TestGetSourceRegister(t *testing.T) {
	registerByteField := byte(0x26)
	result := getSourceRegister(registerByteField)
	if result != 0x2 {
		t.Errorf("Expected 0x2, got %x", result)
	}
}

func TestParseInstruction(t *testing.T) {
	//b7 01 00 00 b8 0b 00 00
	rawInstruction := []byte{0xb7, 0x21, 0x34, 0x58, 0xb8, 0x2b, 0x78, 0xa7}
	result, err := ParseInstruction(rawInstruction)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	expected := Instruction{
		Immediate:           0xa7782bb8,
		Offset:              0x5834,
		SourceRegister:      0x2,
		DestinationRegister: 0x1,
		Opcode:              0xb7,
	}
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

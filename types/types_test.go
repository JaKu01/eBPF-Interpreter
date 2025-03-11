package types

import "testing"

func TestInstruction_String(t *testing.T) {
	inst := Instruction{
		Immediate:           0x12345678,
		Offset:              0x1234,
		SourceRegister:      0x1,
		DestinationRegister: 0x2,
		Opcode:              0xb7,
	}
	expected := "Opcode: 0xb7\n Destination Register: 0x2\nSource Register: 0x1\nOffset: 0x1234\nImmediate: 0x12345678\n"
	result := inst.String()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

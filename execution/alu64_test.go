package execution

import (
	"eBPF-Interpreter/types"
	"testing"
)

func TestAddImmediate(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x5678

	const opcode uint8 = 0x07 | 0x0 | 0x0

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x1234,
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0x68AC
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %x, got %x", expected, applicationState.Registers[1])
	}
}

func TestSubImmediate(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x5678

	const opcode uint8 = 0x07 | 0x0 | 0x10

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x1234,
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0x4444
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %x, got %x", expected, applicationState.Registers[1])
	}
}

func TestMulImmediate(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x5678

	const opcode uint8 = 0x07 | 0x0 | 0x20

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x1234,
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0x6260060
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %x, got %x", expected, applicationState.Registers[1])
	}
}

func TestDiffImmediate(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x1000

	const opcode uint8 = 0x07 | 0x0 | 0x30

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x100,
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0x10
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %x, got %x", expected, applicationState.Registers[1])
	}
}

func TestLogicOrImmediate(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0b10101001

	const opcode uint8 = 0x07 | 0x0 | 0x40

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0b11000100,
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0b11101101
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %b, got %b", expected, applicationState.Registers[1])
	}
}

func TestLogicAndImmediate(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0b10101001

	const opcode uint8 = 0x07 | 0x0 | 0x50

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0b11000100,
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0b10000000
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %b, got %b", expected, applicationState.Registers[1])
	}
}

func TestLeftShiftImmediate(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0b10101001

	const opcode uint8 = 0x07 | 0x0 | 0x60

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x3,
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0b10101001000
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %b, got %b", expected, applicationState.Registers[1])
	}
}

func TestRightShiftImmediate(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0b10101001

	const opcode uint8 = 0x07 | 0x0 | 0x70

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x4,
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0b1010
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %b, got %b", expected, applicationState.Registers[1])
	}
}

func TestNegateImmediate(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x1234

	const opcode uint8 = 0x07 | 0x0 | 0x80

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0b1001011,
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0b11111111111111111111111110110100
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %b, got %b", expected, applicationState.Registers[1])
	}
}

func TestModImmediate(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x12345

	const opcode uint8 = 0x07 | 0x0 | 0x90

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x6,
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0x3
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %x, got %x", expected, applicationState.Registers[1])
	}
}

func TestXorImmediate(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0xfffffffffffffff0 | 0b1010

	const opcode uint8 = 0x07 | 0x0 | 0xa0

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0b0110,
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0xfffffffffffffff0 | 0b1100
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %x, got %x", expected, applicationState.Registers[1])
	}
}

func TestMoveImmediate(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x1234

	const opcode uint8 = 0x07 | 0x0 | 0xb0

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x512347f,
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0x512347f
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %x, got %x", expected, applicationState.Registers[1])
	}
}

func TestArithmeticShiftLeadingOneRightImmediate(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0b1000000000000000000000000000000000000000000000000000000000000101

	const opcode uint8 = 0x07 | 0x0 | 0xc0

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x2,
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0b1110000000000000000000000000000000000000000000000000000000000001
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %b, got %b", expected, applicationState.Registers[1])
	}
}

func TestArithmeticShiftLeadingZeroRightImmediate(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0b0000000000000000000000000000000000000000000000000000000000000101

	const opcode uint8 = 0x07 | 0x0 | 0xc0

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x2,
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0b0000000000000000000000000000000000000000000000000000000000000001
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %b, got %b", expected, applicationState.Registers[1])
	}
}

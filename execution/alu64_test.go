package execution

import (
	"eBPF-Interpreter/types"
	"testing"
)

// Tests for the ALU64 functions that use immediate values

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

func TestToLittleEndian16BitWidth(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x1234567890abcdef

	const opcode uint8 = 0x07 | 0x0 | 0xd0

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x10, // 16 bit width
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0x000000000000cdef // Assuming host byte order is little endian
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %#x, got %#x", expected, applicationState.Registers[1])
	}
}

func TestToLittleEndian32BitWidth(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x1234567890abcdef

	const opcode uint8 = 0x07 | 0x0 | 0xd0

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x20, // 32 bit width
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0x0000000090abcdef // Assuming host byte order is little endian
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %#x, got %#x", expected, applicationState.Registers[1])
	}
}

func TestToLittleEndian64BitWidth(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x1234567890abcdef

	const opcode uint8 = 0x07 | 0x0 | 0xd0

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x40, // 64 bit width
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0x1234567890abcdef // Assuming host byte order is little endian
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %#x, got %#x", expected, applicationState.Registers[1])
	}
}

func TestToBigEndian16BitWidth(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x1234567890abcdef

	const opcode uint8 = 0x07 | 0x08 | 0xd0

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x10, // 16 bit width
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0xefcd000000000000 // Assuming host byte order is little endian
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %#x, got %#x", expected, applicationState.Registers[1])
	}
}

func TestToBigEndian32BitWidth(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x1234567890abcdef

	const opcode uint8 = 0x07 | 0x08 | 0xd0

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x20, // 32 bit width
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0xefcdab9000000000 // Assuming host byte order is little endian
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %#x, got %#x", expected, applicationState.Registers[1])
	}
}

func TestToBigEndian64BitWidth(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x1234567890abcdef

	const opcode uint8 = 0x07 | 0x08 | 0xd0

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x40, // 64 bit width
		SourceRegister:      0x0,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0xefcdab9078563412 // Assuming host byte order is little endian
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %#x, got %#x", expected, applicationState.Registers[1])
	}
}

// Tests for the ALU64 functions that use register values

func TestAddSourceRegister(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x5678 // destination
	applicationState.Registers[2] = 0x1234 // source

	const opcode uint8 = 0x07 | 0x08 | 0x0

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x0,
		SourceRegister:      0x2,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0x68AC
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %x, got %x", expected, applicationState.Registers[1])
	}
}

func TestSubSourceRegister(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x5678 // destination
	applicationState.Registers[2] = 0x1234 // source

	const opcode uint8 = 0x07 | 0x08 | 0x10

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x0,
		SourceRegister:      0x2,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0x4444
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %x, got %x", expected, applicationState.Registers[1])
	}
}

func TestMulSourceRegister(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x5678 // destination
	applicationState.Registers[2] = 0x1234 // source

	const opcode uint8 = 0x07 | 0x08 | 0x20

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x0,
		SourceRegister:      0x2,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0x6260060
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %x, got %x", expected, applicationState.Registers[1])
	}
}

func TestDiffSourceRegister(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x1000 // destination
	applicationState.Registers[2] = 0x100  // source

	const opcode uint8 = 0x07 | 0x08 | 0x30

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x0,
		SourceRegister:      0x2,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0x10
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %x, got %x", expected, applicationState.Registers[1])
	}
}

func TestLogicOrSourceRegister(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0b10101001 // destination
	applicationState.Registers[2] = 0b11000100 // source

	const opcode uint8 = 0x07 | 0x08 | 0x40

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x0,
		SourceRegister:      0x2,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0b11101101
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %b, got %b", expected, applicationState.Registers[1])
	}
}

func TestLogicAndSourceRegister(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0b10101001 // destination
	applicationState.Registers[2] = 0b11000100 // source

	const opcode uint8 = 0x07 | 0x08 | 0x50

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x0,
		SourceRegister:      0x2,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0b10000000
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %b, got %b", expected, applicationState.Registers[1])
	}
}

func TestLeftShiftSourceRegister(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0b10101001
	applicationState.Registers[2] = 0x3

	const opcode uint8 = 0x07 | 0x08 | 0x60

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x0,
		SourceRegister:      0x2,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0b10101001000
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %b, got %b", expected, applicationState.Registers[1])
	}
}

func TestRightShiftSourceRegister(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0b10101001
	applicationState.Registers[2] = 0x4

	const opcode uint8 = 0x07 | 0x08 | 0x70

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x0,
		SourceRegister:      0x2,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0b1010
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %b, got %b", expected, applicationState.Registers[1])
	}
}

func TestNegateSourceRegister(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x1234
	applicationState.Registers[2] = 0b1001011

	const opcode uint8 = 0x07 | 0x08 | 0x80

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x0,
		SourceRegister:      0x2,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0xffffffffffffff00 | 0b10110100
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %b, got %b", expected, applicationState.Registers[1])
	}
}

func TestModSourceRegister(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x12345
	applicationState.Registers[2] = 0x6

	const opcode uint8 = 0x07 | 0x08 | 0x90

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x0,
		SourceRegister:      0x2,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0x3
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %x, got %x", expected, applicationState.Registers[1])
	}
}

func TestXorSourceRegister(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0xfffffffffffffff0 | 0b1010
	applicationState.Registers[2] = 0b0110

	const opcode uint8 = 0x07 | 0x08 | 0xa0

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x0,
		SourceRegister:      0x2,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0xfffffffffffffff0 | 0b1100
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %x, got %x", expected, applicationState.Registers[1])
	}
}

func TestMoveSourceRegister(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0x1234
	applicationState.Registers[2] = 0x512347f

	const opcode uint8 = 0x07 | 0x08 | 0xb0

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x0,
		SourceRegister:      0x2,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0x512347f
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %x, got %x", expected, applicationState.Registers[1])
	}
}

func TestArithmeticShiftLeadingOneRightSourceRegister(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0b1000000000000000000000000000000000000000000000000000000000000101
	applicationState.Registers[2] = 0x2

	const opcode uint8 = 0x07 | 0x08 | 0xc0

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x0,
		SourceRegister:      0x2,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0b1110000000000000000000000000000000000000000000000000000000000001
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %b, got %b", expected, applicationState.Registers[1])
	}
}

func TestArithmeticShiftLeadingZeroRightSourceRegister(t *testing.T) {
	applicationState = types.ApplicationState{}
	applicationState.Registers[1] = 0b0000000000000000000000000000000000000000000000000000000000000101
	applicationState.Registers[2] = 0x2

	const opcode uint8 = 0x07 | 0x08 | 0xc0

	instruction := types.Instruction{
		Opcode:              opcode,
		DestinationRegister: 0x1,
		Immediate:           0x0,
		SourceRegister:      0x2,
		Offset:              0x0,
	}

	ExecuteProgram([]types.Instruction{instruction})

	const expected uint64 = 0b0000000000000000000000000000000000000000000000000000000000000001
	result := applicationState.Registers[1]
	if result != expected {
		t.Errorf("Expected %b, got %b", expected, applicationState.Registers[1])
	}
}

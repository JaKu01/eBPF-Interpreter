package main

import "testing"

func TestPrintVirtualByte(t *testing.T) {
	inputVirtualBytes := [5]VirtualByte{
		{0, 1, 1, 1, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 1},
	}

	expectedValues := [5]string{
		"0111 0101",
		"0000 0000",
		"1111 1111",
		"1111 0000",
		"0000 1111",
	}

	var printedValues [5]string

	for idx, inputVirtualByte := range inputVirtualBytes {
		printedValues[idx] = inputVirtualByte.String()
	}

	for idx, virtualByteStrRepresentation := range printedValues {
		if virtualByteStrRepresentation != expectedValues[idx] {
			t.Errorf("Expected %s, got %s at index %v", expectedValues[idx], virtualByteStrRepresentation, idx)
		}
	}

}

package main

import "testing"

func TestExtractBits(t *testing.T) {
	var inputBytes = [5]byte{0b01110101, 0b00000000, 0b11111111, 0b11110000, 0b00001111}

	expectedValues := [5]VirtualByte{
		{0, 1, 1, 1, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 1},
	}

	var extractedVirtualBytes [5]VirtualByte

	for idx, byteInput := range inputBytes {
		extractedVirtualBytes[idx] = ExtractBits(byteInput)
	}

	for i, virtualByte := range extractedVirtualBytes {
		for j, num := range virtualByte {
			if num != expectedValues[i][j] {
				t.Errorf("Expected %v got %v at VirtualByte %v, at VirtualBit position %v", expectedValues[i][j], num, i, j)
			}
		}
	}
}

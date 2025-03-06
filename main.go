package main

import (
	"fmt"
)

func PrintVirtualByte(virtualByte VirtualByte) {
	fmt.Printf("%b%b%b%b %b%b%b%b", virtualByte[0], virtualByte[1], virtualByte[2], virtualByte[3], virtualByte[4], virtualByte[5], virtualByte[6], virtualByte[7])
}

func ExtractBits(byteToExtract byte) VirtualByte {

	var virtualByte VirtualByte

	for i := 7; i >= 0; i-- {
		extractedBit := (byteToExtract >> i) & 1
		virtualByte[7-i] = VirtualBit(extractedBit)
	}

	return virtualByte

}

func main() {

	bytes := []byte{0b10010110}

	virtualByte := ExtractBits(bytes[0])

	PrintVirtualByte(virtualByte)

}

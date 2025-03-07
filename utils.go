package main

func ExtractBits(byteToExtract byte) VirtualByte {
	var virtualByte VirtualByte

	for i := 7; i >= 0; i-- {
		extractedBit := (byteToExtract >> i) & 1
		virtualByte[7-i] = VirtualBit(extractedBit)
	}

	return virtualByte
}

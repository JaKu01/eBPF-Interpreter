package main

import "fmt"

func main() {

	bytes := []byte{0b10010110}

	virtualByte := ExtractBits(bytes[0])

	fmt.Println(virtualByte.String())

}

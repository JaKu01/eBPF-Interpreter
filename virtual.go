package main

import "fmt"

type VirtualBit uint8

type VirtualByte [8]VirtualBit

func (virtualByte *VirtualByte) String() string {
	return fmt.Sprintf("%b%b%b%b %b%b%b%b", virtualByte[0], virtualByte[1], virtualByte[2],
		virtualByte[3], virtualByte[4], virtualByte[5], virtualByte[6], virtualByte[7])
}

package main

import (
	"fmt"
	"log"
)

func main() {
	sectionContent, err := LoadInstructionsFromFile("./analysis/hello.bpf.o", "tc/egress")
	if err != nil {
		fmt.Println(err)
		return
	}

	var instructions []Instruction
	for i := 0; i < len(sectionContent); i += InstructionLength {
		result, err := ParseInstruction(sectionContent[i : i+InstructionLength])
		if err != nil {
			log.Fatalf("Could not decode instruction: %v", err)
		}
		instructions = append(instructions, result)
	}

	fmt.Println(instructions)
}

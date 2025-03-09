package main

import (
	"eBPF-Interpreter/cli"
	"flag"
	"fmt"
	"log"
)

func main() {
	args, err := cli.SetupCli()
	if err != nil {
		flag.Usage()
		return
	}

	sectionContent, err := LoadInstructionsFromFile(args.Filename, args.SectionName)
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

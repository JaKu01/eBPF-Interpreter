package main

import (
	"eBPF-Interpreter/cli"
	"eBPF-Interpreter/execution"
	"eBPF-Interpreter/parse"
	"eBPF-Interpreter/types"
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

	var instructions []types.Instruction
	for i := 0; i < len(sectionContent); i += parse.InstructionLength {
		result, err := parse.CreateInstructionFromRawInstruction(sectionContent[i : i+parse.InstructionLength])
		if err != nil {
			log.Fatalf("Could not decode instruction: %v", err)
		}
		instructions = append(instructions, result)
	}

	execution.ExecuteProgram(instructions)
}

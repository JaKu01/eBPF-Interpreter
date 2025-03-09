package cli

import (
	"flag"
	"fmt"
)

type InputArgs struct {
	Filename    string
	SectionName string
}

func SetupCli() (InputArgs, error) {
	var inputArgs InputArgs

	flag.StringVar(&inputArgs.Filename, "filename", "", "The filename of the BPF object file")
	flag.StringVar(&inputArgs.Filename, "f", "", "The filename of the BPF object file")
	flag.StringVar(&inputArgs.SectionName, "section-name", "", "The name of the section to extract")
	flag.StringVar(&inputArgs.SectionName, "s", "", "The name of the section to extract")

	// provide custom usage message
	flag.Usage = func() {
		fmt.Println("Usage: eBPF-Interpreter [-f <path> | --filename <path>] [-s <section> | --section-name <section>]")
		fmt.Println("Flags:")
		fmt.Println("  -f, --filename string       The filename of the BPF object file")
		fmt.Println("  -s, --section-name string   The name of the section to extract")
		fmt.Println("  -h, --help                  Show help information")
	}

	flag.Parse()

	// Check if filename and section-name are provided
	if inputArgs.Filename == "" || inputArgs.SectionName == "" {
		return inputArgs, fmt.Errorf("filename and section-name are required")
	}

	return inputArgs, nil
}

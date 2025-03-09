package main

import (
	"debug/elf"
	"fmt"
)

// LoadInstructionsFromFile is a generic helper function to load instructions from an ELF file
// It takes the filename and the section name as input and returns the section content as a byte slice
func LoadInstructionsFromFile(filename string, sectionName string) ([]byte, error) {
	elfFile, err := elf.Open(filename)
	if err != nil {
		return nil, err
	}

	section := elfFile.Section(sectionName)
	if section == nil {
		return nil, fmt.Errorf("could not find section: %v", sectionName)
	}

	sectionContent, err := section.Data()
	return sectionContent, err
}

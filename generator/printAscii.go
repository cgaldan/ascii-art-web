package generator

import (
	"fmt"
	"os"
	"strings"
)

const blockSize = 8

func PrintAsciiArt(inputString string, banner string) (string, error) {
	asciiBlocks := ReadBanner(banner)

	var outputLines []string
	inputLines := strings.Split(inputString, "\r\n")

	// Iterate over each input line and build ASCII art representation
	for _, line := range inputLines {
		if line == "" {
			outputLines = append(outputLines, "")
			continue
		}

		tempOutputLines := make([]string, blockSize) // Temp slice for each line of text

		for _, char := range line {
			asciiIndex := int(char) - 32
			if asciiIndex >= 0 && asciiIndex < len(asciiBlocks) {
				art := asciiBlocks[asciiIndex]
				for i := 0; i < blockSize; i++ {
					tempOutputLines[i] += art[i]
				}
			} else {
				return "", fmt.Errorf("character '%c' not found in banner", char)
			}
		}

		// Append the constructed temp lines to the final output
		outputLines = append(outputLines, tempOutputLines...)
	}
	output := strings.Join(outputLines, "\n")
	return output, nil
}

func ReadBanner(banner string) [][]string {
	content, err := os.ReadFile(fmt.Sprintf("banners/%s.txt", banner))
	if err != nil {
		fmt.Printf("Failed to read banner file: %v", err)
	}
	return CreateLetterBlocks(content, banner)
}

func CreateLetterBlocks(content []byte, banner string) [][]string {
	var lines []string

	if banner == "thinkertoy" {
		lines = strings.Split(string(content), "\r\n")
	} else {
		lines = strings.Split(string(content), "\n")
	}

	var blocks [][]string
	var currentBlock []string

	// Iterate through lines and group them into blocks
	for index, line := range lines {
		if index%9 == 0 {
			continue
		}
		currentBlock = append(currentBlock, line)
		if len(currentBlock) == blockSize {
			blocks = append(blocks, currentBlock)
			currentBlock = []string{} // Reset for the next block
		}
	}
	return blocks
}

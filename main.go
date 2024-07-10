package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go-reloaded/handlers"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("need three arguments: go run main.go input.txt output.txt")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)
	defer writer.Flush()

	for scanner.Scan() {

		line := scanner.Text()
		words := strings.Split(line, " ")
		// logic addressing cases of (hex), (bin), (up), (low), (cap)
		for i, word := range words {
			switch word {
			case "(hex)":
				hexStr := words[i-1]
				words[i-1] = handlers.ChangeHexToDecimal(hexStr)
				words = append(words[:i], words[i+1:]...)
			case "(bin)":
				binStr := words[i-1]
				words[i-1] = handlers.ChangeBinToDecimal(binStr)
				words = append(words[:i], words[i+1:]...)
			case "(up)":
				words[i-1] = handlers.ConvertUpToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
			case "(low)":
				words[i-1] = handlers.ConvertLowToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
			case "(cap)":
				words[i-1] = handlers.ConvertCapToCapitalized(words[i-1])
				words = append(words[:i], words[i+1:]...)
			}
		}
		// adressing instances of (up, <number>) et al. within the sample.txt file

		for i, word := range words {
			if word == "(up," {
				num := handlers.ExtractNumber(words[i+1])
				for j := 1; j <= num; j++ {
					if i-j >= 0 {
						words[i-j] = handlers.ConvertUpToUpper(words[i-j])
					}
				}
				words = append(words[:i], words[i+2:]...)
			} else if word == "(low," {
				num := handlers.ExtractNumber(words[i+1])
				for j := 1; j <= num; j++ {
					if i-j >= 0 {
						words[i-j] = handlers.ConvertLowToLower(words[i-j])
					}
				}
				words = append(words[:i], words[i+2:]...)
			} else if word == "(cap," {
				num := handlers.ExtractNumber(words[i+1])
				for j := 1; j <= num; j++ {
					if i-j >= 0 {
						words[i-j] = handlers.ConvertCapToCapitalized(words[i-j])
					}
				}
				words = append(words[:i], words[i+2:]...)
			}
		}
		handlers.AToAn(words)
		words = handlers.MovePunctuation(words)
		words = handlers.SplitStringWithPunctuation(words)
		fmt.Fprintln(writer, strings.Join(handlers.PunctuationCorrecting(words), " "))

	}
}

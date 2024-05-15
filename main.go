package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"wordify/transformers"
)

type TransformFunc func(string) string

type TransformOperation struct {
	Transform TransformFunc
}

func getTransformations() map[string]TransformOperation {
	transformations := map[string]TransformOperation{
		"(up)":  {strings.ToUpper},
		"(low)": {strings.ToLower},
		"(cap)": {transformers.ToTitle},
		"(hex)": {transformers.HexadecimalToInt},
		"(bin)": {transformers.BinaryToInt},
	}
	return transformations
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: program input_file output_file")
		os.Exit(1)
	}

	// Read input file
	inputText, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	// Split input text into words
	words := strings.Fields(string(inputText))
	transformations := getTransformations()

	// Process transformations
	for i := 0; i < len(words); i++ {
		word := words[i]
		IsDerivativeN := word == "(up," || word == "(low," || word == "(cap,"
		if op, exists := transformations[word]; exists {
			if i > 0 { // Ensure there is a previous word to transform
				words[i-1] = op.Transform(words[i-1])
				words = append(words[:i], words[i+1:]...) // Remove the transformation word
				i--                                       // Adjust index to account for removed word
			}
		} else if strings.HasSuffix(word, ",") && i+1 < len(words) && IsDerivativeN {
			// Handle commands like (up, 3)
			transformType := strings.TrimRight(word, ",")
			b := strings.TrimSuffix(words[i+1], ")")
			number, err := strconv.Atoi(b)
			if err != nil {
				fmt.Println("Error! value provided is not a valid number")
			}

			if number > i {
				fmt.Printf("Error: number greater than length of string upto %d. Expecting string length %d\n", i, i)
				break
			} else if number < 0 {
				fmt.Printf("Error: number (%d) entered is negative\n", number)
				break
			}

			transform, err := getCustomTransformFunc(transformType)
			if err != nil {
				fmt.Printf("Error getting custom transform function: %s\n", err)
				continue
			}

			for j := 1; j <= number && (i-j >= 0); j++ {
				words[i-j] = transform(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			i -= 2 // Adjust index for removed words
		}
	}

	words = transformers.CorrectArticles(words)
	words = transformers.FormatPunctuations(words)
	finalText := strings.Join(words, " ") + "\n"

	// Write the modified text to the output file
	if err := os.WriteFile(args[1], []byte(finalText), 0o644); err != nil {
		fmt.Printf("Error writing file: %s\n", err)
		os.Exit(1)
	}
}

func getCustomTransformFunc(transformType string) (TransformFunc, error) {
	switch transformType {
	case "(up":
		return strings.ToUpper, nil
	case "(low":
		return strings.ToLower, nil
	case "(cap":
		return transformers.ToTitle, nil
	default:
		return nil, fmt.Errorf("unknown custom transformation type: %s", transformType)
	}
}

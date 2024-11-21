package funcs

import (
	"fmt"
	"strings"
)

// Function to print ASCII art to the console
func PrintAsciiArt(sentences []string, textFile []string) (string, error) {
	var result strings.Builder // Use strings.Builder to build the ASCII art output

	for _, word := range sentences {
		if word == "" {
			// Add a newline for blank words
			result.WriteString("\n")
			continue
		}

		// Generate ASCII art for each line of height (assume height is 8)
		for h := 1; h < 9; h++ { // ASCII art character height is 8
			for _, char := range word {
				if char == '\n' {
					// If the character is a newline, just add a newline in the output
					result.WriteString("\n")
					continue
				}

				charIndex := int(char) - 32  // Calculate the index for the character in the ASCII art file
				startLine := charIndex*9 + h // Calculate the line index for the current character and height

				// Ensure the line index is within bounds
				if startLine < 0 || startLine >= len(textFile) {
					return "", fmt.Errorf("character %c not supported in banner", char)
				}

				result.WriteString(textFile[startLine]) // Add the line for the character
			}
			result.WriteString("\n") // New line after each line of ASCII art height
		}
		result.WriteString("\n") // New line after each word
	}

	return result.String(), nil
}

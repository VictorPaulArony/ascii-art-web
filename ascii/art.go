package ascii

import (
	"strings"
)

// DisplayText generates ASCII art from the provided text and content lines
func DisplayText(input string, contentLines []string) string {
	if input == "" {
		return ""
	}

	if input == "\\n" || input == "\n" {
		return "\n"
	}
	// make newline and tab printable in the terminal output
	input = strings.ReplaceAll(input, "\n", "\\n")
	input = strings.ReplaceAll(input, "\\t", "\t")

	wordslice := strings.Split(input, "\\n")
	var result strings.Builder

	for _, word := range wordslice {
		if word == "" {
			result.WriteString("\n")
		} else {
			if English(word) {
				result.WriteString(PrintWord(word, contentLines))
			} else {
				result.WriteString("Invalid input: not accepted\n")
				// Optionally continue processing other words
			}
		}
	}
	return result.String()
}

// English checks if a word contains only English alphabets
func English(words string) bool {
	for _, word := range words {
		if word < 32 || word > 126 {
			return false
		}
	}
	return true
}

// PrintWord generates ASCII art for a word using the content lines
func PrintWord(word string, contentLines []string) string {
	linesOfSlice := make([]string, 9)

	for _, v := range word {
		for i := 1; i <= 9; i++ {
			linesOfSlice[i-1] += contentLines[int(v-32)*9+i]
		}
	}
	return strings.Join(linesOfSlice, "\n") + "\n"
}

package ascii

import (
	"bufio"
	"os"
	"strings"
)

// DisplayText displays the provided text along with content lines
func DisplayText(input, filename string) (string, error) {
	var res string
	if input == "" {
		return "", nil
	}

	if input == "\\n" || input == "\n" {
		return "\n", nil
	}

	// make newline and tab printable in the terminal output
	input = strings.ReplaceAll(input, "\n", "\\n")
	input = strings.ReplaceAll(input, "\t", "\\t")

	contentLines, err := readFile(filename)
	if err != nil {
		return "", err
	}

	wordslice := strings.Split(input, "\\n")

	for _, word := range wordslice {
		if word == "" {
			res += "\n"
		} else {
			if English(word) {
				res += PrintWord(word, contentLines) + "\n"
			} else {
				res += "400 - Bad Request Invalid input: not accepted\n"
			}
		}
	}
	return res, nil
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

// PrintWord prints a word if it exists in the content lines
func PrintWord(word string, contentLines []string) string {
	linesOfSlice := make([]string, 9)
	for _, v := range word {
		for i := 0; i < 9; i++ {
			linesOfSlice[i] += contentLines[int(v-32)*9+i]
		}
	}
	return strings.Join(linesOfSlice, "\n")
}

// readFile reads the banner file and returns its content as a slice of strings
func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

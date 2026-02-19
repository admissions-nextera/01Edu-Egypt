package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 || os.Args[2] == "" {
		return
	}

	pattern := os.Args[1]
	text := os.Args[2]

	// Validate pattern format
	if len(pattern) < 3 || pattern[0] != '(' || pattern[len(pattern)-1] != ')' {
		return
	}

	// Extract content between parentheses
	content := pattern[1 : len(pattern)-1]
	if content == "" {
		return
	}

	// Split by |
	terms := splitPipe(content)

	// Extract words from text
	words := getWords(text)

	// Find and print all matches
	matchNum := 0
	for _, word := range words {
		count := countMatches(word, terms)
		for i := 0; i < count; i++ {
			matchNum++
			fmt.Printf("%d: %s\n", matchNum, word)
		}
	}
}

func splitPipe(s string) []string {
	var result []string
	current := ""

	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			if current != "" {
				result = append(result, current)
				current = ""
			}
		} else {
			current += string(s[i])
		}
	}

	if current != "" {
		result = append(result, current)
	}

	return result
}

func getWords(text string) []string {
	var words []string
	word := ""

	for i := 0; i < len(text); i++ {
		ch := text[i]
		// Include letters and various apostrophe characters
		if isLetter(ch) || ch == '\'' {
			word += string(ch)
		} else {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		}
	}

	if word != "" {
		words = append(words, word)
	}

	return words
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func countMatches(word string, terms []string) int {
	total := 0
	i := 0

	for i < len(word) {
		matched := false

		for _, term := range terms {
			if i+len(term) <= len(word) && matchesAt(word, term, i) {
				total++
				i += len(term)
				matched = true
				break
			}
		}

		if !matched {
			i++
		}
	}

	return total
}

func matchesAt(word, term string, pos int) bool {
	for j := 0; j < len(term); j++ {
		if word[pos+j] != term[j] {
			return false
		}
	}
	return true
}

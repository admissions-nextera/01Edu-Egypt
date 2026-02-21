package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("USAGE: go run . <input-file> <output-file>")
		return
	}
	content, err := readFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	final := processText(content)
	err = writeFile(os.Args[2], final)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
}

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Error: %v", err)
	}
	return string(data), nil
}

func writeFile(filename string, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func processText(text string) string {
	result := strings.Fields(text)
	result = processModifiers(result)
	result = fixPunctuation(result)
	result = fixQuotes(result)
	result = fixArticles(result)
	return strings.Join(result, " ")
}

func hexToDecimal(s string) string {
	dec, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return s
	}
	return strconv.Itoa(int(dec))
}

func binToDecimal(s string) string {
	dec, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return s
	}
	return strconv.Itoa(int(dec))
}

func processModifiers(words []string) []string {
	result := []string{}

	for i := 0; i < len(words); i++ {
		word := words[i]
		switch word {
		case "(hex)":
			if len(result) > 0 {
				result[len(result)-1] = hexToDecimal(result[len(result)-1])
			}
		case "(bin)":
			if len(result) > 0 {
				result[len(result)-1] = binToDecimal(result[len(result)-1])
			}
		case "(up)":
			if len(result) > 0 {
				result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			}
		case "(low)":
			if len(result) > 0 {
				result[len(result)-1] = strings.ToLower(result[len(result)-1])
			}
		case "(cap)":
			if len(result) > 0 {
				result[len(result)-1] = capitalize(result[len(result)-1])
			}
		case "(up,":
			if len(result) > 0 && i+1 < len(words) {
				nextWord := words[i+1]

				if before, ok := strings.CutSuffix(nextWord, ")"); ok {
					n, err := strconv.Atoi(before)
					if err == nil {
						if n > len(result) {
							n = len(result)
						}
						for j := 0; j < n; j++ {
							idx := len(result) - 1 - j
							result[idx] = strings.ToUpper(result[idx])
						}
						i++
					}
				}
			}
		case "(low,":
			if len(result) > 0 && i+1 < len(words) {
				nextWord := words[i+1]

				if before, ok := strings.CutSuffix(nextWord, ")"); ok {
					n, err := strconv.Atoi(before)
					if err == nil {
						if n > len(result) {
							n = len(result)
						}
						for j := 0; j < n; j++ {
							idx := len(result) - 1 - j
							result[idx] = strings.ToLower(result[idx])
						}
						i++
					}
				}
			}
		case "(cap,":
			if len(result) > 0 && i+1 < len(words) {
				nextWord := words[i+1]

				if before, ok := strings.CutSuffix(nextWord, ")"); ok {
					n, err := strconv.Atoi(before)
					if err == nil {
						if n > len(result) {
							n = len(result)
						}
						for j := 0; j < n; j++ {
							idx := len(result) - 1 - j
							result[idx] = capitalize(result[idx])
						}
						i++
					}
				}
			}
		default:
			result = append(result, word)
		}
	}

	return result
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func fixPunctuation(words []string) []string {
	var result []string

	for _, word := range words {
		if len(word) > 0 && isPunctuation(word[0]) {
			i := 0
			for i < len(word) && isPunctuation(word[i]) {
				i++
			}
			puncPart := word[:i]
			restPart := word[i:]

			if len(result) > 0 {
				result[len(result)-1] += puncPart
			} else {
				result = append(result, puncPart)
			}

			if restPart != "" {
				result = append(result, restPart)
			}
		} else {
			result = append(result, word)
		}
	}
	return result
}

func isPunctuation(r byte) bool {
	return r == '.' || r == '!' || r == '?' || r == ':' || r == ';' || r == ','
}

func fixQuotes(words []string) []string {
	var result []string
	isOpenQuote := false
	for _, word := range words {
		if word == "'" {
			if !isOpenQuote {
				result = append(result, word)
				isOpenQuote = true
			} else {
				if len(result) > 0 {
					result[len(result)-1] += "'"
				}
				isOpenQuote = false
			}
		} else {
			if isOpenQuote && len(result) > 0 && result[len(result)-1] == "'" {
				result[len(result)-1] += word
			} else {
				result = append(result, word)
			}
		}
	}
	return result
}

func fixArticles(words []string) []string {
	result := make([]string, len(words))
	copy(result, words)
	for i := 0; i < len(result)-1; i++ {
		nextWord := result[i+1]
		if result[i] == "a" || result[i] == "A" {
			if len(nextWord) > 0 {
				firstChar := rune(strings.ToLower(nextWord)[0])
				if isVowelOrH(firstChar) {
					if result[i] == "a" {
						result[i] = "an"
					} else {
						result[i] = "An"
					}
				}
			}
		}
	}

	return result
}

func isVowelOrH(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h'
}

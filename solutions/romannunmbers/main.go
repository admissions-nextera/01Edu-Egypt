package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := atoi(os.Args[1])
	if err != "" {
		fmt.Println(err)
		return
	}

	if num <= 0 || num > 3999 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	calculation, result := toRoman(num)
	fmt.Println(calculation)
	fmt.Println(result)
}

func toRoman(n int) (string, string) {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	calculations := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}

	var calculation []string
	var roman string

	for i := 0; i < len(values); i++ {
		for n >= values[i] {
			calculation = append(calculation, calculations[i])
			roman += symbols[i]
			n -= values[i]
		}
	}

	result := ""
	for i, calc := range calculation {
		if i > 0 {
			result += "+"
		}
		result += calc
	}

	return result, roman
}

func atoi(s string) (int, string) {
	if len(s) == 0 {
		return 0, "ERROR: cannot convert to roman digit"
	}

	result := 0

	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0, "ERROR: cannot convert to roman digit"
		}
		result = result*10 + int(ch-'0')
	}

	return result, ""
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	result, err := evaluateRPN(os.Args[1])
	if err {
		fmt.Println("Error")
	} else {
		fmt.Println(result)
	}
}

func evaluateRPN(expression string) (int, bool) {
	tokens := strings.Fields(expression)
	if len(tokens) == 0 {
		return 0, true
	}

	var stack []int

	for _, token := range tokens {
		if isOperator(token) {
			if len(stack) < 2 {
				return 0, true
			}

			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			result, err := calculate(a, b, token)
			if err {
				return 0, true
			}

			stack = append(stack, result)
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, true
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return 0, true
	}

	return stack[0], false
}

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/" || s == "%"
}

func calculate(a, b int, op string) (int, bool) {
	switch op {
	case "+":
		return a + b, false
	case "-":
		return a - b, false
	case "*":
		return a * b, false
	case "/":
		if b == 0 {
			return 0, true
		}
		return a / b, false
	case "%":
		if b == 0 {
			return 0, true
		}
		return a % b, false
	}
	return 0, true
}

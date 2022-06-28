package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	err = isEmpty(input)
	if err != nil {
		return "", fmt.Errorf("error whil calculating the sum: %w", err)
	}

	a, b, err := getOperands(input)
	if err != nil {
		return "", fmt.Errorf("error whil calculating the sum: %w", err)
	}

	return strconv.Itoa(a + b), nil
}

func isEmpty(str string) error {
	if len(str) == 0 {
		return errorEmptyInput
	}

	cnt := 0
	cnt += strings.Count(str, " ")
	cnt += strings.Count(str, "\t")
	cnt += strings.Count(str, "\n")
	if cnt == len(str) {
		return errorEmptyInput
	}

	return nil
}

func getOperands(str string) (a, b int, err error) {
	s := strings.Split(str, "+")
	if len(s) != 2 {
		return a, b, errorNotTwoOperands
	}

	a, err = convertToString(s[0])
	if err != nil {
		return 0, 0, fmt.Errorf("error while getting operands: %w", err)
	}

	b, err = convertToString(s[1])
	if err != nil {
		return 0, 0, fmt.Errorf("error while getting operands: %w", err)
	}

	return
}

func convertToString(str string) (int, error) {
	str = strings.TrimSpace(str)
	sign := 1
	if len(str) > 1 && str[0] == '-' {
		sign *= -1
		str = str[1:]
	}
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("error while converting operand to int: %w", err)
	}
	if sign == -1 {
		n = -n
	}
	return n, nil
}

func deleteWhiteSpaces(str string) string {
	return ""
}

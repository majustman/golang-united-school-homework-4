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
	if isEmpty(input) {
		return "", errorEmptyInput
	}

	a, b, err := getOperands(input)
	if err != nil {
		return "", fmt.Errorf("error while getting operands: %w", err)
	}

	return string(a + b), nil
}

func isEmpty(str string) bool {
	if len(str) == 0 {
		return true
	}

	cnt := 0
	cnt += strings.Count(str, " ")
	cnt += strings.Count(str, "\t")
	cnt += strings.Count(str, "\n")
	if cnt == len(str) {
		return true
	}

	return false
}

func getOperands(str string) (a, b int, err error) {
	s := strings.Split(str, "+")
	if len(s) != 2 {
		return a, b, errorNotTwoOperands
	}
	op1 := s[0]
	op2 := s[1]

	sign := 1
	if len(op1) > 1 && op1[0] == '-' {
		sign *= -1
		op1 = op1[1:]
	}
	a, err = strconv.Atoi(op1)
	if err != nil {
		return 0, 0, fmt.Errorf("error while converting the first operand to int: %w", err)
	}
	if sign == -1 {
		a = -a
	}

	sign = 1
	if len(op2) > 1 && op2[0] == '-' {
		sign = -sign
		op2 = op2[1:]
	}
	b, err = strconv.Atoi(op2)
	if err != nil {
		return 0, 0, fmt.Errorf("error while converting the second operand to int: %w", err)
	}
	if sign == -1 {
		b = -b
	}

	return
}

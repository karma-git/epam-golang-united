//use these errors as appropriate, wrapping them with fmt.Errorf function
// var (
// 	// Use when the input is empty, and input is considered empty if the string contains only whitespace
// 	errorEmptyInput = errors.New("input is empty")
// 	// Use when the expression has number of operands not equal to two
// 	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
// )

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/apaxa-go/eval"
)

var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func digitsCount(input string) (result int) {
	for _, i := range input {
		if unicode.IsDigit(i) {
			result++
		}
	}
	return
}

func IsNotEmpty(input string) (result bool) {
	if strings.TrimSpace(input) == "" {
		result = false
	} else {
		result = true
	}
	return
}

func removeWhitespace(input string) (result string) {
	var arr []string
	for _, i := range input {
		if !unicode.IsSpace(i) {
			arr = append(arr, string(i))
		}
	}
	result = strings.Join(arr[:], "")
	return
}

func StringSum(input string) (output string, err error) {
	input = removeWhitespace(input)
	if !IsNotEmpty(input) || digitsCount(input) == 0 {
		err = fmt.Errorf("%w", errorEmptyInput)
		return
	} else if digitsCount(input) != 2 {
		err = fmt.Errorf("%w", errorNotTwoOperands)
		return
	} else {
		expr, err := eval.ParseString(input, "")
		if err != nil {
			return "", err
		}
		r, err := expr.EvalToInterface(nil)
		if err != nil {
			return "", err
		}
		output = fmt.Sprint(r)
		return output, nil
	}
}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	r, err := StringSum(s)
	fmt.Println(r, err)
}

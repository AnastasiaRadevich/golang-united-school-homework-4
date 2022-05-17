package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type MyCustomError struct {
	Message string
}

func (m MyCustomError) Error() string {
	return m.Message
}

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	incorrectReg := `[^\-|^\+|^\d|^[:space:]]`
	numReg, _ := regexp.Compile(`\d+`)
	myError := MyCustomError{Message: "input contains characters, that are not numbers, +, - or whitespace"}
	if len(input) < 1 || input == " " {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	findNum := numReg.FindAllString(input, -1)
	if len(findNum) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	matched, _ := regexp.Match(incorrectReg, []byte(input))
	if matched {
		return "", fmt.Errorf("%s", myError.Error())
	}

	re, _ := regexp.Compile(`\-|\+|\d+`)
	res := re.FindAllString(input, -1)

	result := 0
	for i := 0; i < len(res)-1; i++ {
		x, _ := strconv.Atoi(res[i+1])
		if res[i] == "-" {
			result = result - x
			i++
			continue
		}
		if res[i] == "+" {
			result = result + x
			i++
			continue
		}
		z, _ := strconv.Atoi(res[i])
		result = result + z
	}
	return strconv.Itoa(result), nil
}

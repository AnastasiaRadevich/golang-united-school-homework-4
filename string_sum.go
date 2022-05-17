package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	numReg, _ := regexp.Compile(`\d+`)

	if len(input) < 1 || input == " " {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	findNum := numReg.FindAllString(input, -1)
	if len(findNum) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	incorrectReg, _ := regexp.Compile(`[^\-|^\+|^\d|^[:space:]]`)
	incorrectArg := incorrectReg.FindAllString(input, -1)
	if len(incorrectArg) > 0 {
		_, errArg1 := strconv.Atoi(incorrectArg[0])
		if errArg1 != nil {
			return "", fmt.Errorf("%w", errArg1.(*strconv.NumError))
		}
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

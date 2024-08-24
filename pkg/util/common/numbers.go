package common

import (
	"fmt"
	"strconv"
	"unicode"
)

func ToIntMust(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}
	return i
}

func ToIntOr(str string, or int) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return or
	}
	return i
}

func ExtractFirstInt(s string) (int, error) {
	var numberStr string
	for _, char := range s {
		if unicode.IsDigit(char) {
			numberStr += string(char)
		} else if len(numberStr) > 0 {
			break
		}
	}

	if numberStr != "" {
		return strconv.Atoi(numberStr)
	}
	return 0, fmt.Errorf("no digits found in string")
}

func ExtractFirstIntOr(s string, or int) int {
	i, err := ExtractFirstInt(s)
	if err != nil {
		return or
	}
	return i
}

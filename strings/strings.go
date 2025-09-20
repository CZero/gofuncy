package strings

import (
	"fmt"
	"strconv"
	"strings"
)

// SilentAtoi returns an int or panics
func SilentAtoi(input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Sprintf("SilentAtoi failed with %v as input", val))
	}
	return val
}

// ReverseString returns a reversed string
func ReverseString(word string) string {
	var reversed string
	for i := len(word) - 1; i >= 0; i-- {
		reversed = reversed + string(word[i])
	}
	return reversed
}

// RemoveSpaces removes all extra whitespace from a string
func RemoveSpaces(input string) string {
	return strings.Join(strings.Fields(input), " ")
}

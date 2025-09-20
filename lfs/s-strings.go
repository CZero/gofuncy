// lfs stands for Lib Functions Steph
package lfs

import (
	"strconv"

	"github.com/CZero/gofuncy/errors"
)

// SilentAtoi returns an int or panics
func SilentAtoi(input string) int {
	val, err := strconv.Atoi(input)
	errors.PanErr(err)
	return val
}

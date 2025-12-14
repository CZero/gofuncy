package digitsfuncs

import "math"

// FindMinMaxInxDigits returns the smalles and biggest numbers of x digits
func FindMinMaxInxDigits(digits int) (min, max int) {
	min = int(math.Pow(10, (float64(digits) - 1)))
	max = int(math.Pow(10, float64(digits))) - 1
	return min, max
}

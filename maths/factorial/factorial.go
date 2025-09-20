package factorial

// GetFactorial returns the factorials of a number. It returns 0 when 0 or smaller.
func GetFactorial(input int) int {
	if input <= 0 {
		return 0
	}
	result := 1
	for i := 1; i <= input; i++ {
		result *= (i)
	}
	return result
}

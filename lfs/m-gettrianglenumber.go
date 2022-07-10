package lfs

// Gettrianglenumber gets the n'th triangle number. The sequence of triangle numbers is generated by adding the natural numbers. So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28
func Gettrianglenumber(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}
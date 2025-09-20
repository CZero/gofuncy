package division

import "math"

/*
*
Getdivisors tries like in the linked example to divide all numbers under the squareroot. Each result is also
a divisor. We have to check if it's the same though: 100 / 10 = 10.
The second half works backwards: 100/1 = 100, 100/2 = 50
So we store those apart and in the end return those reversed.
https://www.geeksforgeeks.org/find-all-divisors-of-a-natural-number-set-2/
*
*/
func Getdivisors(i int) []int {
	var reverseme, result []int
	for n := 1; n <= int(math.Sqrt(float64(i))); n++ {
		if i%n == 0 {
			result = append(result, n)
			if i/n != n {
				reverseme = append(reverseme, i/n)
			}
		}
	}
	for m := len(reverseme) - 1; m >= 0; m-- {
		result = append(result, reverseme[m])
	}
	return result
}

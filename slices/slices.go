package slices

import "sort"

// SortIntSlice sorts an intslice
func sortIntSlice(input []int) []int {
	sort.Ints(input)
	return input
}

// SumSlice returns the sum of a slice
func SumSlice(inOrder []int) (sum int) {
	for _, item := range inOrder {
		sum += item
	}
	return sum
}

// ProductSlice returns the product of a slice
func ProductSlice(inOrder []int) (sum int) {
	for _, item := range inOrder {
		sum *= item
	}
	return sum
}

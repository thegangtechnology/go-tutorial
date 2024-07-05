package demo

import "fmt"

func FilterTakeOnlyEvenNumbers(numbers []int) []int {
	var result []int
	// some tutorial will use make([]int, 0)
	// there is a subtlety that we will not get into.
	for _, n := range numbers {
		if n%2 == 0 {
			result = append(result, n)
		}
	}
	return result
}

func Argmin(a []int) (int, error) {
	if len(a) == 0 {
		return 0, fmt.Errorf("slice is empty")
	}
	minSoFar := a[0]
	minIdx := 0
	for i, v := range a {
		if v < minSoFar {
			minSoFar = v
			minIdx = i
		}
	}
	return minIdx, nil
}

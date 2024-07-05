package exercise

import (
	"fmt"
	"sort"
)

// Instruction
// implement and write test both happy and bad cases.

func KthLargestNumber(numbers []int, k int) (int, error) {
	if k <= 0 || k > len(numbers) {
		return 0, fmt.Errorf("k should be positive and less than length of numbers")
	}
	sort.Ints(numbers)
	return numbers[len(numbers)-k], nil
}

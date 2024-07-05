package exercise

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthLargestNumber(t *testing.T) {

	t.Run("Happy Path", func(t *testing.T) {
		numbers := []int{3, 2, 1, 5, 6, 4}
		k := 2
		expected := 5
		result, err := KthLargestNumber(numbers, k)
		assert.Equal(t, expected, result)
		assert.Nil(t, err)
	})

	t.Run("Negative Index", func(t *testing.T) {
		numbers := []int{3, 2, 1, 5, 6, 4}
		k := -2
		_, err := KthLargestNumber(numbers, k)
		assert.NotNil(t, err)
	})

	t.Run("k beyond length", func(t *testing.T) {
		numbers := []int{3, 2, 1, 5, 6, 4}
		k := 100
		_, err := KthLargestNumber(numbers, k)
		assert.NotNil(t, err)
	})

}

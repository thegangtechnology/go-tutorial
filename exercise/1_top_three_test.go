package exercise

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// implement this too
func TestCalculateTopThree(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		tt, err := CalculateTopThree("wordcount.txt")
		assert.Nil(t, err)
		expected := TopThreeWords{
			First:  "the",
			Second: "a",
			Third:  "of",
		}
		assert.Equal(t, expected, tt)
	})
}

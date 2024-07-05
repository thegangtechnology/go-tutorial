package demo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// go test ./... -v
// printout is also pretty shitty

func TestIsPrime(t *testing.T) {
	// yes the go designer decide assert is not necessary in stdlib
	if !IsPrime(7) {
		t.Error("7 should be a prime number stdlib")
	}
	// we need testify/asssert package
	assert.True(t, IsPrime(7), "7 should be a prime number")
	assert.False(t, IsPrime(8), "8 should be a prime number")
	assert.False(t, IsPrime(0), "0 should be a prime number")
	assert.False(t, IsPrime(1), "1 should be a prime number")
}

func TestSumPrime(t *testing.T) {
	assert.Equal(t, 0, SumPrime(0))
	assert.Equal(t, 2, SumPrime(2))
	assert.Equal(t, 10, SumPrime(5))
	assert.Equal(t, 10, SumPrime(6))
}

func TestSumPrintWithError(t *testing.T) {

	t.Run("Happy Path", func(t *testing.T) {
		s, err := SumPrimeWithError(6)
		assert.Equal(t, 10, s)
		assert.Nil(t, err)
	})

	//you can do subtest
	t.Run("Bad Path", func(t *testing.T) {
		t.Run("Negative Should Return Error", func(t *testing.T) {
			_, err := SumPrimeWithError(-1)
			assert.Error(t, err)
		})

		t.Run("Another Negative Should Return Error", func(t *testing.T) {
			_, err := SumPrimeWithError(-3)
			assert.Error(t, err)
		})
	})

}

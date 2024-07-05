package demo

import (
	"fmt"
	"math"
)

// The capital letter signifies that this is a public function
// can be called from another package like demo.Hello()
// dubious convention but yah....

func Hello() {
	fmt.Println("Hello, World! from demo")
	v := "demo"
	fmt.Printf("Hello, World! from %s\n", v)
}

//Control flow demo

// IsPrime returns true if n is prime
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// SumPrime returns the sum of all prime numbers up to and including n
func SumPrime(n int) int {
	s := 0
	for i := 1; i <= n; i++ {
		if IsPrime(i) {
			s += i
		}
	}
	return s
}

// and no there is no exception in golang
// go lang uses concept error as value (which is considered not bad)
// the idiomatic way is to return tuple of value, error
// then the caller check the value every where

func SumPrimeWithError(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("invalid input")
	}
	return SumPrime(n), nil
}

// another way to exit is to panic
// Only Use Panic only if it totally unfixable and you "intend to quit".
// Ex: DB disconnect midway. But not user input error.

func MustSumPrime(n int) int {
	if n <= 0 {
		panic("invalid input") // this is like throwing danger exception
	}
	return SumPrime(n)
}

package solutions

import (
	"fmt"
	"math"
)

// Largest prime factor

// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

func isPrime(num int) bool {
	if num == 2 || num == 3 {
		return true
	}

	maxToCheck := int(math.Floor(math.Sqrt(float64(num))))
	for i:=2; i<=maxToCheck; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func Sol3() {

	const value = 600851475143
	maxToCheck := int(math.Floor(math.Sqrt(float64(value))))
	var maxPrimeFactor int

	for i:=2; i<maxToCheck; i++ {
		if isPrime(i) && value%i == 0 {
			maxPrimeFactor = i
		}
	}

	fmt.Printf("Answer to problem 3 = %d\n", maxPrimeFactor)
}
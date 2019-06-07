package solutions

import "fmt"

// 10001st prime

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10001st prime number?

func Sol7() {

	const expectedPrimeNumber = 10001
	var primeCount int
	for i := 2; ; i++ {
		if isPrime(i) {
			primeCount++
		}
		if primeCount == expectedPrimeNumber {
			fmt.Printf("Answer to problem 7 = %d\n", i)
			break
		}
	}
}

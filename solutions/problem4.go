package solutions

import (
	"fmt"
	"strconv"
)

// Largest palindrome product

// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

func isPalindrome(num int) bool {
	numS := strconv.Itoa(num)
	lenS := len(numS) - 1

	j := lenS
	for i := 0; i <= j; i++ {
		// fmt.Println("i = ", i)
		// fmt.Println("j = ", j)
		if numS[i] != numS[j] {
			return false
		}
		j--
	}
	return true
}

func Sol4() {
	max3digit := 999
	min3digit := 100
	var prod int
	maxProd := 101

	for i := max3digit; i >= min3digit; i-- {
		for j := i; j >= min3digit; j-- {
			prod = i * j
			// fmt.Printf(" %d ", prod)
			if isPalindrome(prod) {
				// fmt.Printf("\nAnswer to problem 4 = %d\n", prod)
				// fmt.Println("i = ", i)
				// fmt.Println("j = ", j)
				if prod > maxProd {
					maxProd = prod
				}
			}
		}
	}
	fmt.Printf("Answer to problem 4 = %d\n", maxProd)
	// fmt.Println(isPalindrome(906609))

}

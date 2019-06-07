package solutions

// Smallest multiple

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

import (
	"fmt"
)

func gcd(a,b int) int {
	for b>0 {
		a,b = b, a%b
	}
	return a
}

func Sol5(){

	// As we know num1*num2 = lcm*gcd
	// lcm = num1*num2/gcd
	// Example -->
	// for 2,3,4 => lcm = 2*3/1 = 6
	// 4,6 => lcm = 4*6/2 = 12

	lcm := 6
	cgcd := 1


	for i:=4;i<20;i++ {
		cgcd = gcd(lcm,i)
		lcm = lcm*i/cgcd
	}

	fmt.Printf("\nAnswer to problem 5 = %d\n", lcm)
}
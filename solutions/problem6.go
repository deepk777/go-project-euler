package solutions

import (
	"math"
	"fmt"
)

func Sol6(){

	const maxNumToTake = 100
	sumOfN := maxNumToTake*(maxNumToTake + 1)/2
	squareOfSumOfN := sumOfN*sumOfN

	var sumOfSquaresOfN int
	for i:=1;i<101;i++ {
		sumOfSquaresOfN += i*i
	}

	fmt.Printf("\nAnswer to problem 6 = %f\n", math.Abs(float64(sumOfSquaresOfN - squareOfSumOfN)))

}
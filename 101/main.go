package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 5, 7, 9, 11, 12, 15, 17, 21}
	primeNums := findPrimeNumber(nums)
	fmt.Println("Prime nmber is : ", primeNums)
}

func findPrimeNumber(nums []int) []int {
	var results []int
	for _, num := range nums {
		isPrime := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				isPrime = false
			}
		}
		if isPrime {
			results = append(results, num)
		}
	}
	return results
}

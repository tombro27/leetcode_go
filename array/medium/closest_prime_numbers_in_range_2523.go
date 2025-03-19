package leetcode

import "math"

// Brute force approach
// O(n*n) time complexity
func closestPrimes(left int, right int) []int {
	primes := []int{}
	first := -1
	second := -1
	if left == 1 {
		first = 2
		left = 3
	} else if !isPrime(left) {
		if left%2 == 0 {
			left++
		} else {
			left += 2
		}
	}
	minDiff := math.MaxInt
	for left <= right {
		if isPrime(left) {
			temp := first
			first = left
			second = temp
			if second != -1 && minDiff > first-second {
				minDiff = first - second
				primes = []int{second, first}
			}
		}
		left += 2
	}
	if second == -1 {
		return []int{-1, -1}
	}
	return primes
}
func isPrime(x int) bool {
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

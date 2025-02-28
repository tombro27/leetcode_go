package leetcode

// Brute force approach
// O(n^2) time complexity and O(1) space complexity
func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i; j < len(prices); j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}

// O(n) time complexity and O(n) space complexity
// traverse prices array twice
// aux[i] stores the largest number in prices[] with index >=i
func maxProfit1(prices []int) int {
	n := len(prices)
	aux := make([]int, n)
	aux[n-1] = prices[n-1]
	for i := n - 2; i >= 0; i-- {
		if prices[i] > aux[i+1] {
			aux[i] = prices[i]
		} else {
			aux[i] = aux[i+1]
		}
	}
	max := 0
	for i := 0; i < n; i++ {
		if aux[i]-prices[i] > max {
			max = aux[i] - prices[i]
		}
	}
	return max
}

// Optimal solution
// O(n) time complexity and O(1) space complexity
func maxProfit2(prices []int) int {
	max := 0
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i+1]-prices[i] > max {
			max = prices[i+1] - prices[i]
		}
		if prices[i] < prices[i+1] {
			prices[i] = prices[i+1]
		}
	}
	return max
}

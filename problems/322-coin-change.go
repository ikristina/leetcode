package problems

// CoinChange solves LeetCode 322: Coin Change
// Time Complexity: O(amount * len(coins))
// Space Complexity: O(amount)
func CoinChange(coins []int, amount int) int {
	// dp[i] represents the fewest number of coins to make up amount i
	// Initialize with amount + 1 (infinity equivalent since max coins is amount of 1s)
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	// Base case: 0 coins needed to make amount 0
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				// If we can use this coin, check if it gives fewer coins than current best
				if dp[i-coin] != amount+1 {
					dp[i] = min(dp[i], dp[i-coin]+1)
				}
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

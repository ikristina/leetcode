package problems

// ClimbStairs solves LeetCode 70: Climbing Stairs
// Time Complexity: O(n)
// Space Complexity: O(1)
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}

	// We only need to store the last two values
	// ways(n) = ways(n-1) + ways(n-2)
	prev2 := 1 // ways(n-2) initially for n=3 (which is ways(1) = 1)
	prev1 := 2 // ways(n-1) initially for n=3 (which is ways(2) = 2)

	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

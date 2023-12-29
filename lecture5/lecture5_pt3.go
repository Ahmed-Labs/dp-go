package lecture5

// Climbing stairs w/ K steps

// Only difference is recurrence relation:
// 	f(n) = f(n-1) + f(n-2) + ... + f(n-k)

// Time complexity: O(n*k)
// Space complexity: O(n)

func climbStairsKSteps(n int, k int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	
	for i:=2; i <= n; i++ {
		for j:=1; j <= k; j++ {
			if i-j < 0 {
				continue
			}
			dp[i] += dp[i-j]
		}
	}
	return dp[n]
}
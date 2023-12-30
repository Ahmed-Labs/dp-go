package lecture7

// Min cost climbing stairs
// Time complexity O(n)
// Space Xomplexity O(n)


func PaidStaircase(n int, p []int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = p[1]

	for i:=2; i <= n; i++ {
		dp[i] = p[i] + min(dp[i-1], dp[i-2])
	}
	return dp[n]
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
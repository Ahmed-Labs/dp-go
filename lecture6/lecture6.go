package lecture6


// Climbing stairs - can't step on red steps

func ClimbStairsKStepsSkipRed(n int, k int, stairs []bool) int {
	dp := make([]int, k)
	dp[0] = 1

	for i := 1; i <= n; i++ {

		for j := 1; j < k; j++ {
			if i-j < 0 {
				continue
			}
			if stairs[i-1] {
				dp[i % k] = 0
				continue
			} else {
				dp[i % k] += dp[(i-j) %k]
			}
		}
	}
	return dp[n % k]
}

func ClimbStairsKStepsSkipRed2(n int, k int, stairs []bool) int {
	dp := make([]int, k)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if stairs[i-1] {
			dp[i % k] = 0
			continue
		}
		for j := 1; j < k; j++ {
			if i-j < 0 {
				continue
			}
			dp[i % k] += dp[(i-j) %k]
		}
	}
	return dp[n % k]
}
package lecture5

// Climbing stairs w/ 3 steps



func climbStairs3Steps(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++{
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}

func climbStairs3StepsOpt(n int) int {
	first, second, third := 1, 1, 2

	for i:= 3; i <= n; i++ {
		temp := first + second + third
		first = second
		second = third
		third = temp
	}
	return third
}


package lecture5

import "github.com/Ahmed-Labs/dp/internal/util"

// Climbing stairs 
// O(n) time
// O(1) space

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i:= 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	util.PrintMemory()
	return dp[n]
}

func climbStairsOpt(n int) int {
	first, second := 1, 1

	for i:=2; i <= n; i++ {
		temp := first + second
		first = second
		second = temp
	}
	util.PrintMemory()
	return second
}
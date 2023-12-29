package lecture4


/*
Climbing Stairs

1. Objective Function:
	f(i) num of distinct ways to reach ith stair
2. Identify base cases:
	f(0) = 1
	f(1) = 1
3. Write down a recurrence relation for optimized objective function
	f(n) = f(n-1) + f(n-2)
4. Order of execution
	bottom-up
5. Where to look for answer
	f(n)

Time complexity: O(n)
Space complexity: O(n)
*/

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i:= 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairsOpt(n int) int {
	first, second := 1, 1

	for i:= 2; i <= n; i++ {
		temp := first + second
		first = second
		second = temp
	}
	return second
}
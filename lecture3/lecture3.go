package lecture3

/*
	Problem:
	Find the sum of the first N numbers.

	Objective function:
	f(i) is the sum of the first i elements.

	Recurrence relation:
	f(n) = f(n-1) + n
*/

func nSumRecur(n int) int {
	if n == 1 {
		return 1
	}
	return n + nSumRecur(n-1)
}

func nSumDP(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + i // recurrence relation
	}
	return dp[n]
}

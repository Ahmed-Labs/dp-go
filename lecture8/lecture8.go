package lecture8


// Min cost climbing stairs with path
// Reconstruct the optimal solution:
// Return the cheapest path to the top of the staircase

/*

Path reconstruction framework (for optimization problems):

path = []
store optimal subprobelms in 'from' array
for curr = best_last_state; (in this case the top of staircase idx n)
	while curr exists -> follow the pointer backwards through intermediary optimaal solutions
	curr = from[curr]
	path.push(curr)
return path.reverse
*/

func PaidStaircase(n int, p []int) []int {
	dp := make([]int, n+1)
	from := make([]int, n+1)

	dp[0] = 0
	dp[1] = p[1]

	for i:=2; i <= n; i++ {
		dp[i] = p[i] + min(dp[i-1], dp[i-2])
		if dp[i-1] < dp[i-2] {
			from[i] = i-1
		} else {
			from[i] = i-2
		}
	}

	path := []int{}
	for curr:=n; curr > 0; curr = from[curr] {
		path = append(path, curr)
	}
	
	path = append(path, 0)
	reverseArray(path, 0, len(path)-1)

	return path
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func reverseArray(arr []int, l, r int) {
	if l >= r {
		return
	}
	temp := arr[l]
	arr[l] = arr[r]
	arr[r] = temp
	reverseArray(arr, l+1, r-1)
}
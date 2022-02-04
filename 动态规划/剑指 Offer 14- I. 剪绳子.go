package main

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = 1
		for j := 1; j <= i-1; j++ {
			n := max(dp[j]*(i-j), j*(i-j))
			if n > dp[i] {
				dp[i] = n
			}
		}
	}
	return dp[n]
	//动态规划[j]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {

}

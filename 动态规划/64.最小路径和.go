package main

import "fmt"

func minPathSum(grid [][]int) int {
	//1.确定dp数组的含义
	//动态规划[m][n]

	//2.确定递推公式
	//动态规划[i][j] = max(动态规划[i-1][j]+values[i][j],动态规划[i][j-1]+values[i][j])

	//3.确定边界
	//动态规划[1][1] = grid[0][0]

	//4.遍历
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m+1)
	for j, _ := range dp {
		dp[j] = make([]int, n+1)
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j]+grid[i][j], dp[i][j-1]+grid[i][j])
		}
	}
	return dp[m-1][n-1]

}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	fmt.Println(minPathSum([][]int{[]int{1, 2, 3}, []int{4, 5, 6}}))
}

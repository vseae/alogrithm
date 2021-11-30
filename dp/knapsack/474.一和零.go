package main

import "strings"

func findMaxForm(strs []string, m int, n int) int {
	//1.确定dp数组的含义
	//2.确定dp数组递推公式
	//dp[i][j] = max(dp[i][j], dp[i - zeroNum][j - oneNum] + 1);
	//3.确定边界
	//4.遍历
	dp := make([][]int, m+1)
	for j, _ := range dp {
		dp[j] = make([]int, n+1)
	}
	for i := 0; i < len(strs); i++ {
		zeroNum, oneNum := 0, 0
		zeroNum = strings.Count(strs[i], "0")
		oneNum = strings.Count(strs[i], "1")
		for j := m; j >= zeroNum; j-- {
			for k := n; k >= oneNum; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zeroNum][k-oneNum]+1)
			}
		}
	}
	return dp[m][n]

}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func main() {

}

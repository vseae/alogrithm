package main

import "fmt"

/*
如果一个由 '0' 和 '1' 组成的字符串，是以一些 '0'（可能没有 '0'）后面跟着一些 '1'（也可能没有 '1'）的形式组成的，那么该字符串是单调递增的。

我们给出一个由字符 '0' 和 '1' 组成的字符串 S，我们可以将任何 '0' 翻转为 '1' 或者将 '1' 翻转为 '0'。

返回使 S 单调递增的最小翻转次数。



示例 1：

输入："00110"
输出：1
解释：我们翻转最后一位得到 00111.
示例 2：

输入："010110"
输出：2
解释：我们翻转得到 011111，或者是 000111。
示例 3：

输入："00011000"
输出：2
解释：我们翻转得到 00000000。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flip-string-to-monotone-increasing
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


*/
func minFlipsMonoIncr(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	n := len(s)
	if s[0] == '1' {
		dp[0][0] = 1
	} else {
		dp[0][1] = 1
	}
	for i := 1; i < n; i++ {
		if s[i] == '1' {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = min(dp[i-1][0], dp[i-1][1])
		} else {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = min(dp[i-1][0]+1, dp[i-1][1]+1)
		}
	}
	return min(dp[n-1][0], dp[n-1][1])
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func main() {
	fmt.Println(minFlipsMonoIncr("010110"))
	fmt.Println(minFlipsMonoIncr("00110"))
	fmt.Println(minFlipsMonoIncr("00011000"))
}

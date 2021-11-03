package main

import "strings"

/*
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
回文串 是正着读和反着读都一样的字符串。

示例 1：

输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
示例 2：

输入：s = "a"
输出：[["a"]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-partitioning
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func partition(s string) [][]string {
	var path []string
	var res [][]string
	backtracking(s, path, &res, 0)
	return res
}
func backtracking(s string, path []string, res *[][]string, startIndex int) {
	if startIndex == len(s) {
		//到达末尾
		pathCopy := make([]string, len(path))
		copy(pathCopy, path)
		*res = append(*res, pathCopy)
	}
	for i := startIndex; i < len(s); i++ {
		//如果是回文子串，加入路径
		if isPartition(s, startIndex, i) {
			path = append(path, s[startIndex:i+1])
		} else {
			continue
		}
		backtracking(s, path, res, i+1)
		path = path[:len(path)-1]
	}
}

//判断回文串
func isPartition(s string, start, end int) bool {
	left, right := start, end
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
func main() {

}

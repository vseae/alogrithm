package main

import (
	"strconv"
	"strings"
)

/*
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按 任何 顺序返回答案。

有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。



示例 1：

输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]
示例 2：

输入：s = "0000"
输出：["0.0.0.0"]
示例 3：

输入：s = "1111"
输出：["1.1.1.1"]
示例 4：

输入：s = "010010"
输出：["0.10.0.10","0.100.1.0"]
示例 5：

输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/restore-ip-addresses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


*/
func restoreIpAddresses(s string) []string {
	var path []string
	var res []string
	backtracking(s, path, 0, &res)
	return res
}

func backtracking(s string, path []string, startIndex int, res *[]string) {
	if startIndex == len(s) && len(path) == 4 {
		IpString := strings.Join(path, ".")
		*res = append(*res, IpString)
	}
	for i := startIndex; i < len(s); i++ {
		path := append(path, s[startIndex:i+1])
		if i-startIndex+1 <= 3 && len(path) <= 4 && isNormalIp(s, startIndex, i) {
			backtracking(s, path, i+1, res)
		} else {
			return
		}
		path = path[:len(path)-1]
	}
}

func isNormalIp(s string, start, end int) bool {
	checkInt, _ := strconv.Atoi(s[start : end+1])
	if end-start+1 > 1 && s[start] == '0' {
		return false
	}
	if checkInt > 255 {
		return false
	}
	return true
}
func main() {

}

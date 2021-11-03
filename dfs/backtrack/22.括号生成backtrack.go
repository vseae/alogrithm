package main

import "fmt"

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

有效括号组合需满足：左括号必须以正确的顺序闭合。



示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

//左括号肯定是优先加入的
//左括号必须比右括号数量多

func generateParenthesis(n int) []string {
	res := []string{}

	dfs(n, n, "", &res, n)
	return res

}
func dfs(left, right int, path string, res *[]string, n int) {
	if len(path) == 2*n {
		//fmt.Println(path)
		*res = append(*res, path)
		return
	}
	if left > 0 {
		dfs(left-1, right, path+"(", res, n)
	}
	if left < right {
		dfs(left, right-1, path+")", res, n)
	}
}

//func generateParenthesis(n int) []string {
//	res := []string{}
//
//	var dfs func(lRemain int, rRemain int, path string)
//	dfs = func(lRemain int, rRemain int, path string) {
//		if 2*n == len(path) {
//			res = append(res, path)
//			return
//		}
//		if lRemain > 0 {
//			dfs(lRemain-1, rRemain, path+"(")
//		}
//		if lRemain < rRemain {
//			dfs(lRemain, rRemain-1, path+")")
//		}
//	}
//
//	dfs(n, n, "")
//	return res
//}

func main() {
	fmt.Println(generateParenthesis(3))
}

package double_pointer

import (
	"fmt"
	"math"
)

/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。



注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。


示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"

示例 2：

输入：s = "a", t = "a"
输出："a"

示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func minWindow(s string, t string) string {
	//1，设置左右指针
	left := 0
	right := 0
	valid := 0
	window := map[byte]int{}
	need := map[byte]int{}
	start, length := 0, math.MaxInt32
	for i := range t {
		need[t[i]]++
	}
	//窗口滑动过程
	for right < len(s) {
		//2.寻找有效解
		addIndex := s[right]
		right++
		if _, ok := need[addIndex]; ok {
			window[addIndex]++
			if window[addIndex] == need[addIndex] {
				valid++
			}
		}
		//3.寻找最优解
		for valid == len(need) {
			fmt.Println(left, right)
			if right-left < length {
				start = left
				length = right - left
			}
			delIndex := s[left]
			left++
			if _, ok := need[delIndex]; ok {
				//window[delIndex]--
				//if window[delIndex] < need[delIndex] {
				//	valid--
				//}
				if window[delIndex] == need[delIndex] {
					valid--
				}
				window[delIndex]--
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	fmt.Println(start, length)
	return s[start : start+length]
}

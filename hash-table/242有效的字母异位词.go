package main

import "fmt"

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。



示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-anagram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func isAnagram(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	exist := make(map[byte]int)
	if sLen != tLen {
		return false
	}
	for i := 0; i < sLen; i++ {
		if v, ok := exist[s[i]]; v > 0 && ok {
			exist[s[i]] = v + 1
		} else {
			exist[s[i]] = 1
		}
	}
	for i := 0; i < tLen; i++ {
		if v, ok := exist[t[i]]; v >= 1 && ok {
			exist[t[i]] = v - 1
		} else {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isAnagram("abc", "csba"))
}

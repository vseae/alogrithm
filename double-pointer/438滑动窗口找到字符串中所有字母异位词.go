package main

import "fmt"

/*
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指字母相同，但排列不同的字符串。



示例 1:

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。


 示例 2:

输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func findAnagrams(s string, p string) []int {
	left, right := 0, 0
	window, need := map[byte]int{}, map[byte]int{}
	valid := 0
	var result []int
	//flag := map[int]bool{}
	for i := range p {
		need[p[i]]++
	}
	for right < len(s) {
		//扩充窗口大小
		tmpAdd := s[right]
		right++
		if _, ok := need[tmpAdd]; ok {
			window[tmpAdd]++
			if window[tmpAdd] == need[tmpAdd] {
				valid++
			}
		}
		for right-left >= len(p) {
			if valid == len(need) {
				fmt.Println(right)
				fmt.Println(left)
				result = append(result, left)
			}
			tmpDel := s[left]
			left++
			if _, ok := need[tmpDel]; ok {
				if window[tmpDel] == need[tmpDel] {
					valid--
				}
				window[tmpDel]--
			}
		}
	}
	fmt.Println(right)
	fmt.Println(left)
	return result
}

func main() {
	fmt.Println(findAnagrams("baa", "a"))
	/*
		"cbaebabacd"
		"abc"
	*/
}

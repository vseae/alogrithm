package main

//KMP

//构建前缀表-全部减1
func getNext(next []int, s string) {
	j := 0      //最大相等前后缀长度,也是前缀起始位置
	next[0] = j //起始位置为0时的最大相等前后缀长度
	// i代表当前字符串的后缀,aabaa=01012
	for i := 1; i < len(s); i++ {

		//如果不相等，那么最大相等前后缀长度向前回退
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}

		//如果前后缀相同，那么最大相等前后缀长度+1
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}
func strStr(haystack string, needle string) int {
	n := len(needle)
	if needle == "" {
		return 0
	}
	next := make([]int, len(needle))
	getNext(next, needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

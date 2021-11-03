package double_pointer

/*
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，s1 的排列之一是 s2 的 子串 。



示例 1：

输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").
示例 2：

输入：s1= "ab" s2 = "eidboaoo"
输出：false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutation-in-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


*/
func checkInclusion(s1 string, s2 string) bool {
	left, right := 0, 0
	window, need := map[byte]int{}, map[byte]int{}
	valid := 0
	for i := range s1 {
		need[s1[i]]++
	}
	for right < len(s2) {
		addIndex := s2[right]
		right++
		if _, ok := need[addIndex]; ok {
			window[addIndex]++
			if window[addIndex] == need[addIndex] {
				valid++
			}
		}
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			delIndex := s2[left]
			left++
			if _, ok := need[delIndex]; ok {

				if window[delIndex] == need[delIndex] {
					valid--
				}
				window[delIndex]--
			}
		}
	}
	return false

}

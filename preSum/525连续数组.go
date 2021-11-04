package main

import "fmt"

/*
给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。



示例 1:

输入: nums = [0,1]
输出: 2
说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
示例 2:

输入: nums = [0,1,0]
输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量 0 和 1 的最长连续子数组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/A1NYOS
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func findMaxLength(nums []int) int {
	//坐标
	cntIndex := map[int]int{0: -1}
	cnt, maxLen := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cnt--
		} else {
			cnt++
		}
		if v, ok := cntIndex[cnt]; ok {
			maxLen = max(maxLen, i-v)
		} else {
			cntIndex[cnt] = i
		}
	}
	return maxLen

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(findMaxLength([]int{0, 0, 1, 0, 0, 0, 1, 1}))
}

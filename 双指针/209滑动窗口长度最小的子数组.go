package main

import (
	"fmt"
)

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。



示例 1：

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
示例 2：

输入：target = 4, nums = [1,4,4]
输出：1
示例 3：

输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-size-subarray-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := 0
	n := len(nums)
	var subLength int
	result := n + 1
	for right < n {
		sum += nums[right]
		right++
		for sum >= target {
			subLength = right - left
			if subLength < result {
				result = subLength
			}
			sum -= nums[left]
			left++
		}
	}
	if result == n+1 {
		return 0
	} else {
		return result
	}
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

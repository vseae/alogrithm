package main

import (
	"fmt"
	"sort"
)

/*
给定一个可包含重复数字的整数集合 nums ，按任意顺序 返回它所有不重复的全排列。



示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/7p8L0Z
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(permuteUnique([]int{1, 2, 1}))
}
func permuteUnique(nums []int) [][]int {
	var track []int
	var res [][]int
	visited := make(map[int]bool)
	sort.Ints(nums)
	backtracking(track, nums, &res, visited)
	return res
}

func backtracking(track, nums []int, res *[][]int, visited map[int]bool) {
	//终止条件
	if len(track) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i >= 1 && nums[i-1] == nums[i] && visited[i-1] == false {
			continue
		}
		if visited[i] == true {
			continue
		}
		track = append(track, nums[i])
		visited[i] = true
		backtracking(track, nums, res, visited)
		track = track[:len(track)-1]
		visited[i] = false
	}
}

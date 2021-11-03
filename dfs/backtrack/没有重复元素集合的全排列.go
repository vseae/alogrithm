package main

import "fmt"

/**
给定一个不含重复数字的整数数组 nums ，返回其 所有可能的全排列 。可以 按任意顺序 返回答案。



示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]

https://leetcode-cn.com/problems/VvJkup
*/

func permute(nums []int) [][]int {
	var track []int
	//track = make([]int, len(nums)+1)
	var visited map[int]bool
	visited = make(map[int]bool)
	var res [][]int
	backtrack(visited, track, nums, &res)
	return res
}

func backtrack(visited map[int]bool, track []int, nums []int, res *[][]int) {
	//终止条件
	if len(nums) == 0 {
		return
	}
	if len(track) == len(nums) {
		tmp := make([]int, len(track))
		copy(tmp, track)         //拷贝
		*res = append(*res, tmp) //放入结果集
		return
	}
	//backtracking

	for i := 0; i < len(nums); i++ {
		//1.更新路径
		if visited[i] {
			continue
		}
		track = append(track, nums[i])
		visited[i] = true
		//2.递归
		backtrack(visited, track, nums, res)
		//3.回溯
		track = track[:len(track)-1]
		visited[i] = false

	}

}

func main() {
	num := []int{1, 2, 3}
	fmt.Println(permute(num))
}

package main

import "debug/dwarf"

func twoSum(nums []int, target int) []int {
	exist := make(map[int]int)

	for i, v := range nums {
		if idx, ok := exist[target-nums[i]]; ok {
			return []int{idx, i}
		}
		exist[nums[i]] = i
	}
	return []int{}

}

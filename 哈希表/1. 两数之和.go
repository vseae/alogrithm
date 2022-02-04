package main

func twoSum(nums []int, target int) []int {
	ha := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if idx, ok := ha[target-nums[i]]; ok {
			return []int{idx, i}
		}
		ha[nums[i]] = i
	}
	return []int{}
}

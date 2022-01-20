package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	var ans []int
	mapNum := map[int]int{}
	for _, item := range nums {
		mapNum[item]++
	}
	for key, _ := range mapNum {
		ans = append(ans, key)
	}
	sort.Slice(ans, func(a, b int) bool {
		return mapNum[ans[a]] > mapNum[ans[b]]
	})
	return ans[:k]
}

package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var track []int
	history := make(map[int]bool)
	n := len(candidates)
	sort.Ints(candidates)
	backtrack(n, 0, 0, target, history, candidates, track, &res)
	return res
}

func backtrack(n, start, sum, target int, history map[int]bool, candidates []int, track []int, res *[][]int) {
	if sum == target {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}
	if sum > target {
		return
	}
	for i := start; i < n; i++ {
		if i > 0 && candidates[i] == candidates[i-1] && history[i-1] == false {
			continue
		}
		track = append(track, candidates[i])
		sum += candidates[i]
		history[i] = true
		backtrack(n, i+1, sum, target, history, candidates, track, res)
		track = track[:len(track)-1]
		history[i] = false
		sum -= candidates[i]
	}
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}

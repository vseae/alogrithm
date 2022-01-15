package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var track []int
	n := len(candidates)
	sort.Ints(candidates)
	backtrack(n, 0, 0, target, candidates, track, &res)
	return res
}

func backtrack(n, sum, start, target int, candidates []int, track []int, res *[][]int) {
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
		track = append(track, candidates[i])
		sum += candidates[i]
		backtrack(n, sum, i, target, candidates, track, res)
		track = (track)[:len(track)-1]
		sum -= candidates[i]
	}
}

func main() {

}

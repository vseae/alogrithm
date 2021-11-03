package main

import (
	"fmt"
	"sort"
)

/**
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用一次。
注意：解集不能包含重复的组合。
*/
func combinationSum2(candidates []int, target int) [][]int {
	var trcak []int
	var res [][]int
	var history map[int]bool
	history = make(map[int]bool)
	sort.Ints(candidates)
	fmt.Println(candidates)
	backtracking(0, 0, target, candidates, trcak, &res, history)
	return res
}
func backtracking(startIndex, sum, target int, candidates, trcak []int, res *[][]int, history map[int]bool) {
	//终止条件
	if sum == target {
		tmp := make([]int, len(trcak))
		fmt.Println(trcak)
		copy(tmp, trcak) //拷贝
		fmt.Println(trcak)
		*res = append(*res, tmp) //放入结果集
		return
	}
	if sum > target {
		return
	}
	//回溯
	for i := startIndex; i < len(candidates); i++ {
		//变动逻辑在这里，
		if i > 0 && candidates[i] == candidates[i-1] && history[i-1] == false {
			continue
		}
		//更新路径集合和sum
		trcak = append(trcak, candidates[i])
		sum += candidates[i]
		history[i] = true
		//递归
		backtracking(i+1, sum, target, candidates, trcak, res, history)
		//回溯
		trcak = trcak[:len(trcak)-1]
		sum -= candidates[i]
		history[i] = false
	}
}
func main() {
	num := []int{10, 1, 2, 7, 6, 1, 5, 111}
	fmt.Println(combinationSum2(num, 8))
}

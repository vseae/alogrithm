package main

import (
	"fmt"
)

/**
【回溯问题】
给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target 的唯一组合。
candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的。
对于给定的输入，保证和为 target 的唯一组合数少于 150 个。
*/

//切入点1：题目为寻找可行解，可以使用回溯算法
//切入点2：无重复元素
//切入点3：无限制重复选取
func combinationSum(candidates []int, target int) [][]int {
	var trcak []int
	var res [][]int
	backtracking(0, 0, target, candidates, trcak, &res)
	return res
}
func backtracking(startIndex, sum, target int, candidates, trcak []int, res *[][]int) {
	//终止条件
	if sum == target {
		tmp := make([]int, len(trcak))
		copy(tmp, trcak)         //拷贝
		*res = append(*res, tmp) //放入结果集
		return
	}
	if sum > target {
		return
	}
	//回溯
	for i := startIndex; i < len(candidates); i++ {
		//更新路径集合和sum
		trcak = append(trcak, candidates[i])
		sum += candidates[i]
		//递归
		backtracking(i, sum, target, candidates, trcak, res)
		//回溯
		trcak = trcak[:len(trcak)-1]
		sum -= candidates[i]
	}

}

func main() {
	num := []int{10, 1, 2, 7, 6, 1, 5}

	fmt.Println(combinationSum(num, 8))
}

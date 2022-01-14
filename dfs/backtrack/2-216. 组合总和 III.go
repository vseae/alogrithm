package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	var track []int
	var result [][]int
	backtrack(n, k, 1, &track, &result)
	return result
}

func backtrack(n, k int, start int, track *[]int, result *[][]int) {
	if len(*track) == k {
		var sum int
		tmp := make([]int, k)
		for k, v := range *track {
			sum += v
			tmp[k] = v
		}
		if sum == n {
			fmt.Println(tmp)
			*result = append(*result, tmp)
		}
		return
	}

	for i := start; i <= 9; i++ {
		*track = append(*track, i)
		backtrack(n, k, i+1, track, result)
		*track = (*track)[:len(*track)-1]
	}
}

func main() {
	fmt.Println(combinationSum3(3, 7))
}

package main

func combine(n int, k int) [][]int {
	var res [][]int
	backTrack(n, k, 1, []int{}, &res)
	return res
}

func backTrack(n, k, start int, track []int, res *[][]int) {
	if len(track) == k {
		temp := make([]int, k)
		copy(temp, track)
		*res = append(*res, temp)
		return
	}
	for i := start; i <= n; i++ {
		track = append(track, i)
		backTrack(n, k, i+1, track, res)
		track = track[:len(track)-1]
	}
}

package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	n := len(matrix)
	m := len(matrix[0])
	i := n - 1
	j := 0
	for i >= 0 && j <= m-1 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		}
	}
	return false
}

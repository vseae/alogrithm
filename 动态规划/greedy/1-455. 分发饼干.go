package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	glen := len(g)
	slen := len(s)
	child := 0
	for i := 0; child < glen && i < slen; i++ {
		if s[i] >= g[child] {
			child++
		}
	}
	return child
}

func main() {

}

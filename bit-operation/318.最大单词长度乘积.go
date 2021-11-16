package main

import "fmt"

func maxProduct(words []string) int {
	masks := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			masks[i] |= 1 << (ch - 'a')
		}
	}
	fmt.Println(masks)
	ans := 0
	for i, x := range masks {
		for j, y := range masks[:i] {
			if x&y == 0 && len(words[i])*len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maxProduct([]string{"a", "ab", "edc"}))
}

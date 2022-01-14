package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	var result []string
	phone := [10]string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}
	if len(digits) > 4 || len(digits) == 0 {
		return nil
	}
	backtrack(0, "", digits, phone, &result)
	return result
}

func backtrack(index int, track, digits string, phone [10]string, result *[]string) {
	if len(track) == len(digits) {
		*result = append(*result, track)
		return
	}
	words := phone[digits[index]-'0']
	for i := 0; i < len(words); i++ {
		track = track + string(words[i])
		backtrack(index+1, track, digits, phone, result)
		track = track[:len(track)-1]
	}

}

func main() {
	fmt.Println(letterCombinations("2"))
}

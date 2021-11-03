package main

import "strings"

func isRegroup(s1, s2 string) bool {
	len_s1 := len([]rune(s1))
	len_s2 := len([]rune(s2))

	if len_s1 > 5000 || len_s2 > 5000 || len_s1 != len_s2 {
		return false
	}
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}

	}
	return true
}

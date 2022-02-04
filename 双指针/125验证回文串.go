package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	n := len(s)
	l, r := 0, n-1
	for l < r {
		if !unicode.IsDigit(rune(s[l])) && !unicode.IsLetter(rune(s[l])) {
			l++
			continue
		}
		if !unicode.IsDigit(rune(s[r])) && !unicode.IsLetter(rune(s[r])) {
			r--
			continue
		}
		if unicode.IsDigit(rune(s[l])) {
			if s[l] != s[r] {
				return false
			}
		} else if unicode.IsLetter(rune(s[l])) {
			if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
				return false
			}
		}
		l++
		r--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

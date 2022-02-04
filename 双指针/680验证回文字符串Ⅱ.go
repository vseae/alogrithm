package main

import "fmt"

func validPalindrome(s string) bool {
	n := len(s)
	l, r := 0, n-1
	for l < r {
		if s[l] != s[r] {
			return isPalindrome(s[l+1:r+1]) || isPalindrome(s[l:r])
		}
		l++
		r--
	}
	return true

}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("abca"))
}

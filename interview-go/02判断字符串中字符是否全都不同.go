package main

import (
	"fmt"
	"strings"
)

func main() {
	var demo_s string
	demo_s = "hello"
	result := isUniqueString((demo_s))
	fmt.Print(result)
}

func isUniqueString(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

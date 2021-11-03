package main

import (
	"strings"
	"unicode"
)

func repalceBlank(s string) (string, bool) {
	if len([]rune(s)) > 1000 {
		return s, false
	}
	for _, v := range s {
		if string(v) != " " && unicode.IsLetter(v) == false {
			return s, false
		}

	}

	//n < 0, there is no limit on the number of replacements.
	return strings.Replace(s, " ", "%20", -1), true
}

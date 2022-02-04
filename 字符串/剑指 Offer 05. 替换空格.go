package 字符串

import "strings"

func replaceSpace(s string) string {
	res := strings.ReplaceAll(s, " ", "%20")
	return res
}

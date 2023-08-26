package utils

import "strings"

/*
IncludesWithString
判断字符串是否在 字符串数组里面
*/
func IncludesWithString(stringList []string, str string) bool {
	str = strings.ToLower(str)

	for _, s := range stringList {
		s = strings.ToLower(s)
		if s == str {
			return true
		}
	}

	return false
}

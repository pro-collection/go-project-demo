package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscoreToUpperCamelCase
// 下划线单词转大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)

	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase
// 下划线单词转小写驼峰单词（首单词字母小写）
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore
// 驼峰单词转下划线单词
func CamelCaseToUnderscore(s string) string {
	var output []rune

	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}

		if unicode.IsUpper(r) {
			output = append(output, '_')
		}

		output = append(output, unicode.ToLower(r))
	}

	return string(output)
}

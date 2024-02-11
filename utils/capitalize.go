package utils

import "strings"

func Capitalize(s string) string {
	newS := strings.Clone(s)
	first := string(newS[0])
	return strings.ToUpper(first) + newS[1:]
}

package strexp

import (
	"regexp"
	"strconv"
	"strings"
)

var pattern = regexp.MustCompile(`^\d+`)

func numFromStart(s string) (int, string) {
	col := pattern.FindStringIndex(s)
	if len(col) == 0 {
		return 1, s
	}

	n, _ := strconv.Atoi(s[:col[1]])
	return n, s[col[1]:]
}

// StrExpand ...
func StrExpand(s string) string {
	if len(s) == 0 {
		return ""
	}
	sr := []rune(s)
	current := sr[0]
	secondPart := sr[1:]
	if _, err := strconv.Atoi(string(current)); err == nil {
		return ""
	}
	if current == '\\' {
		if len(sr) < 2 {
			return ""
		}
		current = sr[1]
		secondPart = sr[2:]
	}
	n, rems := numFromStart(string(secondPart))
	return strings.Repeat(string(current), n) + StrExpand(rems)
}

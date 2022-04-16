package testcase

import "github.com/huandu/xstrings"

func Compare(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func Reverse(s string) string {
	return xstrings.Reverse(s)
}

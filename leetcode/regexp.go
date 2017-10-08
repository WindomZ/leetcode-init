package leetcode

import "regexp"

func mustFindFirstStringSubmatch(expr string, s string) string {
	if strs := regexp.MustCompile(expr).FindStringSubmatch(s); len(strs) >= 1 {
		return strs[1]
	}
	return ""
}

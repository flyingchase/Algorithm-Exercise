package interviewoffer

import (
	"strings"
)

func replaceSpace(s string) string {
	if len(s) == 0 {
		return s
	}
	str := strings.Split(s, " ")
	res := strings.Join(str, "%20")
	return res
}

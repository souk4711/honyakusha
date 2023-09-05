package helpers

import (
	"strings"
)

func SplitTextIntoArray(text string) []string {
	var arr []string
	for _, ele := range strings.Split(strings.TrimSpace(text), "\n") {
		if ele = strings.TrimSpace(ele); ele != "" {
			arr = append(arr, ele)
		}
	}
	return arr
}

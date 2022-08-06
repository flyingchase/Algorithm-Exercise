package qushi

import (
	"strings"
)

func Q1(myTemplate string, keys, values []string) string {
	if len(myTemplate) == 0 {
		return myTemplate
	}
	mapping := make(map[string]string)
	for i := 0; i < len(keys); i++ {
		mapping[keys[i]] = values[i]
	}
	var count int
	for i := 0; i < len(myTemplate); i++ {
		if myTemplate[i] == '%' {
			count++
		}
	}
	s := strings.Split(myTemplate, "%")
	s2 := strings.Join(s, "")
	for index, str := range s {
		if strconv {

		}
	}
	return s2

}

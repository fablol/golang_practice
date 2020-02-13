package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	var temp_str = strings.Fields(s)
	for _, v := range temp_str {
		elem, ok := m[v]
		if ok {
			m[v] = elem + 1
		} else {
			m[v] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}

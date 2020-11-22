package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	strs := strings.Fields(s)
	var dct = make(map[string]int, len(strs))
	for _, val := range strs {
		_, ok := dct[val]
		if ok {
			dct[val]++
		} else {
			dct[val] = 1
		}
	}
	return dct
}

func main() {
	wc.Test(WordCount)
}

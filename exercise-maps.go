package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := map[string]int {}

	for _, v := range(strings.Fields(s)){
		m[v]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}


package main

import (
	"strings"
	"testing"
)

func BenchmarkString2Join(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"Welcome", "To", "China"}
		result := strings.Join(input, " ")
		if result != "Welcome To China" {
			b.Error("Unexcepted result:" + result)
		}
	}
}

func BenchmarkString2Plus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"Welcome", "To", "China"}
		var s, sep string
		for j := 0; j < len(input); j++ {
			s += sep + input[i]
			sep = " "
		}
		if s != "Welcome To China" {
			b.Error("Unexcepted result:" + s)
		}
	}
}

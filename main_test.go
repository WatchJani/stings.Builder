package main

import "testing"

func BenchmarkXxx(b *testing.B) {
	myData := []string{"my", "super", "data", "test"}

	for i := 0; i < b.N; i++ {
		BetterString(myData)
	}
}

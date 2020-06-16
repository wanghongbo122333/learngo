package forstudy

import (
	"fmt"
	"testing"
)

func initSlice() []string {
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = "I love Go"
	}
	return s
}
func initMap() map[int]string {
	m := make(map[int]string, N)
	for i := 0; i < N; i++ {
		m[i] = "Go loves you"
	}
	return m
}
func BenchmarkSprintf(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", num)
	}
}
func BenchmarkForSlice(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForSlice(s)
	}
}

func BenchmarkRangeForSlice(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForSlice(s)
	}
}

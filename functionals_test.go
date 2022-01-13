package functionals

import (
	"testing"
)

func BenchmarkReverse(b *testing.B) {
	xs := make([]int, 10000)
	for i := 0; i < b.N; i++ {
		Reverse(xs)
	}
}

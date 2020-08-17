package rand

import "testing"

func BenchmarkStringByLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringByLetters(i)
	}
}

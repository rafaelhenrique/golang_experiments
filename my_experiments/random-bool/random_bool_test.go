package justforfun

import "testing"

func BenchmarkBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bool()
	}
}

func BenchmarkAnotherBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AnotherBool()
	}
}

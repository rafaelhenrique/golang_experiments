// Ex1.3 make benchmark between examples
package ex13

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testArgs := []string{"a", "b", "c", "d"}
		s, sep := "", ""
		for _, arg := range testArgs {
			s += sep + arg
			sep = " "
		}
		fmt.Println(s)
	}
	// results
	// 1000000	      2570 ns/op
	// 2.593s
}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testArgs := []string{"a", "b", "c", "d"}
		fmt.Println(strings.Join(testArgs, " "))
	}
	// results
	// 1000000	      2513 ns/op
	// 2.537s
}

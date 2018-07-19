/*
Dup1 counts number occurences of same word in stdin

Use:
$ echo -e "banana\nbanana\npêra\npêra\nteste" | go run main.go
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		// line := input.Text()
		// counts[line] = counts[line] + 1

		counts[input.Text()]++
	}

	// note: ignore errors of input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

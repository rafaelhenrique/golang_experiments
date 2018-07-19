// Ex1.4 counts number occurences of same word from files or in stdin.
// And show name of file when duplicated line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex14: %v\n", err)
				continue
			}
			countLines(f)
			f.Close()
		}
	}
}

func countLines(f *os.File) {
	input := bufio.NewScanner(f)
	counts := make(map[string]int)

	fmt.Printf("%s\n", f.Name())
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

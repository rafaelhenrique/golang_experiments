// Ex1.3 test time between examples
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("Arguments: " + s)
	fmt.Printf("%fs elapsed\n", time.Since(start).Seconds())

	start = time.Now()
	fmt.Println("Arguments: " + strings.Join(os.Args[1:], " "))
	fmt.Printf("%fs elapsed\n", time.Since(start).Seconds())
}

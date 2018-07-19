// Echo3 show your arguments from command line.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	// fmt.Println(os.Args[1:])
}

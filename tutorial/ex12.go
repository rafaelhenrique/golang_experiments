// Ex1.2 show your arguments from command line and index line by line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args {
		fmt.Printf("%d %s\n", index, arg)
	}
}

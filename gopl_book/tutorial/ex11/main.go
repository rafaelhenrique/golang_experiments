// Ex1.1 show your arguments from command line and program name on os.Args[0].
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}

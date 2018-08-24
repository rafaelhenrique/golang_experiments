package justforfun

import (
	"fmt"
	"math/rand"
	"time"
)

func Bool() bool {
	return rand.Int31()&(1<<30) == 0
}

func AnotherBool() bool {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(100)%2 == 0)
}

func main() {
	fmt.Println(Bool())
	fmt.Println(AnotherBool())
}

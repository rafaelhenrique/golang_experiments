// Reference: https://blog.learngoprogramming.com/go-functions-overview-anonymous-closures-higher-order-deferred-concurrent-6799008dde7b

package main

import (
	"fmt"
)

type firstClassFunc func(int) int

func mul(n int) int { return n * 2 }

func calculate(nums []int, a firstClassFunc) (rnums []int) {
	rnums = append(rnums, nums...)
	for i, n := range rnums {
		rnums[i] = a(n)
	}
	return
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	newNums := calculate(nums, mul)
	fmt.Println(newNums)
}

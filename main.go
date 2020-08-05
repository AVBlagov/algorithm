package main

import (
	"algorithm/sort"
	"fmt"
)

func main() {
	arg := []int{1, 5, 2, 4, 3}
	sort.Bubble(&arg)
	fmt.Println(arg)
}

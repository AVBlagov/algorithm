package main

import (
	"algorithm/sort"
	"fmt"
)

func main() {
	arr := []int{1, 5, 2, 4, 3}
	//sort.Bubble(&arr)
	sort.Select(&arr)

	fmt.Println(arr)
}

package main

import (
	"algorithm/sort"
	"fmt"
)

func main() {
	arr := []int{1, 5, 2, 4, 3}
	//sort.Bubble(&arr)
	//sort.Select(&arr)
	///https://github.github.com/training-kit/downloads/ru/github-git-cheat-sheet/
	sort.Bubble(&arr)

	fmt.Println(arr)
}

package main

import (
	"fmt"
	"slices"
)

func SortAndMerge(left, right []int) []int{
	var result []int

	slices.Sort(left)
	slices.Sort(right)

	result = append(left, right...)
	slices.Sort(result)

	return result
}

func main() {
	fmt.Println(SortAndMerge([]int{7, 12, 6, 3, 0}, []int{5, 34, 2, 89, 3}))
}

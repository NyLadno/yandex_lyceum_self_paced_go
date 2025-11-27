package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int8{12, 78, 124, 56, 80}

	slices.SortFunc(nums, func(a, b int8) int{
		switch{
		case a < b:
			return 1
		case a > b:
			return -1
		default:
			return 0
		}
})

	fmt.Println(nums)
}

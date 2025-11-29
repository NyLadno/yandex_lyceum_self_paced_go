package main

import (
	"fmt"
	"slices"
)

func SortNames(names []string) []string{
	slices.Sort(names)
	return names
}

func main() {
	fmt.Println(SortNames([]string{"Аксинья", "Варвара", "Арина", "Есения"}))
}
package main

import (
	"fmt"

	binarysearch "github.com/rishavmehra/golangPractice/search/binary_search"
)

func main() {
	// linearSearch := linearsearch.Linear_search([]int{1, 2, 3, 4, 5, 6}, 6)
	binarysearch := binarysearch.Binary_search([]int{1, 2, 3, 4, 5, 6}, 7)
	fmt.Println(binarysearch)
}

package main

import (
	"fmt"

	"github.com/temmyscope/dsa-with-go/customsort"
)

func main() {

	sortedByQS := customsort.QuickSort([]int{7, 0, -1, 12, 2, 5, 1, 10, 15, 1, 2, 13, 9, 7, 4, 5, -3})
	sortedByBS := customsort.BubbleSort([]int{7, 0, -1, 12, 2, 5, 1, 10, 15, 1, 2, 13, 9, 7, 4, 5, -3})
	sortedByIS := customsort.InsertionSort([]int{7, 0, -1, 12, 2, 5, 1, 10, 15, 1, 2, 13, 9, 7, 4, 5, -3})

	fmt.Println(sortedByQS)
	fmt.Println(sortedByBS)
	fmt.Println(sortedByIS)
}

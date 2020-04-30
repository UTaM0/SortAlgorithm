package main

import (
	"fmt"
	sa "sortalgorithm"
)

func main() {
	arr := []int{12, 9, 15, 3, 8, 17, 6, 5}
	fmt.Println("Original array ", arr)

	sa.BubbleSort(arr)
	fmt.Println("   Bubble sort ", arr)

	sa.InsertionSort(arr)
	fmt.Println("Insertion sort ", arr)

	sa.SelectionSort(arr)
	fmt.Println("Selection sort ", arr)

	quickArr := sa.QuickSort(arr)
	fmt.Println("    Quick sort ", quickArr)

	mergeArr := sa.MergeSort(arr)
	fmt.Println("    Merge sort ", mergeArr)

}

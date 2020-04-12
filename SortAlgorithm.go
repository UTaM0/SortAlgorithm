package main

import "fmt"

// Used in Selection sort
func min_idx(a []int) int {
	min := a[0]
	idx := 0
	for i, v := range a {
		if v < min {
			min = v
			idx = i
		}
	}
	return idx
}

func BubbleSort(arr []int) []int {
	ans := make([]int, len(arr))
	copy(ans, arr)
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if ans[j+1] < ans[j] {
				ans[j], ans[j+1] = ans[j+1], ans[j]
			}
		}
	}
	return ans
}

func InsertionSort(arr []int) []int {
	ans := make([]int, len(arr))
	copy(ans, arr)
	for i := 1; i < len(arr)-1; i++ {
		for j := i + 1; j >= 1; j-- {
			if ans[j-1] > ans[j] {
				ans[j-1], ans[j] = ans[j], ans[j-1]
			}
		}
	}
	return ans
}

func SelectionSort(arr []int) []int {
	var idx int
	ans := make([]int, len(arr))
	copy(ans, arr)
	for i := 0; i <= len(arr)-1; i++ {
		idx = min_idx(ans[i:]) + len(ans[:i])
		ans[i], ans[idx] = ans[idx], ans[i]
	}
	return ans
}

func QuickSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	pivot := arr[0]

	right := make([]int, 0)
	left := make([]int, 0)
	for i := 1; i <= n-1; i++ {
		if arr[i] <= pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	r := QuickSort(right)
	l := QuickSort(left)
	return append(append(l, pivot), r...)

}

func MergeSort(arr []int) []int {
	n := len(arr)
	if n == 1 {
		return arr
	}

	var mid int = n / 2
	left := arr[:mid]
	right := arr[mid:]

	l := MergeSort(left)
	r := MergeSort(right)

	sortedArr := make([]int, 0)
	for len(l) > 0 && len(r) > 0 {

		if l[0] < r[0] {
			sortedArr = append(sortedArr, l[0])
			l = l[1:]
		} else {
			sortedArr = append(sortedArr, r[0])
			r = r[1:]
		}

		if len(l) == 0 {
			sortedArr = append(sortedArr, r...)
		} else if len(r) == 0 {
			sortedArr = append(sortedArr, l...)
		}
	}
	return sortedArr

}

func main() {
	arr := []int{12, 9, 15, 3, 8, 17, 6, 5}
	fmt.Println("Original array ", arr)
	fmt.Println("   Bubble sort ", BubbleSort(arr))
	fmt.Println("Insertion sort ", InsertionSort(arr))
	fmt.Println("Selection sort ", SelectionSort(arr))
	fmt.Println("    Quick sort ", QuickSort(arr))
	fmt.Println("    Merge sort ", MergeSort(arr))
}

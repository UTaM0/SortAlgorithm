package sortalgorithm

// Used in Selection sort
func minIdx(a []int) int {
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

// BubbleSort is one of a sort algorithm
func BubbleSort(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return
}

// InsertionSort is one of a sort algorithm
func InsertionSort(arr []int) {
	for i := 1; i < len(arr)-1; i++ {
		for j := i + 1; j >= 1; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return
}

// SelectionSort is one of a sort algorithm
func SelectionSort(arr []int) {
	var idx int
	for i := 0; i <= len(arr)-1; i++ {
		idx = minIdx(arr[i:]) + len(arr[:i])
		arr[i], arr[idx] = arr[idx], arr[i]
	}
	return
}

// QuickSort is one of a sort algorithm
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

// MergeSort is one of a sort algorithm
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

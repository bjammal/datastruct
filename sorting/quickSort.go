package sorting

//QuickSort implements quick sorting of slice of int
func QuickSort(array []int, left, right int) {
	index := partition(array, left, right)
	if left < index-1 {
		QuickSort(array, left, index-1)
	}
	if index < right {
		QuickSort(array, index, right)
	}
}

func partition(s []int, left, right int) int {
	for left <= right {
		mid := (left + right) / 2
		for s[left] < s[mid] {
			left++
		}
		for s[right] > s[mid] {
			right--
		}
		if left <= right {
			s[left], s[right] = s[right], s[left]
			left++
			right--
		}
	}
	return left
}

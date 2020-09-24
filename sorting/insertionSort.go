package sorting

//InsertionSort implemens insertion sort on integers
func InsertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		for ; j >= 0 && a[j] > key; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = key
	}
}

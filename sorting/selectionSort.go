package sorting

//SelectionSort implements selection sorting on a slice of int
func SelectionSort(s []int) {
	n := len(s)
	for i := range s {
		for j := i + 1; j < n; j++ {
			if s[j] < s[i] {
				s[j], s[i] = s[i], s[j]
			}
		}
	}
}

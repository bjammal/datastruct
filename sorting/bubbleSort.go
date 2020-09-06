package sorting

//BubbleSort implements bubble sorting on slice of int
func BubbleSort(s []int) {
	n := len(s)
	for i := range s {
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

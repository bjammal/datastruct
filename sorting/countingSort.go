package sorting

const maxInput int = 128

//CountingSort implements counting sort and return a sorted array of integers
func CountingSort(seq []int) []int {
	counts := [maxInput]int{}
	ordered := make([]int, len(seq))

	for _, n := range seq {
		counts[n]++
	}

	//modified counts
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	for _, k := range seq {
		ordered[counts[k]-1] = k
		counts[k]--
	}

	return ordered
}

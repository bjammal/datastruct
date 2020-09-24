package sorting

import (
	"sort"
)

//BucketSort implements bucket sorting
func BucketSort(a []float64) []float64 {
	// Creating 10 buckets
	buckets := make([][]float64, 10)
	for i := 0; i < len(a); i++ {
		index := int(10 * a[i])
		buckets[index] = append(buckets[index], a[i])
	}

	output := []float64{}
	for i := 0; i < 10; i++ {
		//because we are not using linked list, we need to sort after appending the elements
		sort.Float64s(buckets[i])
		output = append(output, buckets[i]...)
	}

	return output
}

package sorting

//MergeSort implements merge sorting on a slice of int
func MergeSort(s []int, left, right int) []int {
	if left == right {
		return []int{s[left]}
	}
	mid := (left + right) / 2
	leftHalf := MergeSort(s, left, mid)
	rightHalf := MergeSort(s, mid+1, right)
	return merge(leftHalf, rightHalf)

}

func merge(left, right []int) []int {
	lleft := len(left)
	lright := len(right)
	var output []int
	var i, j int = 0, 0
	for i < lleft && j < lright {
		if left[i] <= right[j] {
			output = append(output, left[i])
			i++
		} else {
			output = append(output, right[j])
			j++
		}
	}

	//copy remaining elements
	for ; i < lleft; i++ {
		output = append(output, left[i])
	}
	for ; j < lright; j++ {
		output = append(output, right[j])
	}
	return output
}

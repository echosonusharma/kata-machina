package example

func qs(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)

	qs(arr, lo, pivotIdx-1)
	qs(arr, pivotIdx+1, hi)
}

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}

	// move pivot - so we are true to the rule that
	// everything left to the pivot is <= to the pivot
	// and everything to the right is >= to the pivot

	idx++
	arr[hi], arr[idx] = arr[idx], arr[hi]

	return idx
}

func QuickSort(arr []int) {
	// lo and hi are both inclusive
	qs(arr, 0, len(arr)-1)
}

/*
Search an element in a sorted and rotated array
*/
package rotation

func findPivot(a *[]int, low, high int) int {
	// base cases
	if high == low {
		return -1
	}

	if high == low {
		return low
	}

	midPoint := (low + high) / 2

	if midPoint < high && (*a)[midPoint] > (*a)[midPoint+1] {
		return midPoint
	}

	if midPoint > low && (*a)[midPoint] < (*a)[midPoint-1] {
		return midPoint - 1
	}

	if (*a)[low] >= (*a)[midPoint] {
		return findPivot(a, low, midPoint-1)
	}

	return findPivot(a, midPoint+1, high)

}

func BinarySearch(a *[]int, low, high, searchTerm int) int {

	if high < low {
		return -1
	}

	midPoint := (low + high) / 2

	if (*a)[midPoint] == searchTerm {
		return midPoint
	}

	if (*a)[midPoint] > searchTerm {
		return BinarySearch(a, midPoint+1, high, searchTerm)
	}

	return BinarySearch(a, low, midPoint-1, searchTerm)
}

func PivotBinarySearch(a *[]int, searchTerm int) int {
	pivot := findPivot(a, 0, len(*a)-1)

	if pivot == -1 {
		// in this case array is not rotated
		BinarySearch(a, 0, len(*a)-1, searchTerm)
	}

	if (*a)[pivot] == searchTerm {
		return pivot
	}

	if (*a)[pivot] < searchTerm {
		return BinarySearch(a, 0, pivot, searchTerm)
	}

	return BinarySearch(a, pivot+1, len(*a)-1, searchTerm)
}

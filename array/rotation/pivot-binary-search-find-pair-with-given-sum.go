/*
Package rotation - ...
Search an element in a sorted and rotated array
*/
package rotation

func findPivotForSum(a *[]int, low, high int) int {
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
		return findPivotForSum(a, low, midPoint-1)
	}

	return findPivotForSum(a, midPoint+1, high)

}

func BinarySearchForSum(a *[]int, low, high, searchTerm int) int {

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

func PivotBinarySearchForSum(a *[]int, sum int) bool {
	pivot := findPivotForSum(a, 0, len(*a)-1)

	var searchTerm int

	if pivot == -1 {
		// in this case array is not rotated
		for i := 0; i < len(*a)-1; i++ {
			searchTerm = sum - (*a)[i]
			if BinarySearchForSum(a, 0, len(*a)-1, searchTerm) != -1 {
				return true
			}
		}
		return false
	}

	// in this case array is not rotated
	for i := 0; i < len(*a)-1; i++ {
		searchTerm = sum - (*a)[i]

		if (*a)[pivot] < searchTerm {
			return BinarySearchForSum(a, 0, pivot, searchTerm)
		}

		if BinarySearchForSum(a, 0, len(*a)-1, searchTerm) != -1 {
			return true
		}
	}

	return false
}

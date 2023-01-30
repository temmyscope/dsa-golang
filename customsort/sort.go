package customsort

// sort array of integers using quick sort "divide and conquer" algorithm
func QuickSort(unSortedArr []int) []int {
	arrayLength := len(unSortedArr)
	if arrayLength < 2 {
		return unSortedArr
	}

	var pivot int = unSortedArr[0]
	var leftList, rightList []int

	//since first value is pivot, start iteration from next value
	for i := 1; i < arrayLength; i++ {
		if unSortedArr[i] < pivot {
			leftList = append(leftList, unSortedArr[i])
		} else {
			rightList = append(rightList, unSortedArr[i])
		}
	}

	//recursively sort both partitions of the List, adding pivot to the left and merge them
	return append(
		append(QuickSort(leftList), pivot), QuickSort(rightList)...,
	)
}

// sort array of integers using optimized bubble sort algorithm
func BubbleSort(unSortedArr []int) []int {
	length := len(unSortedArr)
	totalIterations := 0

	for range unSortedArr {
		iterations := (length - totalIterations) - 1
		for i := 0; i < iterations; i++ {
			if unSortedArr[i] > unSortedArr[i+1] {
				unSortedArr[i], unSortedArr[i+1] = unSortedArr[i+1], unSortedArr[i]
			}
		}
		totalIterations += 1
	}
	return unSortedArr
}

// sort array of integers using insertion sort algorithm
func InsertionSort(unSortedArr []int) []int {
	length := len(unSortedArr)

	outerLimit := length - 1
	for i := 0; i < outerLimit; i++ {
		innerLimit := i + 1
		for j := innerLimit; j > 0; j-- {
			if unSortedArr[j] < unSortedArr[j-1] {
				unSortedArr[j-1], unSortedArr[j] = unSortedArr[j], unSortedArr[j-1]
			}
		}
	}
	return unSortedArr
}

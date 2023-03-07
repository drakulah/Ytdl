package array

func InsertionSort[T any](arr []T, less func(a, b T) bool) {
	arrSize := len(arr)
	for i := 0; i < arrSize; i++ {
		for j := 0; j < arrSize; j++ {
			if less(arr[i], arr[j]) {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}

func QuickSort[T any](arr []T, less func(a, b T) bool) {
	arrSize := len(arr)

	if arrSize <= 0 {
		return
	} else if arrSize <= 3 {
		InsertionSort(arr, less)
	}

	midPos := arrSize / 2
	midElem := arr[midPos]

	beforeArr := []T{}
	afterArr := []T{}

	for i := 0; i < arrSize; i++ {
		if i == midPos {
			continue
		} else if less(arr[i], midElem) {
			beforeArr = append(beforeArr, arr[i])
		} else {
			afterArr = append(afterArr, arr[i])
		}
	}

	QuickSort(afterArr, less)
	QuickSort(beforeArr, less)

	count := 0
	beforeLen := len(beforeArr)
	afterLen := len(afterArr)

	for i := 0; i < beforeLen; i++ {
		arr[count] = beforeArr[i]
		count++
	}

	arr[count] = midElem
	count++

	for i := 0; i < afterLen; i++ {
		arr[count] = afterArr[i]
		count++
	}

}

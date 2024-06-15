package sort

func SelectionSort(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[i] > nums[i+j] {
				nums[i], nums[i+j] = nums[i+j], nums[i]
			}
		}
	}

	return nums
}

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func partition(arr []int, low, high int) int {
	if len(arr) == 0 {
		return -1
	}

	pivot := arr[high]
	partition := low - 1

	for i := low; i < high; i++ {
		if arr[i] < pivot {
			partition++
			swap(arr, partition, i)
		}
	}

	swap(arr, partition+1, high)

	return partition + 1
}

func quickSort(arr []int, low, high int) {

	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}

}

func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)

	return arr
}

func partitionV2(arr []int, low, high int) int {
	pivot := arr[(low+high)/2]
	i, j := low, high

	for {

		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}

		if i >= j {
			break
		}
		swap(arr, i, j)

		i, j = i+1, j-1

	}

	return j
}

func quickSortV2(arr []int, low, high int) {
	if low < high {
		pi := partitionV2(arr, low, high)
		quickSortV2(arr, low, pi)
		quickSortV2(arr, pi+1, high)
	}
}

func QuickSortV2(arr []int) []int {
	quickSortV2(arr, 0, len(arr)-1)

	return arr
}
func IsSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	return true
}

package sort

func BubbleSort(nums []int) []int {

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}

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

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	low := start - 1

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			low++
			swap(arr, low, i)
		}
	}

	swap(arr, low+1, end)

	return low + 1
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

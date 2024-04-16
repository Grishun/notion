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

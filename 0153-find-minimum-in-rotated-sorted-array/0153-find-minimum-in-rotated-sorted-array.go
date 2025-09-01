func Partition(arr []int, low int, high int) int {
	pivot := arr[low]

	i := low + 1

	for j := low + 1; j <= high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i-1], arr[low] = arr[low], arr[i-1]

	return i - 1
}

func Quick(arr []int, low int, high int) {
	if low < high {
		p := Partition(arr, low, high)
		Quick(arr, low, p-1)
		Quick(arr, p+1, high)
	}
}


func findMin(nums []int) int {
     Quick(nums,0,len(nums)-1)

       return nums[0]
}
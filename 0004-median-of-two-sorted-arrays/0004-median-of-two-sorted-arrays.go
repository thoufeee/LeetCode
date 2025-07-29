func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    arr := append(nums1, nums2...)
	sort.Ints(arr)
	length := len(arr)

	if length%2 == 1 {
		val := length / 2
		return float64(arr[val])
	} else {
		val := length / 2
		sum := arr[val-1] + arr[val]
		return float64(sum) / 2
	}
}
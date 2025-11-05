func findDifference(nums1 []int, nums2 []int) [][]int {

    
	res := [][]int{}

	hash := make(map[int]bool)
	hash2 := make(map[int]bool)

	for _, val := range nums1 {
		hash[val] = true
	}

	for _, val := range nums2 {
		hash2[val] = true
	}
	r := []int{}
	for _, val := range nums1 {
		if !hash2[val] {
			hash2[val] = true
			r = append(r, val)
		}
	}

	e := []int{}

	for _, val := range nums2 {
		if !hash[val] {
			hash[val] = true
			e = append(e, val)
		}
	}

	res = append(res, r)
	res = append(res, e)

	return res
}
func containsNearbyDuplicate(nums []int, k int) bool {

    result := false

	hash := make(map[int]int)

	for i, val := range nums {
		if _,err := hash[val]; !err {
			hash[val] = i
		} else {
			a := hash[val]
			b := i - a
			if b <= k {
				result = true
			}
			hash[val] = i
		}

	}
	return result
    
}
func containsNearbyDuplicate(nums []int, k int) bool {

    result := false

	hash := make(map[int]*int)

	for i, val := range nums {
		if err := hash[val]; err == nil {
			hash[val] = &i
		} else {
			a := hash[val]

			// fmt.Println(i, int(*a))
			b := i - int(*a)
			// fmt.Println(b)
			if b <= k {
				result = true
			}
			hash[val] = &i
		}

	}
	return result
    
}
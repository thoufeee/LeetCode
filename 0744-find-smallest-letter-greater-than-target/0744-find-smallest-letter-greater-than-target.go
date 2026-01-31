func nextGreatestLetter(letters []byte, target byte) byte {

     min := letters[0]
	for _, val := range letters {
		if val > target {
			min = val
			break
		}
	}

	return min

}
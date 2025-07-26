func reverseDegree(s string) int {
    
    // ma := make(map[int]int)

	// for i, val := range s {

	// 	if val >= 'a' && val <= 'z' {
	// 		degree := 26 - int(val-'a')

	// 		ma[i+1] = degree
	// 	}
	// }

	// result := 0

	// for i, val := range ma {
	// 	result += i * val
	// }

    //     return result




    result := 0

	for i, val := range s {
		result += (i + 1) * int(26-(val-'a'))
	}

	     return result
}
func judgeCircle(moves string) bool {
	r, u := 0, 0
	for i := 0; i < len(moves); i++ {

		switch {
		case string(moves[i]) == "U":
			u++
		case string(moves[i]) == "R":
			r++
		case string(moves[i]) == "L":
			r--
		case string(moves[i]) == "D":
			u--
		}
	}

	return r == 0 && u == 0 
	
}
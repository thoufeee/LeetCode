func judgeCircle(moves string) bool {
	sv := strings.ToLower(moves)
	r, u := 0, 0
	for i := 0; i < len(sv); i++ {

		switch {
		case string(sv[i]) == "u":
			u++
		case string(sv[i]) == "r":
			r++
		case string(sv[i]) == "l":
			r--
		case string(sv[i]) == "d":
			u--
		}
	}

	if r == 0 && u == 0 {
		return true
	} else {
		return false
	}

}
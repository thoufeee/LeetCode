func reversePrefix(s string, k int) string {

    s2 := s[:k]
	res := ""
	for i := len(s2) - 1; i >= 0; i-- {
		res += string(s2[i])
	}

	res += s[k:]

    return res
    
}
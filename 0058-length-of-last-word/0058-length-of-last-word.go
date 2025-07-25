func lengthOfLastWord(s string) int {
    a := strings.TrimSpace(s)
	val := strings.Split(a, " ")
	c := len(val) - 1
	return len(val[c])
}
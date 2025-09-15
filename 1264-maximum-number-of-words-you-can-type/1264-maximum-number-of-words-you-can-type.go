func canBeTypedWords(text string, brokenLetters string) int {
           
	arr := strings.Split(text, " ")
	result := 0

	for _, val := range arr {
		if !strings.ContainsAny(val, brokenLetters) {
			result++
		}
	} 
	  return result
}
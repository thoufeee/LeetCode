func detectCapitalUse(word string) bool {
    con1 := strings.ToUpper(word)
	con2 := strings.ToLower(word)
	con3 := strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))

    if word == con1 {
          return true
    }else if word == con2 {
          return true
    }else if word == con3 {
          return true
    }
       return false
}
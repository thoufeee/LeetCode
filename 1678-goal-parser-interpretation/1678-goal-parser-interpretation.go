func interpret(command string) string {
    res := ""

    for i := 0; i < len(command); i++ {
        if string(command[i]) == "G" {
            res += "G"
        } else if string(command[i+1]) == ")" {
            res += "o"
            i += 1
        } else if string(command[i+1]) == "a" && string(command[i+2]) == "l" && string(command[i+3]) == ")" {
            res += "al"
            i += 3
        }
    }

    return res
    
}
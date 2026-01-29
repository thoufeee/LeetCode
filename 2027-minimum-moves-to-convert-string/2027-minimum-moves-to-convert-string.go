func minimumMoves(s string) int {

    res := 0
    i := 0

    for i < len(s) {
           if string(s[i]) == "X" {
                res++
                i+= 3
           }else {
             i++
           }
    }
       return res
    
}
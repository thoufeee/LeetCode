
func Climb (val int) int {
      if val <= 2 {
           return val
      }

      hash := make([]int, val+1) 
      hash[1] = 1
      hash[2] = 2

      for i:=3;i<=val;i++ {
           hash[i] = hash[i-1] + hash[i-2]
      } 
        return hash[val]
}


func climbStairs(n int) int {
       
       result := Climb(n)

        return result
}
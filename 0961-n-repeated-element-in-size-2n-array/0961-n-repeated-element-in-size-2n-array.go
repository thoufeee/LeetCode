func repeatedNTimes(nums []int) int { 
      hash := make(map[int]int)

      for _,val := range nums {
          if hash[val] == 1 {
             return val
          }
            hash[val]++
      }

         return -1

         
}
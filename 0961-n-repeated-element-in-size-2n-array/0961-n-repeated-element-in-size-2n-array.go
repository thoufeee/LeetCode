func repeatedNTimes(nums []int) int {

      hash := make(map[int]int)

      for _,val := range nums {
            hash[val]++
      }

      for val,count := range hash {
           if count > 1 {
               return val
           }
      }
          return -1
}
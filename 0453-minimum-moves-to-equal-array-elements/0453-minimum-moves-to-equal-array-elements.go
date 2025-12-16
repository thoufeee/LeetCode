func minMoves(nums []int) int {
     min := nums[0]
    res := 0

    for _,val := range nums {
         if val < min {
              min = val
         }
    }

    for _,val := range nums {
                 res += val - min
    }

      return res
    
}
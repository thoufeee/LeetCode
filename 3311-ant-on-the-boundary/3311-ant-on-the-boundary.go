func returnToBoundaryCount(nums []int) int {
      result := 0
      sum := 0;

   for _,val := range nums {
        sum += val

        if sum == 0 {
             result++
        }
   }

     return result
}
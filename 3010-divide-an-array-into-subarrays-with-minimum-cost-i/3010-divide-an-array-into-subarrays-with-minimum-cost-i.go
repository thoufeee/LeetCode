func minimumCost(nums []int) int {

        min := 1000000
        min2 := 1000000

          for i:=1;i<len(nums);i++ {
                if nums[i] < min {
                       min2 = min
                       min = nums[i]
                }else if nums[i] < min2 {
                        min2 = nums[i]
                }
          }
               return nums[0] + min + min2
}
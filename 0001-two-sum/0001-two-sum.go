func twoSum(nums []int, target int) []int {
       hash := make(map[int] int)

       for i,val := range nums {
            diff := target - val
           if ind,value := hash[diff]; value {
                     return [] int {ind, i}
            }
               hash[val] = i
       }
          return [] int {}
}
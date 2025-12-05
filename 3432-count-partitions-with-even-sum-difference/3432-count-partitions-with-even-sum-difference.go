func countPartitions(nums []int) int {

    res := 0

    for i:=0;i<len(nums)-1;i++ {
           sum := 0
           sum2 := 0

           for j:=0;j<=i;j++ {
               sum += nums[j]
           }

           for k:=i+1;k<len(nums);k++ {
               sum2 += nums[k]
           }

             even := sum - sum2

             if even % 2 == 0 {
                  res++
             }
    }

          return res
}
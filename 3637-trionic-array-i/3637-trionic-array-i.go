func isTrionic(nums []int) bool {
     
      n := len(nums)

      if n < 4 {
           return false
      }
    
      i := 0

      for i+1 < n && nums[i] < nums[i+1] {
           i++
      }

      if i == 0 {
           return false
      }
       s := i

      for i+1 < n && nums[i] > nums[i+1] {
           i++
      }

      if s == i {
           return false
      }

       s2 := i

      for i+1 < n && nums[i] < nums[i+1] {
        i++
      }

        if s2 == i {
              return false
        }

        return i == n-1
}
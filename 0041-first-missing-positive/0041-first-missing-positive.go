func firstMissingPositive(nums []int) int {
    hash := make(map[int]bool)

    for _,val := range nums {
           hash[val] = true
    }

    for i:=1;i<=len(nums)+1;i++ {
            if !hash[i] {
                  return i
            }
    }
            return 2
}
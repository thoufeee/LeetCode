func search(nums []int, target int) int {
    result := -1;
    for i,val := range nums {
          if val == target {
                result = i;
          }
    }
        return result;
}
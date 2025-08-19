func findDuplicates(nums []int) []int {
        hash := make(map[int]int)
        result := []int {}

        for _,val := range nums {
               hash[val]++
        }
        for i,val := range hash {
               if val > 1 {
                   result = append(result,i)
               }
        }

           return result
    
}
func findDuplicates(nums []int) []int {
        hash := make(map[int]bool)
        check := make(map[int]bool)
        result := []int {}
 
        for _,val := range nums {
              if hash[val] {
                  if !check[val] {
                       result = append(result,val)
                       check[val] = true
                  }
                   
              }else {
                   hash[val] = true
              }
        }
     
           return result
    
}
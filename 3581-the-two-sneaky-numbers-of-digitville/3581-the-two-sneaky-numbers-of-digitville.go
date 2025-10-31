func getSneakyNumbers(nums []int) []int {

    result := [] int {}

    hash := make(map[int]int)

    for _,val := range nums {
           hash[val]++
    }

    for val,count := range hash {
          if count > 1 {
               result = append(result,val)
          }
    }

       return result
    
}
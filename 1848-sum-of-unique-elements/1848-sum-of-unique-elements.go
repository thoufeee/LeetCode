func sumOfUnique(nums []int) int {

  hash := make(map[int]int)

  for _,val := range nums {
       hash[val]++
  }

  result := 0


  for va,co := range hash {
        if co == 1 {
               result += va
        }
  }

     return result 
}
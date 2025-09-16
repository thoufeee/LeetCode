func twoSum(numbers []int, target int) []int {

      hash := make(map[int]int)

      for i,val := range numbers {
            diff := target - val

        if j,res := hash[diff]; res {
              return [] int {j+1, i+1}
        }

          hash[val] = i
      }

         return [] int {}
}
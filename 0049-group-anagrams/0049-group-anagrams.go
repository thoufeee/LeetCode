func groupAnagrams(strs []string) [][]string {
            hash := make(map[string][]string)

            for _,val := range strs {
                  s := Sort(val)
                hash[s] = append(hash[s], val)
            }

             result := [][] string{}
             for _,val := range hash{
                    result = append(result,val)
             }

            return result
}

 func Sort(s string)string{
         r := []rune(s)
         sort.Slice(r, func(i,j int) bool{
             return  r[i] < r[j]
         })
             return string(r)
 }
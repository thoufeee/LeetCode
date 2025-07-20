func longestCommonPrefix(strs []string) string {

      if len(strs) == 0 {
           return ""
      }

       base := strs[0];

       for i:=0;i<len(base);i++ {
            for _,val := range strs[1:] {
                   if i == len(val) || base[i] != val[i] {
                         return base[:i]
                   }
            }
       }
           return base
    
}
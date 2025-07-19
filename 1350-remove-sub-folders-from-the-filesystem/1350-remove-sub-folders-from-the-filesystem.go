func removeSubfolders(folder []string) []string {
    
         sort.Strings(folder)
        result := [] string {}

        for _,val := range folder {
            if len(result) == 0 || !strings.HasPrefix(val, result[len(result)-1] + "/") {
                       result = append(result, val)
            }
        }

           return result;
}
func numOfSubarrays(arr []int, k int, threshold int) int {

    sum := 0
    res := 0

    for i:=0;i<k;i++ {
          sum += arr[i]
    }

    var avg float64
    sum2 := sum

    avg = float64(sum2) / float64(k)

    if avg >= float64(threshold) {
              res++
           }
           
    for i:=k;i<len(arr);i++ {
           
     sum2 -= arr[i-k]
     sum2 += arr[i]

     avg = float64(sum2) / float64(k)

     if avg >= float64(threshold) {
              res++
           }

    }

    return res
    
}
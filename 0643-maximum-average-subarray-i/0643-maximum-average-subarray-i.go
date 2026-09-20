func findMaxAverage(nums []int, k int) float64 {

    var n float64 = float64(k)

    var maxsum float64 = 0

    for i:=0;i<k;i++ {
           maxsum += float64(nums[i])
    }

     slidmax := maxsum

     for i:=k;i<len(nums);i++ {
          slidmax -= float64(nums[i-k])

           slidmax += float64(nums[i])

           if slidmax > maxsum  {
                maxsum = slidmax
           }
     }

     return maxsum / n
}
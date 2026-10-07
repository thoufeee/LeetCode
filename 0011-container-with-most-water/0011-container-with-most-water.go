func maxArea(height []int) int {

    max := 0

    left := 0
    right := len(height)-1

    for left < right {
          sum := 0

          smallheight := 0

          if height[left] < height[right] {
                 smallheight = height[left]
          }else {
             smallheight =  height[right]
          }


      width := right - left

      sum = width * smallheight

      if height[left] < height[right] {
           left++
      }else {
        right--
      }

      if sum > max {
         max = sum
      }
    }

    return max
    
}

// func Sub(arr []int, n int, res *[][] int){
//        if n == len(arr) {
//            return 
//        }

//        for i:=n;i<len(arr);i++ {
//              sub := make([]int, i-n+1)
//               copy(sub, arr[n:i+1])
//               *res = append(*res, sub)
//        }

//        Sub(arr,n+1,res)
// }


func maxSubArray(nums []int) int {

    // if len(nums) == 1 {
    //       return nums[0]
    // }
    
    //    res := [][]int{}

    //    Sub(nums,0,&res)

    //    result := -100000

    //    for _,val := range res {
    //         sum := 0;
    //         for _,va := range val {
    //                sum += va;
    //         }

    //         if sum > result {
    //               result = sum
    //         }
    //    }  
    //      return result  

        max := nums[0]
        result := nums[0]

        for i:=1;i<len(nums);i++ {
               if max + nums[i] > nums[i] {
                  max = max + nums[i]
               }else {
                  max = nums[i]
               }

               if max > result {
                  result = max
               }
        }
                  return result
}
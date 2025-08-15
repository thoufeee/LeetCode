
// func Quick(arr []*big.Int, low,high int) {
//       if low < high {
//           p := Partition(arr, low,high)
//           Quick(arr,low,p-1)
//           Quick(arr,p+1,high)

//       }
// }

// func Partition(arr []*big.Int, low,high int)int{
//        pivot := arr[low]
//        i := low+1

//        for j:=low+1;j<=high;j++ {
//             if arr[j].Cmp(pivot) > 0 {
//                    arr[i],arr[j] = arr[j],arr[i]
//                    i++
//             }
//        }
//          arr[low],arr[i-1] = arr[i-1],arr[low]
//            return i-1
// }


func kthLargestNumber(nums []string, k int) string {
    
        arr := []*big.Int{}

	for _, val := range nums{
		num := new(big.Int)
        num.SetString(val,10)
		arr = append(arr, num)
	}

	sort.Slice(arr, func(i,j int)bool {
           return arr[i].Cmp(arr[j]) > 0
    })
  

    return arr[k-1].String()
}
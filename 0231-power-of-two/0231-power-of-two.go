func isPowerOfTwo(n int) bool {
//    for i := 0; i <= n; i++ {
// 		exp := big.NewInt(int64(i))
// 		base := big.NewInt(2)
// 		re := new(big.Int).Exp(base, exp, nil)

// 		val := big.NewInt(int64(n))
// 		if re.Cmp(val) == 0 {
// 			return true
// 		}
// 	}

// 	return false

     return n > 0 && (n & (n-1)) == 0
}
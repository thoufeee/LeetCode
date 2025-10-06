func multiply(num1 string, num2 string) string {
    
	num3 := new(big.Int)

	num4 := new(big.Int)

	num3.SetString(num1, 10)
	num4.SetString(num2, 10)

	e := new(big.Int).Mul(num3, num4)

	result := e.String()

      return result

}
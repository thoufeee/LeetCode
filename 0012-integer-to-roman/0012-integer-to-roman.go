func intToRoman(num int) string {
    val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	rom := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

    result := ""

    for i:=0;i<len(val);i++ {
          for num >= val[i] {
                result += rom[i]
                num -= val[i]
          }
    }

       return result
}
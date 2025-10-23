func fib(n int) int {

   
    if n < 1 {
      return 0
    }else if n == 1 {
        return 1
    }
     sum := 0
     
      a,b := 0,1
    for i := 1; i <= n; i++ {
		a, b = b, a+b

		if i == n-2 {
			sum += a
		} else if i == n-1 {
			sum += a
		}
	}
    
    return sum
}
func numWaterBottles(numBottles int, numExchange int) int {
    result := 0
	bottle := 0

    for numBottles > 0 {
          result += numBottles
          bottle += numBottles 

          numBottles = bottle / numExchange
          bottle = bottle % numExchange
    }

	return result
}
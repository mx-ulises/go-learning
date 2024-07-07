func numWaterBottles(numBottles int, numExchange int) int {
	drunkBottles := numBottles
	emptyBottles := numBottles
	for numExchange <= emptyBottles {
		drinkableBottles := emptyBottles / numExchange
		emptyBottles %= numExchange
		drunkBottles += drinkableBottles
		emptyBottles += drinkableBottles
	}
	return drunkBottles
}

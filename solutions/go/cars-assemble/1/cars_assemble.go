package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// panic("CalculateWorkingCarsPerHour not implemented")
    return float64(float64(productionRate) * successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// panic("CalculateWorkingCarsPerMinute not implemented")
    productionPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
    return int(productionPerHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// panic("CalculateCost not implemented")
    costOfOneCar := 10000
	costOfTenCars := 95000

	carsOf10 := int(carsCount / 10)
	remainingCars := int(carsCount % 10)

	return uint(carsOf10*costOfTenCars + remainingCars*costOfOneCar)
}

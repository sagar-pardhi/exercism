package lasagna

func PreparationTime(layers []string, minutes int) int {
	if minutes == 0 {
		minutes = 2
	}

	return len(layers) * minutes
}

func Quantities(layers []string) (int, float64) {
	noodleQuantity := 0
	sauceQuantity := 0.0

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodleQuantity += 50
		} else if layers[i] == "sauce" {
			sauceQuantity += 0.2
		}
	}

	return noodleQuantity, sauceQuantity

}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, numberOfPortionsToCook int) []float64 {
	newQuantities := []float64{}

	for i := 0; i < len(quantities); i++ {
		newQuantities = append(newQuantities, quantities[i]*float64(numberOfPortionsToCook)/2)
	}

	return newQuantities
}

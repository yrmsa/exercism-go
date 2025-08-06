package lasagna

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(layers []string, averageMinutePerLayer int) int {
	if (averageMinutePerLayer == 0) {
		averageMinutePerLayer = 2
	}
	
	return len(layers) * averageMinutePerLayer
}

// Quantities calculates how many noodles and sauces needed to prepare for the lasagna based on the layers
func Quantities(layers []string) (noodles int, sauces float64) {
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		}

		if layers[i] == "sauce" {
			sauces += 0.2
		}
	}

	return
}

// AddSecretIngredient replace the last ingredient in the second argument with the last ingredient in the first argument
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe scale the default amount for two portions into specific portions
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := append([]float64 {}, quantities...)
	for i := 0; i < len(scaledQuantities); i++ {
		scaledQuantities[i] = scaledQuantities[i] / 2 * float64(portions)
	}

	return scaledQuantities
}

package helpers

// CleanedArr removes empty strings from the inputArr and updates the inputArr and index i accordingly
func CleanedArr(inputArr *[]string, i *int) {
	// Create a new slice to store non-empty strings
	CleanedArr := make([]string, 0, len(*inputArr))
	// Iterate through the inputArr and filter out empty strings
	for _, arg := range *inputArr {
		if arg != "" {
			CleanedArr = append(CleanedArr, arg)
		}
	}

	// Adjust the index i if there are empty strings removed
	if len(CleanedArr) < len(*inputArr) {
		*i = *i - (len(*inputArr) - len(CleanedArr))
	}

	// Update the inputArr with the cleaned version
	*inputArr = CleanedArr
}

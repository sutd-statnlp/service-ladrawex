package sliceutil

// ContainString returns true when the slice contains given strings, otherwise false.
func ContainString(slice []string, inputStrings ...string) bool {
	inputLen := len(inputStrings)
	if inputLen == 0 {
		return false
	}
	var matchingCount uint8
	for _, inputString := range inputStrings {
		for _, item := range slice {
			if item == inputString {
				matchingCount++
			}
		}
	}
	if matchingCount == uint8(inputLen) {
		return true
	}
	return false
}

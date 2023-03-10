package slices

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func SlicesAreEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func IntToSlicePadded(rawNumber, digits int) []int {

	slice := []int{}
	currentDigit := 0

	for currentDigit < digits {
		slice = append([]int{rawNumber % 10}, slice...)
		rawNumber = rawNumber / 10
		currentDigit++
	}

	return slice
}

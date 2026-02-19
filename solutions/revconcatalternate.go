package piscine

func RevConcatAlternate(slice1, slice2 []int) []int {
	lenS1 := len(slice1)
	lenS2 := len(slice2)
	result := make([]int, 0, lenS1+lenS2)
	i := lenS1 - 1
	j := lenS2 - 1

	// Add extra elements from the larger slice first
	if lenS1 > lenS2 {
		for k := 0; k < lenS1-lenS2; k++ {
			result = append(result, slice1[i])
			i--
		}
	} else if lenS2 > lenS1 {
		for k := 0; k < lenS2-lenS1; k++ {
			result = append(result, slice2[j])
			j--
		}
	}

	// Now alternate between the remaining elements
	// At this point, both have equal remaining elements
	for i >= 0 && j >= 0 {
		result = append(result, slice1[i])
		i--
		result = append(result, slice2[j])
		j--
	}

	return result
}

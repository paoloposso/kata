package equalizearrays

func equalizeArray(arr []int32) int32 {
	m := map[int32]int{}

	for _, v := range arr {
		value := m[v]
		m[v] = value + 1
	}

	biggest := 0

	// Iterate over the map and append values to the slice
	for _, quantity := range m {
		if quantity > biggest {
			biggest = quantity
		}
	}

	return int32(len(arr) - biggest)
}

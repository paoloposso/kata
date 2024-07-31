package arraysaresimilar

func solution(a, b []int) bool {
	differencesIndex := []int{}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			differencesIndex = append(differencesIndex, i)
		}

		if len(differencesIndex) > 2 {
			return false
		}
	}

	if len(differencesIndex) == 0 {
		return true
	}

	if len(differencesIndex) != 2 {
		return false
	}

	return a[differencesIndex[0]] == b[differencesIndex[1]] && a[differencesIndex[1]] == b[differencesIndex[0]]
}

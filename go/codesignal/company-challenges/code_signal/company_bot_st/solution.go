package companybotst

func solution(trainingData [][]int) float64 {
	total := 0
	count := 0

	for i := range trainingData {
		if trainingData[i][1] == 1 {
			total += trainingData[i][0]
			count++
		}
	}

	if count == 0 {
		return 0
	}

	return float64(float64(total) / float64(count))
}

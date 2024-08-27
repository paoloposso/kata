package blur

func solution(image [][]int) [][]int {

	result := [][]int{}

	for x := 0; x+3 <= len(image); x++ {
		result = append(result, []int{})
		for y := 0; y+3 <= len(image[x]); y++ {
			r := getAvg(image, x, y)
			result[x] = append(result[x], r)
		}
	}

	return result
}

func getAvg(square [][]int, startX, startY int) int {
	total := 0
	for i := startX; i < 3+startX; i++ {
		for j := startY; j < 3+startY; j++ {
			total += square[i][j]
		}
	}
	return total / 9
}

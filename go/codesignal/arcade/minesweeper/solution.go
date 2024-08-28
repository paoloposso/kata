package minesweeper

func solution(matrix [][]bool) [][]int {
	result := [][]int{}
	for i := 0; i < len(matrix); i++ {
		result = append(result, []int{})
		for j := 0; j < len(matrix[i]); j++ {
			result[i] = append(result[i], countNeighbors(matrix, i, j))
		}
	}

	return result
}

func countNeighbors(matrix [][]bool, i, j int) int {
	total := 0

	if i < len(matrix)-1 {
		if matrix[i+1][j] {
			total++
		}
		if j < len(matrix[i])-1 && matrix[i+1][j+1] {
			total++
		}
		if j > 0 && matrix[i+1][j-1] {
			total++
		}
	}
	if i > 0 {
		if matrix[i-1][j] {
			total++
		}
		if j > 0 && matrix[i-1][j-1] {
			total++
		}
		if j < len(matrix[i])-1 && matrix[i-1][j+1] {
			total++
		}
	}
	if j > 0 && matrix[i][j-1] {
		total++
	}
	if j < len(matrix[i])-1 && matrix[i][j+1] {
		total++
	}

	return total
}

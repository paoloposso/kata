package chessboard

const black string = "black"
const white string = "white"

func solution(cell1 string, cell2 string) bool {
	return getColor(cell1) == getColor(cell2)
}

func getColor(cell string) string {
	if (cell[0])%2 == 0 && cell[1]%2 == 0 {
		return black
	} else if (cell[0])%2 == 1 && cell[1]%2 == 1 {
		return black
	}
	return white
}

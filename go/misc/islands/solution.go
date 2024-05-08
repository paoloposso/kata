package islands

type Coordinate struct {
	i, j int
}

func countIslands(m [][]bool) int {
	totatIslands := 0
	visited := []Coordinate{}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if isVisited(i, j, visited) {
				continue
			}

			if isLand(i, j, m) {
				totatIslands++
				visited = bfs(i, j, m, visited)
			} else {
				visited = append(visited, Coordinate{i, j})
			}
		}
	}

	return totatIslands
}

func isVisited(i, j int, coordinates []Coordinate) bool {
	for _, v := range coordinates {
		if v.i == i && v.j == j {
			return true
		}
	}
	return false
}

func bfs(i, j int, n [][]bool, visited []Coordinate) []Coordinate {
	if !isLand(i, j, n) || isVisited(i, j, visited) {
		return visited
	}

	visited = append(visited, Coordinate{i, j})

	visited = bfs(i-1, j, n, visited)
	visited = bfs(i+1, j, n, visited)
	visited = bfs(i, j-1, n, visited)
	visited = bfs(i, j+1, n, visited)

	return visited
}

func isLand(i, j int, n [][]bool) bool {
	if i < 0 || i >= len(n) {
		return false
	}
	if j < 0 || j >= len(n[i]) {
		return false
	}
	return n[i][j]
}

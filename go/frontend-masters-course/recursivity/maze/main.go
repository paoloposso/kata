package main

//todo - finish

type Point struct {
	x int
	y int
}

func Walk(maze [][]string, wall string, current Point, end Point, seen [][]bool, pathStack []Point) bool {
	//off the map
	if current.x < 0 || current.x > len(maze) {
		return false
	}
	if current.y < 0 || current.y > len(maze[0]) {
		return false
	}
	if maze[current.x][current.y] == wall {
		return false
	}
	//end of the maze
	if current.x == end.x && current.y == end.y {
		return true
	}
	if seen[current.x][current.y] {
		return false
	}

	pathStack = append(pathStack, current)

	return true
}

func main() {

}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//todo - finish

type Point struct {
	x int
	y int
}

func generateRandomMaze(x int, y int, sparseness int) (maze [][]string, start Point, end Point) {
	for i := 0; i < x; i++ {
		line := []string{}

		for j := 0; j < y; j++ {
			line = append(line, " ")
		}
		maze = append(maze, line)
	}

	end = getRandomPoint(x, y)
	maze[end.x][end.y] = "$"

	start = getRandomPoint(x, y)
	maze[start.x][start.y] = "@"

	return maze, start, end
}

func printMaze(maze [][]string) {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			fmt.Printf("[%s]", maze[i][j])
		}
		fmt.Println()
	}
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
	maze, _, _ := generateRandomMaze(4, 4, 0)
	printMaze(maze)
}

func getRandomPoint(maxX, maxY int) Point {

	result := Point{}

	rand.Seed(time.Now().UnixNano())
	result.x = rand.Intn(maxX)

	rand.Seed(time.Now().UnixNano())
	result.y = rand.Intn(maxY)

	return result
}

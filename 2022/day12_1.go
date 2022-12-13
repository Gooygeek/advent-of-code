package main

import (
	"bufio"
	"fmt"
	"os"
)

type Square struct {
	x         int
	y         int
	elevation int
	visited   bool
	neighbors []*Square
	from      *Square
}

func main() {
	// fmt.Println("Hello, world!")
	input := []string{}
	f, err := os.Open("./day12_1.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// generate a graph of paths.
	// For each square: look at the surrounding squares and note any with an elevation at most 1 higher or lower.
	// From this we can use a breadth-first search to find the shortest path between the start and end squares.

	// Generate grid of elevations
	var start *Square
	var end *Square
	grid := [][]Square{}
	for i, row := range input {
		grid = append(grid, []Square{})
		for j, square := range row {
			curSquare := Square{x: j, y: i, elevation: int(square)}
			if curSquare.elevation == 83 {
				start = &curSquare
				curSquare.elevation = 97
			}
			if curSquare.elevation == 69 {
				end = &curSquare
				curSquare.elevation = 122
			}
			grid[i] = append(grid[i], curSquare)
		}
	}
	// NOTE: the starting elevation (S) is '83' and the ending elevation (E) is '69'

	// for _, v := range grid {
	// 	for _, s := range v {
	// 		fmt.Printf("%v, ", s.elevation)
	// 	}
	// 	fmt.Printf("\n")
	// }

	// Populate neighbors
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			curSquare := grid[i][j]
			if i != 0 {
				if grid[i-1][j].elevation-curSquare.elevation <= 1 {
					curSquare.neighbors = append(curSquare.neighbors, &grid[i-1][j])
				}
			}
			if i != len(grid)-1 {
				if grid[i+1][j].elevation-curSquare.elevation <= 1 {
					curSquare.neighbors = append(curSquare.neighbors, &grid[i+1][j])
				}
			}
			if j != 0 {
				if grid[i][j-1].elevation-curSquare.elevation <= 1 {
					curSquare.neighbors = append(curSquare.neighbors, &grid[i][j-1])
				}
			}
			if j != len(grid[i])-1 {
				if grid[i][j+1].elevation-curSquare.elevation <= 1 {
					curSquare.neighbors = append(curSquare.neighbors, &grid[i][j+1])
				}
			}
			grid[i][j] = curSquare
		}
	}

	// for _, v := range grid {
	// 	for _, s := range v {
	// 		fmt.Printf("%+v, ", s)
	// 	}
	// 	fmt.Printf("\n")
	// }

	// Perform Breadth first search
	// fmt.Printf("%+v\n", start)
	// fmt.Printf("%+v\n", end)
	breadthCount := BreadthFirstSearch(&grid[start.y][start.x], &grid[end.y][end.x])
	fmt.Printf("Number of steps from S to E: %v\n", breadthCount-1)
}

// BreadthFirstSearch performs a breadth-first search on a graph
// starting at the given node it finds all the next nodes to traverse in 'waves', keeping track of the number of waves
// finally, it returns the number of waves searched till the end node was found
func BreadthFirstSearch(start *Square, end *Square) int {
	// Start square comes from itself
	start.from = start

	// Create a slice to hold the visited nodes
	visited := make([]Square, 0)

	// Create a queue to hold the nodes we need to visit
	queue := make([]*Square, 0)

	// Add the start node to the queue
	queue = append(queue, start)

	breadthCount := 0

	// Keep looping until the queue is empty
	for len(queue) > 0 {
		breadthCount++
		// fmt.Printf("Starting Wave: %v\n", breadthCount)
		queue, visited = findNextBreadth(queue, visited, end)
	}

	printVisited(visited)

	return breadthCount
}

func findNextBreadth(queue []*Square, visited []Square, end *Square) ([]*Square, []Square) {
	newQueue := make([]*Square, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if !node.visited {
			node.visited = true
			// fmt.Printf("Visiting: %+v\n", node)

			// Add the node's value to the visited slice
			visited = append(visited, *node)

			if node.x == end.x && node.y == end.y {
				// fmt.Println("FOUND IN WAVE!!")
				return []*Square{}, visited
			} else {
				// Add the node's edges to the queue
				for _, neighbor := range node.neighbors {
					if neighbor.visited == false {
						neighbor.from = node
						newQueue = append(newQueue, neighbor)
					}
				}
			}
		}
	}
	// fmt.Println("NOT FOUND IN WAVE")
	return newQueue, visited
}

func printVisited(visited []Square) {
	maxX := 163
	maxY := 41

	visitedGrid := [][]string{}
	for y := 0; y < maxY; y++ {
		visitedGrid = append(visitedGrid, []string{})
		for x := 0; x < maxX; x++ {
			visitedGrid[y] = append(visitedGrid[y], ".")
		}
	}

	for _, v := range visited {
		// fmt.Println(v)
		from := v.from
		visitedGrid[v.y][v.x] = fmt.Sprintf("%v (x%v y%v)", v.elevation, from.x, from.y)
	}

	fmt.Printf("-,")
	for i := range visitedGrid[0] {
		fmt.Printf("%v,", i)
	}
	fmt.Printf("\n")
	for i, row := range visitedGrid {
		fmt.Printf("%v,", i)
		for _, square := range row {
			fmt.Printf("%v,", square)
		}
		fmt.Printf("\n")
	}
}

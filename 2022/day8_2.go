package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello, world!")
	input := []string{}
	f, err := os.Open("./day8_1.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// Save as a grid (slice of slices of ints) and generate a 'scenic grid'
	grid := [][]int{}
	scenicGrid := [][]int{}
	for i, row := range input {
		grid = append(grid, []int{})
		scenicGrid = append(scenicGrid, []int{})
		for _, tree := range row {
			treeHeight, _ := strconv.Atoi(string(tree))
			grid[i] = append(grid[i], treeHeight)
			scenicGrid[i] = append(scenicGrid[i], 0)
		}
	}
	// Print the grids
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("")

	// For each tree, calculate it's scenic score
	gridHeight := len(grid)
	gridWidth := len(grid[0])
	for i := 0; i < gridHeight; i++ {
		for j := 0; j < gridWidth; j++ {
			scenicGrid[i][j] = scenicScore(grid, i, j)
		}
	}

	// fmt.Println("")
	for _, row := range scenicGrid {
		fmt.Println(row)
	}
	fmt.Println("")

	// find the tree with the highest scenic score
	bestTree := 0
	for i := 0; i < gridHeight; i++ {
		for j := 0; j < gridWidth; j++ {
			if scenicGrid[i][j] > bestTree {
				bestTree = scenicGrid[i][j]
			}
		}
	}
	fmt.Println(bestTree)

}

func scenicScore(grid [][]int, i int, j int) int {
	// up
	distanceUp := 0
	tmpJ := j
	for tmpI := i - 1; tmpI > -1; tmpI-- {
		distanceUp++
		if grid[tmpI][tmpJ] >= grid[i][j] {
			break
		}
	}

	// right
	distanceRight := 0
	tmpI := i
	for tmpJ := j + 1; tmpJ < len(grid[i]); tmpJ++ {
		distanceRight++
		if grid[tmpI][tmpJ] >= grid[i][j] {
			break
		}
	}

	// down
	distanceDown := 0
	tmpJ = j
	for tmpI := i + 1; tmpI < len(grid); tmpI++ {
		distanceDown++
		if grid[tmpI][tmpJ] >= grid[i][j] {
			break
		}
	}

	// left
	distanceLeft := 0
	tmpI = i
	for tmpJ := j - 1; tmpJ > -1; tmpJ-- {
		distanceLeft++
		if grid[tmpI][tmpJ] >= grid[i][j] {
			break
		}
	}

	// fmt.Printf("tree[%v][%v]: distanceUp = %v, distanceRight = %v, distanceDown = %v, disdistanceULeft = %v\n", i, j, distanceUp, distanceRight, distanceDown, distanceLeft)

	scenicScore := distanceUp * distanceRight * distanceDown * distanceLeft
	return scenicScore
}

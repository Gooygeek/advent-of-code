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

	// Save as a grid (slice of slices of ints) and generate a 'visible grid'
	grid := [][]int{}
	visibleGrid := [][]int{}
	for i, row := range input {
		grid = append(grid, []int{})
		visibleGrid = append(visibleGrid, []int{})
		for _, tree := range row {
			treeHeight, _ := strconv.Atoi(string(tree))
			grid[i] = append(grid[i], treeHeight)
			visibleGrid[i] = append(visibleGrid[i], 0)
		}
	}
	// Print the grids
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("")

	// For each of the 4 directions, iterate along that row/column and find the highest tree
	// Each time a new max is found, mark it as a 1 in the visibleGrid

	gridHeight := len(grid)
	gridWidth := len(grid[0])

	// Looking from the Top
	for i := 0; i < gridWidth; i++ {
		max := -1
		for j := 0; j < gridHeight; j++ {
			curTree := grid[j][i]
			if curTree > max {
				max = curTree
				visibleGrid[j][i] = 1
			}
		}
	}

	// Looking from the Right
	for i := 0; i < gridHeight; i++ {
		max := -1
		for j := gridWidth - 1; j > -1; j-- {
			curTree := grid[i][j]
			if curTree > max {
				max = curTree
				visibleGrid[i][j] = 1
			}
		}
	}

	// Looking from the Bottom
	for i := gridWidth - 1; i > -1; i-- {
		max := -1
		for j := gridHeight - 1; j > -1; j-- {
			curTree := grid[j][i]
			if curTree > max {
				max = curTree
				visibleGrid[j][i] = 1
			}
		}
	}

	// Looking from the Left
	for i := gridHeight - 1; i > -1; i-- {
		max := -1
		for j := 0; j < gridWidth; j++ {
			curTree := grid[i][j]
			if curTree > max {
				max = curTree
				visibleGrid[i][j] = 1
			}
		}
	}

	// fmt.Println("")
	for _, row := range visibleGrid {
		fmt.Println(row)
	}
	fmt.Println("")

	// Count the number of 1's in the visibleGrid to calculate the number of visible trees
	sum := 0
	for _, row := range visibleGrid {
		for _, tree := range row {
			sum = sum + tree
		}
	}
	fmt.Println(sum)

}

package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, world!")
	input := []string{}
	f, err := os.Open("./day9_1.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// For each step, calculate the position of the head relative to the starting point,
	// use this to then calculate the position of the tail relative to the starting point
	// positions are saved as a pair of [x, y] coords
	numOfTails := 9
	headPosList := [][]int{[]int{0, 0}}
	tailPosList := [][][]int{}
	for i := 0; i < numOfTails; i++ {
		tailPosList = append(tailPosList, [][]int{[]int{0, 0}})
	}
	posUpdateIndex := 0
	for _, step := range input {
		fmt.Println(step)
		stepSplit := strings.Split(step, " ")
		direction := stepSplit[0]
		distance, _ := strconv.Atoi(stepSplit[1])
		for j := 0; j < distance; j++ {
			if direction == "U" {
				headX := headPosList[posUpdateIndex][0]
				headY := headPosList[posUpdateIndex][1] + 1
				headPosList = append(headPosList, []int{headX, headY})
			}
			if direction == "D" {
				headX := headPosList[posUpdateIndex][0]
				headY := headPosList[posUpdateIndex][1] - 1
				headPosList = append(headPosList, []int{headX, headY})
			}
			if direction == "R" {
				headX := headPosList[posUpdateIndex][0] + 1
				headY := headPosList[posUpdateIndex][1]
				headPosList = append(headPosList, []int{headX, headY})
			}
			if direction == "L" {
				headX := headPosList[posUpdateIndex][0] - 1
				headY := headPosList[posUpdateIndex][1]
				headPosList = append(headPosList, []int{headX, headY})
			}
			for i := range tailPosList {
				var tailPos []int
				if i == 0 {
					tailPos = calcNewTailPos(headPosList, tailPosList[i])
				} else {
					tailPos = calcNewTailPos(tailPosList[i-1], tailPosList[i])
				}
				tailPosList[i] = append(tailPosList[i], tailPos)
			}
			posUpdateIndex++

			// fmt.Printf("%v, %v\n", headPosList[posUpdateIndex], tailPosList)
		}
	}

	// Using the list of tail positions, deduplicate and get the length to determine the number of spots the tail has been in at least once
	dedupedTailPos := removeDuplicateValues(tailPosList[len(tailPosList)-1])
	fmt.Println(len(dedupedTailPos))

	// DEBUGGING
	// get max size of grid
	// fmt.Println(dedupedTailPos)
	maxX := 0
	minX := 0
	maxY := 0
	minY := 0
	for i := range headPosList {
		if headPosList[i][0] > maxX {
			maxX = headPosList[i][0]
		}
		if headPosList[i][1] > maxY {
			maxY = headPosList[i][1]
		}
		if headPosList[i][0] < minX {
			minX = headPosList[i][0]
		}
		if headPosList[i][1] < minY {
			minY = headPosList[i][1]
		}
	}
	fmt.Printf("maxX: %v, minX: %v, maxY: %v, minY: %v\n", maxX, minX, maxY, minY)

	traceGrid := [][]int{}
	for y := maxY - 1; y > minY-1; y-- {
		traceGrid = append(traceGrid, []int{})
		for x := minX; x < maxX; x++ {
			posMarker := 0
			for pos := 0; pos < len(dedupedTailPos); pos++ {
				posString := fmt.Sprint(x) + "," + fmt.Sprint(y)
				if posString == dedupedTailPos[pos] {
					posMarker = 1
				}
			}
			traceGrid[len(traceGrid)-1] = append(traceGrid[len(traceGrid)-1], posMarker)
		}
	}
	// for _, row := range traceGrid {
	// 	fmt.Println(row)
	// }

	convert_to_image(traceGrid)

}

func calcNewTailPos(headPosList [][]int, tailPosList [][]int) []int {
	headPos := headPosList[len(headPosList)-1]
	headX := float64(headPos[0])
	headY := float64(headPos[1])
	oldTailPos := tailPosList[len(tailPosList)-1]
	oldTailX := float64(oldTailPos[0])
	oldTailY := float64(oldTailPos[1])
	newTailX := oldTailPos[0]
	newTailY := oldTailPos[1]
	deltaX := headX - oldTailX
	deltaY := headY - oldTailY
	if math.Abs(deltaX) >= 2 {
		newTailX = int(oldTailX + math.Copysign(1, deltaX))
		newTailY = int(headY)
	}
	if math.Abs(deltaY) >= 2 {
		newTailY = int(oldTailY + math.Copysign(1, deltaY))
		newTailX = int(headX)
	}
	if math.Abs(deltaX) >= 2 && math.Abs(deltaY) >= 2 {
		newTailX = int(oldTailX + math.Copysign(1, deltaX))
		newTailY = int(oldTailY + math.Copysign(1, deltaY))
	}

	return []int{newTailX, newTailY}
}

func removeDuplicateValues(intSlice [][]int) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		stringEntry := fmt.Sprint(entry[0]) + "," + fmt.Sprint(entry[1])
		if _, value := keys[stringEntry]; !value {
			keys[stringEntry] = true
			list = append(list, stringEntry)
		}
	}
	return list
}

func convert_to_image(grid [][]int) {
	// Create a new grayscale image with the given dimensions.
	var width, height = len(grid[0]), len(grid)
	img := image.NewGray(image.Rect(0, 0, width, height))

	// Iterate over the pixels in the image.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// Set the color of the pixel based on the value in the grid.
			var c color.Gray
			switch grid[y][x] {
			case 0:
				c = color.Gray{0}
			case 1:
				c = color.Gray{255}
			}
			img.SetGray(x, y, c)
		}
	}

	// Save the image to a PNG file.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

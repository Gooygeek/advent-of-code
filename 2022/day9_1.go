package main

import (
	"bufio"
	"fmt"
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
	headPosList := [][]int{[]int{0, 0}}
	tailPosList := [][]int{[]int{0, 0}}
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
			tailPos := calcNewTailPos(headPosList, tailPosList)
			tailPosList = append(tailPosList, tailPos)
			posUpdateIndex++

			fmt.Printf("%v, %v\n", headPosList[posUpdateIndex], tailPosList[posUpdateIndex])
		}
	}

	// Using the list of tail positions, deduplicate and get the length to determine the number of spots the tail has been in at least once
	dedupedTailPos := removeDuplicateValues(tailPosList)
	fmt.Println(len(dedupedTailPos))

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

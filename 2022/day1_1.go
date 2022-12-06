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
	f, err := os.Open("./day1_1.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	elves := []int{}
	currentElfSum := 0
	currentMax := 0
	for i := range input {
		fmt.Println(input[i])
		if input[i] != "" {
			intVar, err := strconv.Atoi(input[i])
			if err != nil {
				panic(err)
			}
			currentElfSum += intVar
		} else {
			elves = append(elves, currentElfSum)
			if currentElfSum > currentMax {
				currentMax = currentElfSum
			}
			currentElfSum = 0
		}
	}
	fmt.Println(elves)
	fmt.Println(currentMax)

}

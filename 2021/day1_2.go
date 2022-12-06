package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := []int{}
	f, err := os.Open("./day1_1.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		intVar, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		input = append(input, intVar)
	}

	increaseCount := 0
	prevDepth := input[0] + input[1] + input[2]
	for i := range input[:len(input)-2] {
		sum := input[i] + input[i+1] + input[i+2]
		if sum > prevDepth {
			fmt.Printf("%d (increase)\n", input[i])
			increaseCount++
		} else {
			fmt.Printf("%d\n", input[i])
		}
		prevDepth = sum
	}
	fmt.Println(increaseCount)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, world!")
	input := []string{}
	f, err := os.Open("./day10_1.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// For each step, keep track of the value of the register.
	// If the step is noop, then append the value of the register to the tracker.
	// If the step is a addx, then first append the value of the register to the tracker and then append the updated value to the register.

	registerX := 1
	registerXTracker := []int{1}

	for _, v := range input {
		step := strings.Split(v, " ")
		if step[0] == "noop" {
			registerXTracker = append(registerXTracker, registerX)
		}
		if step[0] == "addx" {
			registerXTracker = append(registerXTracker, registerX)
			value, _ := strconv.Atoi(step[1])
			registerX = registerX + value
			registerXTracker = append(registerXTracker, registerX)
		}
	}

	for _, v := range registerXTracker {
		fmt.Println(v)
	}

	// Add the appropriate signals
	// Note: the tracker is zero-indexed, but the signal is 'during' a cycle and thus takes the result of the previous cycle, but we also force the first cycle to have a value of 1
	// thus the '-1'
	cycle20 := registerXTracker[20-1] * 20
	cycle60 := registerXTracker[60-1] * 60
	cycle100 := registerXTracker[100-1] * 100
	cycle140 := registerXTracker[140-1] * 140
	cycle180 := registerXTracker[180-1] * 180
	cycle220 := registerXTracker[220-1] * 220
	result := cycle20 + cycle60 + cycle100 + cycle140 + cycle180 + cycle220
	fmt.Println(result)

}

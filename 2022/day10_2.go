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

	// For each cycle, determine if the value of the register is within 1 space away from the cycle count.
	// If so, draw '#', else draw '.'
	screen := []string{}
	for i, v := range registerXTracker {
		rowCycle := math.Mod(float64(i), 40)
		if math.Abs(float64(v)-rowCycle) <= 1 {
			screen = append(screen, "#")
		} else {
			screen = append(screen, " ")
		}
	}
	for i := range screen {
		fmt.Printf("%v", screen[i])
		if math.Mod(float64(i)+1, 40) == 0 {
			fmt.Printf("\n")
		}
	}

}

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
	f, err := os.Open("./day4_1.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	runningTotal := 0

	for _, v := range input {
		// split in to each elf
		elfPair := strings.Split(v, ",")

		// define starting and ending sections for each elf
		elf1_strings := strings.Split(elfPair[0], "-")
		elf1_start, _ := strconv.Atoi(elf1_strings[0])
		elf1_end, _ := strconv.Atoi(elf1_strings[len(elf1_strings)-1])
		elf2_strings := strings.Split(elfPair[1], "-")
		elf2_start, _ := strconv.Atoi(elf2_strings[0])
		elf2_end, _ := strconv.Atoi(elf2_strings[len(elf2_strings)-1])
		fmt.Printf("elf1_start = %v, elf1_end = %v, elf2_start = %v, elf2_end = %v\n", elf1_start, elf1_end, elf2_start, elf2_end)

		//determine if the end section on elf 2 isn't above end section of elf 1 OR ending section of elf 2 isn't below starting section of elf 2
		if !((elf2_start > elf1_end) || (elf2_end < elf1_start)) {
			fmt.Println("Overlap!!")
			runningTotal += 1
		}
	}

	fmt.Printf("Total = %d\n", runningTotal)
}

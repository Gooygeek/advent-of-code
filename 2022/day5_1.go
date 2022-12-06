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
	f, err := os.Open("./day5_1.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// First, get the number of stacks

	lineOfStackCount := 0
	stacks := [][]string{}
	for i, v := range input {
		if string(v[:2]) == " 1" {
			stackStrings := strings.Split(v, " ")
			stackStrings = delete_empty(stackStrings)
			stackCount := len(stackStrings)
			for i := 0; i < stackCount; i++ {
				stacks = append(stacks, []string{})
			}
			lineOfStackCount = i
			break
		}
	}

	// Now populate the stacks
	for i := lineOfStackCount - 1; i >= 0; i-- {
		v := string(input[i])
		if string(v[:2]) == " 1" {
			break
		}
		for j := 0; j < len(v); j = j + 4 {
			if v[j] == '[' {
				stacks[j/4] = push(stacks[j/4], string(v[j+1]))
			}
		}
	}
	fmt.Println(stacks)

	// for each instruction, parse out the number of crates to move from which stack to which stack
	// execute by using pop & push stack operations between the designated stacks, X amount of times

	for i := lineOfStackCount + 2; i < len(input); i++ {
		v := input[i]
		fmt.Println(v)
		instructionsSplit := strings.Split(v, " ")
		moveCount, _ := strconv.Atoi(instructionsSplit[1])
		fromStack, _ := strconv.Atoi(instructionsSplit[3])
		fromStack-- // zero-index
		toStack, _ := strconv.Atoi(instructionsSplit[5])
		toStack-- // zero-index
		for j := 0; j < moveCount; j++ {
			item := ""
			item, stacks[fromStack] = pop(stacks[fromStack])
			stacks[toStack] = push(stacks[toStack], item)
		}
		fmt.Println(stacks)
	}

	// get top of each stack
	for i := range stacks {
		fmt.Printf("%v", stacks[i][len(stacks[i])-1])
	}
	fmt.Printf("\n")
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func pop(s []string) (string, []string) {
	lastItem := s[len(s)-1]
	remaining := s[0 : len(s)-1]
	return lastItem, remaining
}

func push(s []string, input string) []string {
	return append(s, input)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world!")
	input := []string{}
	f, err := os.Open("./day3_1.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	letterValueMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}

	runningTotal := 0
	for _, v := range input {
		duplicateItem := 'A'
		length := len(v)
		compartment1 := v[:length/2]
		compartment2 := v[length/2:]
		fmt.Printf("total: %v = %v, comp1: %v =  %v, comp2: %v =  %v\n", v, length, compartment1, len(compartment1), compartment2, len(compartment2))
		for _, comp1Item := range compartment1 {
			for _, comp2Item := range compartment2 {
				if comp1Item == comp2Item {
					duplicateItem = comp1Item
				}
			}
		}
		priority := letterValueMap[string(duplicateItem)]
		runningTotal += priority
		fmt.Printf("Duplicate found! %v, number: %v, priority: %v, running total: %d\n", string(duplicateItem), duplicateItem, priority, runningTotal)
	}
	fmt.Printf("Total = %d\n", runningTotal)
}

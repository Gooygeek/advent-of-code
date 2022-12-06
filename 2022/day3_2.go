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
	fmt.Println(letterValueMap)

	runningTotal := 0

	i := 0
	for i < len(input) {
		sack1 := input[i]
		sack2 := input[i+1]
		sack3 := input[i+2]

		fmt.Printf("sack1 = %v, sack2 = %v, sack3 = %v\n", sack1, sack2, sack3)
		duplicateFound := false
		for _, sack1Item := range sack1 {
			for _, sack2Item := range sack2 {
				if sack1Item == sack2Item {
					for _, sack3Item := range sack3 {
						if sack1Item == sack3Item {
							runningTotal += letterValueMap[string(sack1Item)]
							fmt.Printf("Duplicate item = %v, value = %v, running total = %v\n", string(sack1Item), letterValueMap[string(sack1Item)], runningTotal)
							duplicateFound = true
						}
						if duplicateFound {
							break
						}
					}
				}
				if duplicateFound {
					break
				}
			}
			if duplicateFound {
				break
			}
		}

		i += 3
	}

	fmt.Printf("Total = %d\n", runningTotal)
}

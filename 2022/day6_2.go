package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, world!")
	input := []string{}
	f, err := os.Open("./day6_1.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	stream := input[0]
	// generate window 4 chars wide
	for i := 0; i < len(stream)-13; i++ {
		substream := stream[i : i+14]

		// compare all chars are different
		for j := 0; j < len(substream); j++ {
			fmt.Printf("Looking for %v in '%v'\n", string(substream[j]), substream[j+1:])
			if len(substream[j+1:]) == 0 {
				return
			}
			if !(strings.Contains(substream[j+1:], string(substream[j]))) {
				// report input index at end of relevant window
				fmt.Println(i + 14)
			} else {
				fmt.Println("Duplicate detected!")
				break
			}
		}
	}
}

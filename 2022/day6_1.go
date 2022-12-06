package main

import (
	"bufio"
	"fmt"
	"os"
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
	for i := 0; i < len(stream); i++ {
		a := stream[i]
		b := stream[i+1]
		c := stream[i+2]
		d := stream[i+3]

		// compare all chars are different
		if !(a == b) && !(a == c) && !(a == d) && !(b == c) && !(b == d) && !(c == d) {
			fmt.Println(i + 4)
			break
		}

		// report input index at end of relevant window
	}
}

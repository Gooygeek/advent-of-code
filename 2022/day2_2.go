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
	f, err := os.Open("./day2_1.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	shapeMapper := map[string]map[string]string{
		"A": {"X": "C", "Z": "B", "Y": "A", "points": "1"},
		"B": {"X": "A", "Z": "C", "Y": "B", "points": "2"},
		"C": {"X": "B", "Z": "A", "Y": "C", "points": "3"},
	}
	outcomeMapper := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	runningTotalPoints := 0

	for _, v := range input {
		opponentShape := string(v[0])
		desiredOutcome := string(v[2])

		fmt.Printf("Your opponent played %s, you need to %s, therefore you need to play %s\n", opponentShape, desiredOutcome, shapeMapper[opponentShape][desiredOutcome])
		shapePoints, _ := strconv.Atoi(shapeMapper[shapeMapper[opponentShape][desiredOutcome]]["points"])
		outcomePoints := outcomeMapper[desiredOutcome]
		roundPoints := shapePoints + outcomePoints
		fmt.Printf("You score: %d points\n", roundPoints)
		runningTotalPoints += roundPoints
	}
	fmt.Println(runningTotalPoints)
}

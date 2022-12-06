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
		"A": {"beats": "Z", "points": "1"},
		"B": {"beats": "X", "points": "2"},
		"C": {"beats": "Y", "points": "3"},
		"X": {"beats": "C", "points": "1"},
		"Y": {"beats": "A", "points": "2"},
		"Z": {"beats": "B", "points": "3"},
	}

	runningTotalPoints := 0

	for _, v := range input {
		// fmt.Println(v)
		opponentShape := string(v[0])
		yourShape := string(v[2])
		fmt.Printf("Opponent: %s (%s), You: %s (%s)\n", opponentShape, shapeMapper[opponentShape], yourShape, shapeMapper[yourShape])
		if shapeMapper[yourShape]["beats"] == opponentShape {
			shapePoints, _ := strconv.Atoi(shapeMapper[yourShape]["points"])
			roundPoints := shapePoints + 6
			runningTotalPoints = runningTotalPoints + roundPoints
			fmt.Printf("You win! You scored: %v!\n", roundPoints)
		} else if shapeMapper[opponentShape]["beats"] == yourShape {
			shapePoints, _ := strconv.Atoi(shapeMapper[yourShape]["points"])
			roundPoints := shapePoints + 0
			runningTotalPoints = runningTotalPoints + roundPoints
			fmt.Printf("You lose! You scored: %v\n", roundPoints)
		} else {
			shapePoints, _ := strconv.Atoi(shapeMapper[yourShape]["points"])
			roundPoints := shapePoints + 3
			runningTotalPoints = runningTotalPoints + roundPoints
			fmt.Printf("Draw, You scored: %v!\n", roundPoints)
		}
	}
	fmt.Println(runningTotalPoints)
}

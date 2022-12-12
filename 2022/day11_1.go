package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	number          string
	items           []int
	consideredItems int
	operation       string
	operationSize   string
	test            int
	nextIfTrue      string
	nextIfFalse     string
}

func main() {
	fmt.Println("Hello, world!")
	input := []string{}
	f, err := os.Open("./day11_1.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// Load each money into a struct
	monkeys := map[string]Monkey{} // order is not guaranteed with maps
	monkeyList := []string{}       // This allows us to use a map for easier reference while also maintaining ordering
	curMonkey := "0"
	for _, v := range input {
		row := strings.Split(v, " ")
		fmt.Println(row)
		if len(v) != 0 {
			if row[0] == "Monkey" {
				monkeyNumber := strings.Trim(row[1], ":")
				monkey := Monkey{number: monkeyNumber}
				monkeys[monkeyNumber] = monkey
				curMonkey = monkeyNumber
				monkeyList = append(monkeyList, curMonkey)
			} else if row[2] == "Starting" {
				startingItemsString := row[4:]
				startingItems := []int{}
				for i := range startingItemsString {
					item, _ := strconv.Atoi(strings.Trim(startingItemsString[i], ","))
					startingItems = append(startingItems, item)
				}
				tmpMonkey := monkeys[curMonkey]
				tmpMonkey.items = startingItems
				monkeys[curMonkey] = tmpMonkey
			} else if row[2] == "Operation:" {
				tmpMonkey := monkeys[curMonkey]
				tmpMonkey.operation = row[6]
				tmpMonkey.operationSize = row[7]
				monkeys[curMonkey] = tmpMonkey
			} else if row[2] == "Test:" {
				tmpMonkey := monkeys[curMonkey]
				tmpMonkey.test, _ = strconv.Atoi(row[5])
				monkeys[curMonkey] = tmpMonkey
			} else if row[5] == "true:" {
				tmpMonkey := monkeys[curMonkey]
				tmpMonkey.nextIfTrue = row[9]
				monkeys[curMonkey] = tmpMonkey
			} else if row[5] == "false:" {
				tmpMonkey := monkeys[curMonkey]
				tmpMonkey.nextIfFalse = row[9]
				monkeys[curMonkey] = tmpMonkey
			}
		}
	}
	for _, v := range monkeys {
		fmt.Println(v)
	}
	fmt.Println(monkeyList)

	// For 20 rounds, compute the stuff-slinging, keeping track of the number of items considered each round per monkey
	for i := 0; i < 20; i++ {
		// for each round
		fmt.Printf("Round %d\n", i)
		for j := 0; j < len(monkeyList); j++ {
			// for each monkey
			curMonkey := monkeys[monkeyList[j]]
			fmt.Println(curMonkey)
			curMonkey.consideredItems += len(curMonkey.items)
			// for each item
			for _, item := range curMonkey.items {
				fmt.Printf("Item Start: %v\n", item)
				// calc the scaler used in the operation
				operationSize := 0
				if curMonkey.operationSize == "old" {
					operationSize = item
				} else {
					operationSize, _ = strconv.Atoi(curMonkey.operationSize)
				}

				// perform the operation
				if curMonkey.operation == "+" {
					item = item + operationSize
				} else if curMonkey.operation == "*" {
					item = item * operationSize
				}

				// post operation division
				item = item / 3 // integer division, rounds down
				fmt.Printf("Item End: %v\n", item)

				// run the test
				if math.Mod(float64(item), float64(curMonkey.test)) == 0 {
					tmpMonkey := monkeys[curMonkey.nextIfTrue]
					tmpMonkey.items = append(tmpMonkey.items, item)
					monkeys[curMonkey.nextIfTrue] = tmpMonkey
				} else {
					tmpMonkey := monkeys[curMonkey.nextIfFalse]
					tmpMonkey.items = append(tmpMonkey.items, item)
					monkeys[curMonkey.nextIfFalse] = tmpMonkey
				}

				// Remove Item from curMonkey
				curMonkey.items = curMonkey.items[1:]
			}
			// Save curMonkey back to map
			monkeys[monkeyList[j]] = curMonkey
		}
	}

	for _, v := range monkeyList {
		fmt.Println(monkeys[v])
	}

	// Calculate the result
	monkeyConsiderItemsCounts := []int{}
	for _, v := range monkeyList {
		monkeyConsiderItemsCounts = append(monkeyConsiderItemsCounts, monkeys[v].consideredItems)
	}
	fmt.Println(monkeyConsiderItemsCounts)
	sort.Ints(monkeyConsiderItemsCounts)
	fmt.Println(monkeyConsiderItemsCounts)
	result := monkeyConsiderItemsCounts[len(monkeyConsiderItemsCounts)-2] * monkeyConsiderItemsCounts[len(monkeyConsiderItemsCounts)-1]
	fmt.Println(result)
}

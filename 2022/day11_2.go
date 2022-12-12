package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	number          string
	items           []big.Int
	consideredItems int
	operation       string
	operationSize   string
	test            big.Int
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
		// fmt.Println(row)
		if len(v) != 0 {
			if row[0] == "Monkey" {
				monkeyNumber := strings.Trim(row[1], ":")
				monkey := Monkey{number: monkeyNumber}
				monkeys[monkeyNumber] = monkey
				curMonkey = monkeyNumber
				monkeyList = append(monkeyList, curMonkey)
			} else if row[2] == "Starting" {
				startingItemsString := row[4:]
				startingItems := []big.Int{}
				for i := range startingItemsString {
					item, _ := strconv.ParseInt(strings.Trim(startingItemsString[i], ","), 10, 64)
					bigItem := big.NewInt(item)
					startingItems = append(startingItems, *bigItem)
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
				test, _ := strconv.ParseInt(row[5], 10, 64)
				tmpMonkey.test = *big.NewInt(test)
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

	// For 10000 rounds, compute the stuff-slinging, keeping track of the number of items considered each round per monkey
	worryDecreaser := big.NewInt(1) // divide an item by this amount each time it is inspected
	// modulo the item by the product of all tests after each time it is inspected.
	//   This works because, if y = n * m, then x % y preserves x % n and x % m.
	//   It means that x can be decreased (thus not get too big) while preserving the modulo relationship between x, n, and m.
	moduloser := big.NewInt(1)
	for _, v := range monkeyList {
		test := monkeys[v].test
		moduloser.Mul(moduloser, &test)
	}
	for i := 0; i < 10000; i++ {
		monkeys = calcRound(monkeys, monkeyList, *worryDecreaser, *moduloser, i)
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

func calcRound(monkeys map[string]Monkey, monkeyList []string, worryDecreaser big.Int, moduloser big.Int, i int) map[string]Monkey {
	// for each round
	fmt.Printf("Round %d\n", i)
	for j := 0; j < len(monkeyList); j++ {
		// for each monkey
		curMonkey := monkeys[monkeyList[j]]
		// fmt.Println(curMonkey)
		curMonkey.consideredItems += len(curMonkey.items)
		// for each item
		for _, item := range curMonkey.items {
			// fmt.Printf("Item Start: %v\n", item)
			// calc the scaler used in the operation
			var operationSize big.Int
			if curMonkey.operationSize == "old" {
				operationSize = item
			} else {
				size, _ := strconv.ParseInt(curMonkey.operationSize, 10, 64)
				operationSize = *big.NewInt(size)
			}

			// perform the operation
			if curMonkey.operation == "+" {
				item.Add(&item, &operationSize)
			} else if curMonkey.operation == "*" {
				item.Mul(&item, &operationSize)
			}

			// post operation division
			item = *item.Div(&item, &worryDecreaser) // integer division, rounds down
			item.Mod(&item, &moduloser)
			// fmt.Printf("Item End: %v\n", item)

			// run the test
			var modRes big.Int
			modRes.Mod(&item, &curMonkey.test)
			var zero int64 = 0
			if modRes.Int64() == zero {
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
	return monkeys
}

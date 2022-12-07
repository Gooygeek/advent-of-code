package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	name  string
	dirs  map[string]Directory
	files map[string]int
	size  int
}

func main() {
	fmt.Println("Hello, world!")
	input := []string{}
	f, err := os.Open("./day7_1.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// generate a map of the file system based onn the input
	root, _ := cd("/", input[1:])
	fmt.Printf("root: %v\n", root)

	fmt.Println("-")
	fmt.Println("-")

	// iterate through all directories and calculate the smallest folder that's above the minimum required size
	min := 30000000 - (70000000 - root.size)
	smallestDir := tree(root, 4, map[string]string{"name": "/", "size": strconv.Itoa(root.size)}, min)
	fmt.Printf("running total: %v\n", smallestDir)

	// Sum the sizes of directories above 100,000

}

func cd(dir string, input []string) (Directory, int) {
	fmt.Printf("traversing into %v\n", dir)
	curDir := Directory{name: dir, dirs: map[string]Directory{}, files: map[string]int{}, size: 0}
	for i := 0; i < len(input); i++ {
		v := input[i]
		lineBreakdown := strings.Split(v, " ")
		fmt.Printf("LineBreakdown = %v\n", lineBreakdown)
		if lineBreakdown[0] == "$" && lineBreakdown[1] == "cd" {
			if lineBreakdown[2] != ".." {
				traversedDir, returnedIndex := cd(lineBreakdown[2], input[i+1:])
				fmt.Printf("Traversed Dir: %v\n", traversedDir)
				curDir.dirs[lineBreakdown[2]] = traversedDir
				i = i + returnedIndex + 1
				curDir.size += traversedDir.size
			} else if lineBreakdown[2] == ".." {
				fmt.Printf("Going back up, returning index: %v\n", i)
				return curDir, i
			}
		} else if lineBreakdown[0] == "dir" {
			curDir.dirs[lineBreakdown[1]] = Directory{name: lineBreakdown[1], dirs: map[string]Directory{}, files: map[string]int{}, size: 0}
		} else if lineBreakdown[0] != "$" {
			curDir.files[lineBreakdown[1]], _ = strconv.Atoi(lineBreakdown[0])
			fileSize, _ := strconv.Atoi(lineBreakdown[0])
			curDir.size += fileSize
		}
	}
	return curDir, len(input)
}

func tree(curDir Directory, indentSize int, curSmallest map[string]string, min int) map[string]string {
	indent := ""
	for i := 0; i < indentSize-4; i++ {
		indent += " "
	}
	sizeMarker := ""
	fmt.Printf("%v%v - %v%v\n", indent, curDir.name, curDir.size, sizeMarker)
	indent += "    "
	for i := range curDir.files {
		fmt.Printf("%v%v %v\n", indent, curDir.files[i], i)
	}
	for i := range curDir.dirs {
		curSmallest = tree(curDir.dirs[i], indentSize+4, curSmallest, min)
	}
	curSmallestSize, _ := strconv.Atoi(curSmallest["size"])
	if curDir.size < curSmallestSize && curDir.size > min {
		return map[string]string{"name": curDir.name, "size": strconv.Itoa(curDir.size)}
	}
	return curSmallest
}

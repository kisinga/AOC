package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	fileArray := []string{}
	for scanner.Scan() {
		fileArray = append(fileArray, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slope1 := checkTrees(fileArray, 1, 1)
	slope2 := checkTrees(fileArray, 3, 1)
	slope3 := checkTrees(fileArray, 5, 1)
	slope4 := checkTrees(fileArray, 7, 1)
	slope5 := checkTrees(fileArray, 1, 2)
	fmt.Println(slope1 * slope2 * slope3 * slope4 * slope5)
}

func checkTrees(input []string, xoffset int, yoffset int) int {
	trees := 0

	for y, x := 0, 0; y < len(input); y, x = y+yoffset, x+xoffset {
		fmt.Println(string(input[y]), x, len(input[y]), x%len(input[y]), y, string(input[y][x%len(input[y])]))
		if string(input[y][x%len(input[y])]) == "#" {
			trees++
		}
	}

	return trees
}

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
	treeCount := checkTrees(fileArray, 3, 1)
	fmt.Println(treeCount)
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

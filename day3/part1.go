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
	offset := 0
	treeCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if checkTree(offset, []byte(line)) {
			treeCount++
		}
		if offset+3 > len(line)-1 {
			x := (len(line) - offset)
			offset = 3 - x
			continue
		}
		offset += 3
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(treeCount)
}
func checkTree(offset int, line []byte) bool {
	found := false
	if line[offset] == '#' {
		found = true
		line[offset] = 'X'
	} else {
		line[offset] = 'O'
	}
	fmt.Println(offset, string(line))

	return found
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var mandatoryKeys []string = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	count := countYes(string(content))
	// count = countYes("abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb")
	fmt.Println(count)
}

func countYes(answers string) int {
	total := 0
	for _, answers := range strings.Split(answers, "\n\n") {
		found := map[rune]bool{}
		fields := strings.Fields(answers)
		for _, field := range fields {
			for _, answer := range field {
				found[answer] = true
			}
		}
		total += len(found)
	}
	return total
}

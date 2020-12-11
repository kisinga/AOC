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
	for _, groupAnswers := range strings.Split(answers, "\n\n") {
		totalFound := map[rune]int{}
		singlePersonAnsers := strings.Fields(groupAnswers)
		memberCount := len(singlePersonAnsers)

		for _, answers := range singlePersonAnsers {
			yeses := map[rune]bool{}
			for _, answer := range answers {
				yeses[answer] = true
			}
			for key := range yeses {
				totalFound[key]++
			}
		}

		for _, yesCount := range totalFound {
			if yesCount == memberCount {
				total++
			}
		}
	}
	return total
}

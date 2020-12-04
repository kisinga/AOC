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
	passports := strings.Split(string(content), "\n\n")
	validCount := 0
	for _, passport := range passports {
		if validatePassport(passport) {
			validCount++
		}
	}
	fmt.Println(validCount)
}
func validatePassport(passport string) bool {
	fields := strings.Fields(passport)
	found := map[string]bool{}
	valid := true
	for _, field := range fields {
		key := strings.Split(field, ":")[0]
		found[key] = true
	}
	for _, key := range mandatoryKeys {
		if !found[key] {
			valid = false
		}
	}
	return valid
}

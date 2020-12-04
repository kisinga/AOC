package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type heightType int

const (
	cm heightType = iota
	in heightType = iota
)

var heighRegex = regexp.MustCompile(`#([0-9a-f]){6}`)
var eyeColorRegex = regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
var passportIDRegex = regexp.MustCompile(`^[0-9]{9}$`)
var temp []string

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	passports := strings.Split(string(content), "\n\n")
	validCount := 0
	// fmt.Println(len(passports))
	for _, passport := range passports {
		if validatePassport(passport) {
			validCount++
		}
	}
	fmt.Println(validCount)
	fmt.Println(temp)
}
func validatePassport(passport string) bool {
	fields := strings.Fields(passport)
	validCount := 0
	for _, field := range fields {
		split := strings.Split(field, ":")
		key, value := split[0], split[1]
		switch key {
		case "byr":
			val, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			if val < 1920 || val > 2002 {
				return false
			}
			validCount++
			break

		case "iyr":
			val, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			if val < 2010 || val > 2020 {
				return false
			}
			validCount++
			break

		case "eyr":
			val, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			if val < 2020 || val > 2030 {
				return false
			}
			validCount++
			break

		case "hgt":
			hgtType, val, err := isolateHeight(value)
			if err != nil {
				return false
			}
			switch hgtType {
			case cm:
				if val < 150 || val > 193 {
					return false
				}
				break
			case in:
				if val < 59 || val > 76 {
					return false
				}
				break
			}
			validCount++
			break

		case "hcl":
			match := heighRegex.MatchString(value)
			if !match {
				return false
			}
			validCount++
			break

		case "ecl":
			match := eyeColorRegex.MatchString(value)
			if !match {
				return false
			}
			validCount++
			break

		case "pid":
			match := passportIDRegex.MatchString(value)
			if !match {
				temp = append(temp, value)
				return false
			}
			validCount++
			break

		default:
			// fmt.Println(key)
			break
		}

	}
	return validCount == 7
}

func isolateHeight(height string) (heightType, int, error) {
	if strings.Contains(height, "cm") {
		val, err := strconv.Atoi(strings.Trim(height, "cm"))
		if err != nil {
			return 0, 0, err
		}
		return cm, val, nil
	}
	val, err := strconv.Atoi(strings.Trim(height, "in"))
	if err != nil {
		return 0, 0, err
	}
	return in, val, nil
}

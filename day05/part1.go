package main

import (
	"bufio"
	"errors"
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
	highestID := 0
	fmt.Println(getSeatID("FBFBBFFRLR"))
	for scanner.Scan() {
		if seatID := getSeatID(scanner.Text()); seatID > highestID {
			highestID = seatID
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(highestID)
}
func getSeatID(line string) int {
	row, err := getRow(string([]byte(line)[0:7]))
	if err != nil {
		log.Fatal(err)
	}
	col, err := getCol(string([]byte(line)[7:10]))
	if err != nil {
		log.Fatal(err)
	}
	return (row * 8) + col

}
func getRow(val string) (int, error) {
	if len(val) != 7 {
		return 0, errors.New("row string must be exactly 7 letters")
	}
	upper := 127
	lower := 0
	for _, letter := range val {
		func(letter rune, lower *int, upper *int) {
			_range := *upper - *lower
			mid := (_range / 2) + 1
			switch letter {
			case 'F':
				*upper = *upper - mid
				break
			default:
				*lower = *lower + mid
				break
			}
		}(letter, &lower, &upper)
	}
	return upper, nil
}

func getCol(val string) (int, error) {
	if len(val) != 3 {
		return 0, errors.New("col string must be exactly 3 letters")
	}
	upper := 7
	lower := 0
	for _, letter := range val {
		func(letter rune, lower *int, upper *int) {
			_range := *upper - *lower
			mid := (_range / 2) + 1
			switch letter {
			case 'L':
				*upper = *upper - mid
				break
			default:
				*lower = *lower + mid
				break
			}
		}(letter, &lower, &upper)
	}
	return upper, nil
}

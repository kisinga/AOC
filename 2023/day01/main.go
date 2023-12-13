package day01

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var numberDict = map[string]int{
	"one":         1,
	"two":         2,
	"three":       3,
	"four":        4,
	"five":        5,
	"six":         6,
	"seven":       7,
	"eight":       8,
	"nine":        9,
	"ten":         10,
	"eleven":      11,
	"twelve":      12,
	"thirteen":    13,
	"fourteen":    14,
	"fifteen":     15,
	"sixteen":     16,
	"seventeen":   17,
	"eighteen":    18,
	"nineteen":    19,
	"twenty":      20,
	"thirty":      30,
	"forty":       40,
	"fifty":       50,
	"sixty":       60,
	"seventy":     70,
	"eighty":      80,
	"ninety":      90,
	"hundred":     100,
	"thousand":    1000,
	"million":     1000000,
	"billion":     1000000000,
	"trillion":    1000000000000,
	"quadrillion": 1000000000000000,
	"quintillion": 1000000000000000000,
}

type day struct {
	// store any necessary configs here
	dataDir string
	answer  struct {
		part1 int
		part2 int
	}
}

type containedNumber struct {
	index int
	value int
}

type numberWordsContainedInStringMap struct {
	// map of number words to their index in the string. The index is the index of the word, the value is the number
	smallestIndex *containedNumber
	largestIndex  *containedNumber
}

func (d *day) Solve() error {
	file, err := os.Open(d.dataDir)
	if err != nil {
		return err
	}
	defer file.Close()
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		twoDigits := make([]int, 0, 2)

		// Part 1
		for index, s := range line {
			if isNumber(s) {
				digit := int(s - '0')
				if len(twoDigits) == 2 {
					twoDigits = twoDigits[0:1]
				}
				twoDigits = append(twoDigits, digit)
			}
			if len(line) == index+1 {
				var lineDigits int
				if len(twoDigits) == 2 {
					lineDigits = combineDigits(twoDigits[0], twoDigits[1])
				} else if len(twoDigits) == 1 {
					lineDigits = combineDigits(twoDigits[0], twoDigits[0])
				}
				d.answer.part1 += lineDigits
				// fmt.Println(lineDigits, d.answer)
			}
		}

		// Part 2
		// reset twoDigits
		twoDigits = make([]int, 0, 2)

		// get the numbers contained in the string

		indexStore := new(numberWordsContainedInStringMap)

		for index, s := range line {
			if isNumber(s) {
				digit := int(s - '0')
				if len(twoDigits) == 2 {
					twoDigits = twoDigits[0:1]
				}
				if indexStore.smallestIndex == nil || index < indexStore.smallestIndex.index {
					indexStore.smallestIndex = &containedNumber{index: index, value: digit}
				}

				if indexStore.largestIndex == nil || index > indexStore.largestIndex.index {
					indexStore.largestIndex = &containedNumber{index: index, value: digit}
				}
			}

			if len(line) == index+1 {
				indexStore = compareIndexesWithDictIndexes(line, *indexStore)

				twoDigits = []int{indexStore.smallestIndex.value, indexStore.largestIndex.value}

				lineDigits := combineDigits(twoDigits[0], twoDigits[1])

				d.answer.part2 += lineDigits
				fmt.Println(lineDigits, d.answer)
			}

		}

	}
	return nil
}

func ensureSingleDigit(digit int) int {
	if digit > 9 {
		return digit % 10
	}
	return digit
}

func compareIndexesWithDictIndexes(s string, indexStore numberWordsContainedInStringMap) *numberWordsContainedInStringMap {
	for key, v := range numberDict {
		lastIndex := strings.LastIndex(s, key)
		firstIndex := strings.Index(s, key)

		if firstIndex != -1 {
			if indexStore.smallestIndex == nil || firstIndex < indexStore.smallestIndex.index {
				indexStore.smallestIndex = &containedNumber{
					index: firstIndex,
					value: ensureSingleDigit(v),
				}
			}
		}

		if lastIndex != -1 {
			if indexStore.largestIndex == nil || lastIndex > indexStore.largestIndex.index {
				indexStore.largestIndex = &containedNumber{
					index: lastIndex,
					value: ensureSingleDigit(v),
				}
			}
		}
	}
	return &indexStore
}

// since we're only dealing with two digits, we can just use this shortcut
func combineDigits(digit1, digit2 int) int {
	combined := digit1*10 + digit2
	return combined
}

func isNumber(s rune) bool {
	return s >= '0' && s <= '9'
}

func NewChallenge(dataDir string) *day {
	return &day{
		dataDir: dataDir,
	}
}

func Run() *cobra.Command {

	d := NewChallenge("day01/data.txt")

	cmd := &cobra.Command{
		Use: "1",
		Run: func(cmd *cobra.Command, args []string) {
			err := d.Solve()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Answer:", d.answer)
		},
	}
	return cmd
}

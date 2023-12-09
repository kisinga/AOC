package day01

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var numberMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
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
		wordIndexes := numberWordsContainedInString(line)

		digitIndexes := new(numberWordsContainedInStringMap)

		for index, s := range line {
			if isNumber(s) {
				digit := int(s - '0')
				if len(twoDigits) == 2 {
					twoDigits = twoDigits[0:1]
				}
				if digitIndexes.smallestIndex != nil {
					digitIndexes.smallestIndex = &containedNumber{
						index: index,
						value: digit,
					}
				} else if digitIndexes.largestIndex != nil {
					digitIndexes.largestIndex = &containedNumber{
						index: index,
						value: digit,
					}
				} else {
					if digitIndexes.largestIndex == nil {
						digitIndexes.largestIndex = &containedNumber{
							index: index,
							value: digit,
						}
					} else if index > digitIndexes.largestIndex.index {
						digitIndexes.largestIndex.index = digit
					}
				}
			}
			if wordIndexes.smallestIndex != nil && digitIndexes.smallestIndex != nil {
				if wordIndexes.smallestIndex.index < digitIndexes.smallestIndex.index {
					twoDigits = append(twoDigits, wordIndexes.smallestIndex.value)
				}
			} else if wordIndexes.smallestIndex != nil {
				twoDigits = append(twoDigits, wordIndexes.smallestIndex.value)
			} else if digitIndexes.smallestIndex != nil {
				twoDigits = append(twoDigits, digitIndexes.smallestIndex.value)
			}

			if wordIndexes.largestIndex != nil && digitIndexes.largestIndex != nil {
				if wordIndexes.largestIndex.index > digitIndexes.largestIndex.index {
					twoDigits = append(twoDigits, wordIndexes.largestIndex.value)
				}
			} else if wordIndexes.largestIndex != nil {
				twoDigits = append(twoDigits, wordIndexes.largestIndex.value)
			} else if digitIndexes.largestIndex != nil {
				twoDigits = append(twoDigits, digitIndexes.largestIndex.value)
			}
		}

	}
	return nil
}

func numberWordsContainedInString(s string) numberWordsContainedInStringMap {
	var numbersContained numberWordsContainedInStringMap
	for key, v := range numberMap {
		index := strings.Index(s, key)
		if index != -1 {
			if numbersContained.smallestIndex == nil {
				numbersContained.smallestIndex = &containedNumber{
					index: index,
					value: v,
				}
			} else if numbersContained.largestIndex == nil {
				numbersContained.largestIndex = &containedNumber{
					index: index,
					value: v,
				}
			} else {
				if numbersContained.largestIndex != nil {
					numbersContained.largestIndex = &containedNumber{
						index: index,
						value: v,
					}
				} else if index > numbersContained.largestIndex.index {
					numbersContained.largestIndex.index = index
				}
			}

		}
	}
	return numbersContained
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

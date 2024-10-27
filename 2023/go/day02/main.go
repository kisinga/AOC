package day02

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type day struct {
	// store any necessary configs here
	dataDir string
	answer  struct {
		part1 int
		part2 int
	}
}
type Set struct {
	Blue  int
	Red   int
	Green int
}

func (d *day) Solve(constraint Set) error {

	re := regexp.MustCompile(`Game \d+:`)
	gameIdRegex := regexp.MustCompile(`[^\d]+`)
	file, err := os.Open(d.dataDir)
	if err != nil {
		return err
	}

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	// Iterate through each line
	gameIDs := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		extractedGameId := string(re.Find([]byte(line)))

		gameIDString := gameIdRegex.ReplaceAllString(extractedGameId, "")

		gameID, err := strconv.Atoi(gameIDString)
		if err != nil {
			return err
		}

		// remove the game number
		strippedString := re.ReplaceAllString(line, "")

		g, err := processLine(strippedString, constraint)
		if err != nil {
			return err
		}
		if g {
			gameIDs = append(gameIDs, gameID)
		}
	}
	file.Close()
	totals := 0
	for _, v := range gameIDs {
		totals += v
	}
	d.answer.part1 = totals
	return nil
}

func processLine(line string, constraint Set) (bool, error) {
	// each set is separated by a semicolon
	sets := strings.Split(line, ";")
	possible := true
	// each value is separated by a comma
	for _, set := range sets {
		values := strings.Split(set, ",")
		for _, strValue := range values {
			strValue = strings.TrimSpace(strValue)
			valueKeyPair := strings.Split(strValue, " ")
			value, err := strconv.Atoi(valueKeyPair[0])
			if err != nil {
				return false, err
			}
			switch valueKeyPair[1] {
			case "blue":
				if value > constraint.Blue {
					return false, nil
				}
			case "red":
				if value > constraint.Red {
					return false, nil
				}
			case "green":
				if value > constraint.Green {
					return false, nil
				}
			}
		}
	}
	return possible, nil
}

func NewChallenge(dataDir string) *day {
	return &day{
		dataDir: dataDir,
	}
}

func Run() *cobra.Command {

	d := NewChallenge("day02/data.txt")

	cmd := &cobra.Command{
		Use: "2",
		Run: func(cmd *cobra.Command, args []string) {
			err := d.Solve(Set{
				Red:   12,
				Green: 13,
				Blue:  14,
			})
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Answer:", d.answer)
		},
	}
	return cmd
}

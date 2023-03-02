package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type day struct {
	// store any necessary configs here
	dir string
}

// run the day solution independently
func main() {
	d, err := prepare("data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = d.solve()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func prepare(dir string) (*day, error) {
	// read data from disk
	return &day{dir}, nil
}

func (d *day) solve() error {
	// read data from disk
	file, err := os.Open(d.dir)
	if err != nil {
		return err
	}

	defer file.Close()

	highestElf := 0
	currentElf := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// found empty line, stop reading

			if currentElf > highestElf {
				highestElf = currentElf
			}
			// reset currentElf to 0
			currentElf = 0
			continue
		}
		currenLine, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			return err
		}
		currentElf += currenLine
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Highest elf:", highestElf)
	return nil
}

func Runner() *cobra.Command {
	var d *day
	cmd := &cobra.Command{
		Use:   "1",
		Short: "Day 1 solution",
		
		PreRun: func(cmd *cobra.Command, args []string) {
			err := error(nil)
			d, err = prepare("day01/data.txt")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},

		Run: func(cmd *cobra.Command, args []string) {
			err := d.solve()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	return cmd
}

package day02

import (
	"fmt"
	"os"

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

func (d *day) Solve() error {

	return nil
}

func NewChallenge(dataDir string) *day {
	return &day{
		dataDir: dataDir,
	}
}

func Run() *cobra.Command {

	d := NewChallenge("day02/data.txt")

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

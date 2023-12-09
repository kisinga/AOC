package day01

import (
	"fmt"
	"os"

	"github.com/kisinga/AOC/2023/types"
	"github.com/spf13/cobra"
)

type day struct {
	// store any necessary configs here
	dataDir string
	file    *os.File
}

func (d *day) LoadData() error {
	file, err := os.Open(d.dataDir)
	if err != nil {
		return err
	}
	d.file = file
	return nil
}

func (d *day) Solve() error {
	d.LoadData()
	return nil

}

func NewDay(dataDir string) types.Day {
	return &day{
		dataDir: dataDir,
	}
}

func Run() *cobra.Command {

	d := NewDay("data.txt")

	cmd := &cobra.Command{
		Use: "1",
		Run: func(cmd *cobra.Command, args []string) {
			err := d.Solve()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

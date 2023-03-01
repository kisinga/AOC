package day02

import (
	"fmt"

	"github.com/spf13/cobra"
)

type day struct {
	// data
}

func prepare() *day {
	// read data from disk
	return &day{}
}

func (*day) solve() {
	// solve the problem
	fmt.Println("Hello from mycommand!")
}

func Runner() *cobra.Command {
	var d *day
	cmd := &cobra.Command{
		Use:   "2",
		Short: "Day 2 solution",
		PreRun: func(cmd *cobra.Command, args []string) {
			d = prepare()
		},

		Run: func(cmd *cobra.Command, args []string) {
			d.solve()
		},
	}

	return cmd
}

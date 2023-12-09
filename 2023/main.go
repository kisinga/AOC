package main

import (
	"fmt"
	"os"

	"github.com/kisinga/AOC/2023/day01"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Short: "My solutions to the Advent of Code challenges",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Please provide a day number to run!")
		},
	}

	rootCmd.AddCommand(day01.Run())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package playground

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Play struct {
}

func findSubsets(arr []int) {
	n := len(arr)
	fmt.Println(1 & 2)
	for i := 0; i < (1 << n); i++ {
		fmt.Print("{ ")
		for j := 0; j < n; j++ {
			x := 1 << j
			y := i & x
			fmt.Sprintf("x: %d, y: %d\n", x, y)
			if (i & (1 << j)) > 0 {
				fmt.Printf("%d ", arr[j])
			}
		}
		fmt.Println("}")
	}
}

func (p *Play) play() error {
	array := []int{7, 12, 31, 99}
	findSubsets(array)
	return nil
}

func NewPlay(dataDir string) *Play {
	return &Play{}
}

func Run() *cobra.Command {

	d := NewPlay("day02/data.txt")

	cmd := &cobra.Command{
		Use: "p",
		Run: func(cmd *cobra.Command, args []string) {
			err := d.play()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

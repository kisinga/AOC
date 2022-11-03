package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// describe values of the hourglass
// top left
// top left + 1 (mid)
// top left + 2
// shift mid
// double shift left (bottom)
// bottom + 1
// bottom + 2

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != 6 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)

	}
	highestTotal := int32(math.MinInt32)

	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr[0])-2; j++ {
			t := hourglassTotal(i, j, arr)
			if highestTotal < t {
				highestTotal = t
			}
		}
	}
	fmt.Println(highestTotal)
}
func hourglassTotal(x, y int, arr [][]int32) int32 {
	diagonal1 := arr[x][y] + arr[x+1][y+1] + arr[x+2][y+2]
	diagonal2 := arr[x][y+2] + arr[x+2][y]
	vertical := arr[x][y+1] + arr[x+2][y+1]
	xx := diagonal1 + diagonal2 + vertical
	return xx
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

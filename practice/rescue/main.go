package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sizeTemp := scanner.Text()
	size, _ := strconv.Atoi(sizeTemp)
	scanner.Scan()
	valuesTemp := strings.Split(scanner.Text(), " ")
	mx, _ := strconv.Atoi(valuesTemp[0])
	my, _ := strconv.Atoi(valuesTemp[1])

	px := 0
	py := 0
	lineTemp := ""
	// loop:=
	for i := 0; i < size; i++ {
		scanner.Scan()
		lineTemp = scanner.Text()
		for pos, val := range lineTemp {
			if string(val) == "p" {
				px, py = i, pos
			}
		}
	}
	ydist := ydistance(py, my)
	xdist := xdistance(px, mx)
	loops := float64(0)

	loops = float64(xdist)

	if ydist < 0 {
		loops = math.Abs(float64(xdist))
		for i := 0.0; i < loops; i++ {
			fmt.Println("UP")
		}
	} else {
		for i := 0.0; i < loops; i++ {
			fmt.Println("DOWN")
		}
	}

	loops = float64(ydist)

	if xdist < 0 {
		loops = math.Abs(float64(ydist))
		for i := 0.0; i < loops; i++ {
			fmt.Println("LEFT")
		}
	} else {
		for i := 0.0; i < loops; i++ {
			fmt.Println("RIGHT")
		}
	}
}

func xdistance(pyPos, myPos int) int {
	val := myPos - pyPos
	if val < 0 {
		return val + 1
	} else if val > 0 {
		return val - 1
	} else {
		return 0
	}
}

func ydistance(pxPos, mxPos int) int {
	return mxPos - pxPos
}

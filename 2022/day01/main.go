package day01

import "fmt"

type day01 struct {
}

func NewDay01() day01 {
	return day01{}
}

func (day01) Run() {
	fmt.Println("Testing day01")
}

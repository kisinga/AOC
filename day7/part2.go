package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bag struct {
	color    string
	children []child
}

type child struct {
	count int
	color string
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	// make a map so that it's easy to search, just in case
	bags := make(map[string]bag)
	for scanner.Scan() {
		lineBag := getBag(scanner.Text())
		bags[lineBag.color] = *lineBag
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	count := countChildren(bags, "shiny gold")

	fmt.Println(count)
}

func countChildren(bags map[string]bag, color string) int {
	count := 0
	for _, v := range bags[color].children {
		count += v.count
		for _, child := range bags[v.color].children {
			count += v.count * child.count
			count += v.count * countChildren(bags, child.color) * child.count
		}
	}
	return count
}

func getBag(line string) *bag {
	split := strings.Split(line, "contain")
	parent, children := strings.ReplaceAll(split[0], "bags", ""), getChildren(split[1])
	return &bag{
		color:    strings.TrimSpace(parent),
		children: children,
	}
}

func getChildren(childrenString string) []child {
	childrenArray := strings.Split(childrenString, ",")
	children := []child{}
	for _, childString := range childrenArray {
		temp := strings.TrimSpace(strings.Trim(strings.Trim(childString, "bags."), "bag."))
		words := strings.Split(temp, " ")
		count, err := strconv.Atoi(words[0])
		if err != nil {
			continue
		}
		children = append(children, child{
			count: count,
			color: strings.Join(words[1:], " "),
		})
	}
	return children
}

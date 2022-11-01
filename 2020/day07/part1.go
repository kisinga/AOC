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
	bags := make(map[string]*bag)
	for scanner.Scan() {
		lineBag := getBag(scanner.Text())
		bags[lineBag.color] = lineBag
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	directRoutes := parents(bags, "shiny gold")
	indirectRoutes := map[string]bag{}
	for key := range directRoutes {
		for k, v := range indirectParents(bags, key) {
			indirectRoutes[k] = v
		}
	}
	merged := indirectRoutes
	for k, v := range directRoutes {
		merged[k] = v
	}
	fmt.Println(len(merged))
}

func indirectParents(bags map[string]*bag, color string) map[string]bag {
	indirectRoutes := map[string]bag{}
	p := parents(bags, color)
	for _, parent := range p {
		indirectRoutes[parent.color] = parent
		for k, v := range indirectParents(bags, parent.color) {
			indirectRoutes[k] = v
		}
	}

	return indirectRoutes
}

func parents(bags map[string]*bag, color string) map[string]bag {
	found := map[string]bag{}
	for _, bag := range bags {
		for _, child := range bag.children {
			if child.color == color {
				found[bag.color] = *bag
			}
		}
	}
	return found
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

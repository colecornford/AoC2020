package main

// Attempt 2 was created using the assistance of Reddit.com contributors to AoC. I was frustrated
// and didn't want to spend an enormous amount of time on this puzzle, but found some help
// I was on the right pathway with my parsing, and it was possible to do it with a tree, but
// overall the better solution was hashmap > hashmap[int] and use that to follow between 
// each parent and their child bags.

// I have included Day7.go here to remind me of where i was going initially.

import (
	"io/ioutil"
	s "strings"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	part1()
	part2()
}

func getData()(lines []string){
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return s.Split(string(content), "\n")
}

func part1() {

	data := getData()
	parent2Children := map[string]map[string]int{} // Maps Parent to [Children]NumberOfBags

	for _, line := range data {
		parent, children := parse(line)
		parent2Children[parent] = children
	}

	fmt.Println(len(uniqueParents(parent2Children, "shiny gold")))
	fmt.Println(countChildren(parent2Children, "shiny gold"))
}

func part2() {
	// data := getData()
}

func parse(line string) (parent string, children map[string]int) {
	
	reParent := regexp.MustCompile(`\w+ \w+`)
	reChildren := regexp.MustCompile(`\d+\s(\w+\s\w+)`)
	reNoChildren := regexp.MustCompile(`no other bags`)

	parent = string(reParent.Find([]byte(line)))
	children = map[string]int{}
	
	if reNoChildren.Find([]byte(line)) != nil {
		return // No Kids
	}

	allChildren := reChildren.FindAllString(line, -1)
	fmt.Println(allChildren)
	for _, child := range allChildren {
		vals := s.Split(string(child), " ")
		children[vals[1] + " " + vals[2]], _ = strconv.Atoi(vals[0])
	}

	return
}

func uniqueParents(parent2Children map[string]map[string]int, child string) map[string]bool {
	parents := map[string]bool{}
	for parent, children := range parent2Children {
		for c := range children {
			if c == child {
				parents[parent] = true
				for parentsParent := range uniqueParents(parent2Children, parent) {
					parents[parentsParent] = true
				}
			}
		}
	}
	return parents
}

func countChildren(parent2Children map[string]map[string]int, parent string) int {
	children := parent2Children[parent]
	total := 0
	for child, count := range children {
		total += count * (1 + countChildren(parent2Children, child))
	}
	return total
}
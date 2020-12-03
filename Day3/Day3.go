package main

import (
	"io/ioutil"
	"strings"
	//"strconv"
	"fmt"
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
	return strings.Split(string(content), "\n")
}

func part1() {
	data := getData()
	fmt.Println(countTrees(data, 3, false))
	//fmt.Print(data)
}

func part2() {
	data := getData()
	r1d1 := countTrees(data, 1, false)
	r3d1 := countTrees(data, 3, false)
	r5d1 := countTrees(data, 5, false)
	r7d1 := countTrees(data, 7, false)
	r1d2 := countTrees(data, 1, true)
	fmt.Print(r1d1 * r3d1 * r5d1 * r7d1 * r1d2)
}

func changePos(currentPos int, slopeLength int) int {
	newPos := currentPos + slopeLength
	if newPos > 30 {
		return newPos % 31
	}
	return newPos
}

func countTrees(lines []string, slopeLength int, skipTWOlol bool) int {
	currentPos := 0
	treeCount := 0
	for x, line := range lines {
		if skipTWOlol {
			if x % 2 != 0 {
				continue
			}
		}
		if string(line[currentPos]) == "#" {
			treeCount = treeCount + 1
		}
		currentPos = changePos(currentPos, slopeLength)
	}
	return treeCount
}
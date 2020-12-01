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
	fmt.Print(data)
}

func part2() {
	// data := getData()
}
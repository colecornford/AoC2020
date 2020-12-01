package main

import (
	"io/ioutil"
	"strings"
	"strconv"
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
	//fmt.Print(data)
	part1:
	for _, num1 := range data {
        for _, num2 := range data {
			x, err := strconv.Atoi(num1)
			if err != nil {

			}
			y, err := strconv.Atoi(num2)
			if err != nil {
				
			}
			if x + y == 2020 {
				product := x * y
				fmt.Printf("%d * %d = %d\n", x, y, product)
				break part1
			}
		}
    }

}

func part2() {
	data := getData()
	//fmt.Print(data)
	part2:
	for _, num1 := range data {
        for _, num2 := range data {
			for _, num3 := range data {
				x, err := strconv.Atoi(num1)
				if err != nil {

				}
				y, err := strconv.Atoi(num2)
				if err != nil {
					
				}
				z, err := strconv.Atoi(num3)
				if err != nil {
					
				}
				if x + y + z == 2020 {
					product := x * y * z
					fmt.Printf("%d * %d * %d = %d\n", x, y, z, product)
					break part2
				}
			}
		}
    }
}
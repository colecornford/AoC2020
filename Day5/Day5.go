package main

import (
	"io/ioutil"
	s "strings"
	"fmt"
	"sort"
)

var seatRow int
var seatCol int

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
	rows := make([]int, 0)
	cols := make([]int, 0)
	highest := 0
	for i := 0; i < 128; i++ {
		rows = append(rows, i)
	}
	for i := 0; i < 8; i++ {
		cols = append(cols, i)
	}
	for _, input := range data {
		parseRows(rows, input[:7])
		parseCols(cols, input[7:])
		seatId := seatRow * 8 + seatCol
		if seatId > highest {
			highest = seatId
		}
		fmt.Printf("Row: %d, Col %d, seatID: %d \n", seatRow, seatCol, seatId)
	}
	fmt.Printf("\nHighest is: %d \n", highest)

}

func part2() {
	data := getData()
	rows := make([]int, 0)
	cols := make([]int, 0)
	highest := 0
	allSeats := make([]int, 0)
	for i := 0; i < 128; i++ {
		rows = append(rows, i)
	}
	for i := 0; i < 8; i++ {
		cols = append(cols, i)
	}
	for _, input := range data {
		parseRows(rows, input[:7])
		parseCols(cols, input[7:])
		seatId := seatRow * 8 + seatCol
		allSeats = append(allSeats, seatId)
		if seatId > highest {
			highest = seatId
		}
	}




	 
	sort.Ints(allSeats)
	for x := 1; x < highest - 1; x++ {
		if (allSeats[x] + allSeats[x+1]) % 2 == 0 {
			fmt.Printf("Your Seat: %d", allSeats[x] + 1)
		}
	}
}



func parseRows(rows []int, input string){
	if len(rows) != 1 {
		if string(input[0]) == "F" {
			rows = rows[:len(rows) / 2]
		} else if string(input[0]) == "B" {
			rows = rows[len(rows) / 2:]
		}
		input = input[1:]
		parseRows(rows, input)
	} else {
		seatRow = rows[0]
	}
}

func parseCols(cols []int, input string){
	if len(cols) != 1 {
		if string(input[0]) == "L" {
			cols = cols[:len(cols) / 2]
		} else if string(input[0]) == "R" {
			cols = cols[len(cols) / 2:]
		}
		input = input[1:]
		parseCols(cols, input)
	} else {
		seatCol = cols[0]
	}
}

/* func keepHalf(ForB string) int {
	if ForB != "F" || ForB != "B"
		return


}
 */

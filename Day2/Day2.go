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
	parsedWords := parseAllPasswords(data)
	count := 0
	for _, p := range parsedWords {
		count = count + validatePasswordpart1(p)
	}
	fmt.Print(count)
}

func part2() {
	data := getData()
	parsedWords := parseAllPasswords(data)
	count := 0
	for _, p := range parsedWords {
		count = count + validatePasswordpart2(p)
	}
	fmt.Print(count)
}

func parseAllPasswords(lines []string) []corruptPassword {

	var parsedWords []corruptPassword
	for _, line := range lines {
        parsedWords = append(parsedWords, parseCorruptPassword(line))
	}
	return parsedWords
}

func parseCorruptPassword(input string) corruptPassword {
	tempMin := strings.Split(input,"-")[0]
	temp := strings.Split(input,"-")[1]

	tempMax := strings.Split(temp," ")[0]
	temp = strings.Split(temp," ")[1] + strings.Split(temp," ")[2]

	req := strings.Split(temp,":")[0]
	temp = strings.Split(temp,":")[1]

	pwd := strings.Split(temp," ")[0]

	// Messy Conversions
	min, err := strconv.Atoi(tempMin)
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(tempMax)
	if err != nil {
		panic(err)
	}

	lineItem := corruptPassword{password: pwd, minimumChar: min, maximumChar: max, requiredChar: req}
	return lineItem
}

func validatePasswordpart1(p corruptPassword) int {

	if (strings.Count(p.password, p.requiredChar) >= p.minimumChar) && (strings.Count(p.password, p.requiredChar) <= p.maximumChar) {
		return 1
	}
	return 0
}

func validatePasswordpart2(p corruptPassword) int {

	//fmt.Print(p.password + " " + string(p.password[p.minimumChar]) + "\n")
	if string(p.password[p.minimumChar-1]) == p.requiredChar && string(p.password[p.maximumChar-1]) != p.requiredChar {
		return 1
	}
	if string(p.password[p.minimumChar-1]) != p.requiredChar && string(p.password[p.maximumChar-1]) == p.requiredChar {
		return 1
	}
	return 0
}

type corruptPassword struct {
    password string
	minimumChar  int
	maximumChar  int
	requiredChar  string
}
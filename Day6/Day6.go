package main

import (
	"io/ioutil"
	s "strings"
	"fmt"
	"regexp"
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
	return s.Split(string(content), "\n\n")
}

func part1() {
	data := getData()
	total := 0
	var decs map[string]bool
	decs = make(map[string]bool)
	for _, group := range data {
		decs = make(map[string]bool)
		for _, declaration := range group {
			notWhitespace, _ := regexp.MatchString(`[a-z]`, string(declaration))
			if notWhitespace {
				decs[string(declaration)] = true
			}
		}
		total = total + len(decs)
	}
	fmt.Println(total)    
}

func part2() {
	data := getData()

	total := 0
	for _, group := range data {
		decs := make(map[string]string)
		for _, declaration := range group {
			notWhitespace, _ := regexp.MatchString(`[a-z]`, string(declaration))
			if notWhitespace {
				decs[string(declaration)] = string(declaration)
			}
		}

		users := s.Split(string(group), "\n")
		//fmt.Print(users)
		//fmt.Print(decs)    

		for _, dec := range decs {
			total = total + checkContains(users, dec)
		}
	}
	fmt.Println(total)    
}

func checkContains(users []string, dec string) int {
	for _, user := range users {
		if !s.Contains(user, dec) {
			return 0
		}
	}
	return 1
}
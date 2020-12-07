package main

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
	count := 0
	// fmt.Print(data)
	bagRules := make([]bag, 0)
	for _, line := range data {
		bagRules = append(bagRules, storeBagRules(line))
	}
	shinyGoldBag := bag{adjective:"shiny",colour:"gold"}

	for _, rule := range bagRules {
		//fmt.Println(rule)
		count = count + validateCanContain(rule, shinyGoldBag, 0)
	}
	
	fmt.Println("hello" + string(count))
}

func part2() {
	// data := getData()
}



func validateCanContain(rule bag, inner bag, total int) int {
	if inner.contains == nil {
		return total
	} else {
		for _, innerBag := range inner.contains {
			if innerBag.adjective == inner.adjective && innerBag.colour == inner.colour {
				total = total + 1
			} else {
				newBag := retrieveBagRule(rule)
				total = total + validateCanContain(newBag, inner, total)
			}
		}
	}
	return total
}

func retrieveBagRule(ruleBag bag) bag {
	data := getData()
	bagRules := make([]bag, 0)
	for _, line := range data {
		bagRules = append(bagRules, storeBagRules(line))
	}
	
	for _, bag := range bagRules {
		if bag.adjective == ruleBag.adjective && bag.colour == ruleBag.colour {
			fmt.Println(bag)
			return bag
		}
	}
	return bag{}
} 

// allRules = append(allRules[:i], allRules[i+1:]...)

type bag struct {
	adjective string
	colour string
	contains []bag
}

func storeBagRules(bagString string) bag{
	//fmt.Printf(bagString)

	regexOuter := regexp.MustCompile(`\w+ \w+`)
	regexInner := regexp.MustCompile(`\d+ \w+ \w+`)

	outerBagString := s.Split(string(regexOuter.Find([]byte(bagString))), " ")
	outerBag := bag{adjective: outerBagString[0], colour: outerBagString[1]}

	innerBagStrings := regexInner.FindAllString(bagString, -1)
	for _, innerBag := range innerBagStrings {
		newBagString := s.Split(string(regexInner.Find([]byte(innerBag))), " ")
		
		newBag := bag{adjective: newBagString[1], colour: newBagString[2]}
		for x, _ := strconv.Atoi(newBagString[0]); x > 0; x-- {
			outerBag.contains = append(outerBag.contains, newBag)
		}
	}
	fmt.Println(outerBag)
	return outerBag
}
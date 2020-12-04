package main

import (
	"io/ioutil"
	s "strings"
	"strconv"
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
	for _, passport := range data {
		if parsePassport1(passport) {
			total = total + 1
		}
	}
	fmt.Print(total)    
}

func part2() {
	data := getData()
	total := 0
	for _, passport := range data {
		if parsePassport2(passport) {
			total = total + 1
		}
	}
	fmt.Print(total)    
}

func parsePassport1(input string) bool {
	req := make([]string, 0)
	req = append(req, "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid")
	for _, val := range req {
		if ! s.Contains(input, val) {
			return false
		}
	}
	return true
}

func parsePassport2(input string) bool {
	input = s.ReplaceAll(input, "\n", " ") 
	vals := s.Split(input, " ") 
	//fmt.Print(vals)
	if !parsePassport1(input) {
		fmt.Print("notEnoughFields")
		fmt.Print(vals)
		return false
	}
	for _, val := range vals {
		if !validate(val) {
			return false
		}
	}
	return true
}

func validate(input string) bool {
	if ! s.Contains(input, ":"){
		fmt.Println("EOF")
		return true
	}
	prefix := s.Split(input, ":")[0]
	input = s.Split(input, ":")[1]
	//fmt.Println(prefix)
	//fmt.Println(input)

	if prefix == "cid" {
		return true
	}

	if prefix == "hgt" {
		hgtCm, _ := regexp.MatchString(`[0-9]{3}(cm)`, input)
		hgtInch, _ := regexp.MatchString(`[0-9]{2}(in)`, input)
		if hgtCm {
			x := s.Split(input, "cm") 
			hgt, _ := strconv.Atoi(x[0])
			if ! (hgt >= 150 && hgt <= 193 ) {
				fmt.Println(hgt)
				fmt.Println("Wrong CM")
				return false
			}
		} else if hgtInch {
			x := s.Split(input, "in")
			hgt, _ := strconv.Atoi(x[0])
			if !(hgt >= 59 && hgt <= 76 ) {
				fmt.Println(hgt)
				fmt.Println("Wrong INCH")
				return false
			}

		} else if !hgtCm || !hgtInch {
			fmt.Println(hgt)
			fmt.Println("Wrong HGT")
			return false
		}
	}

	if prefix == "hcl" {
		validHcl, _ := regexp.MatchString(`#[a-f0-9]{6}`, input)
		if !validHcl {
			fmt.Println(input)
			fmt.Println("Wrong HCL")
			return false
		}
	}

	if prefix == "ecl" {
		validEcl, _ := regexp.MatchString(`amb|blu|brn|gry|grn|hzl|oth`, input)
		if !validEcl {
			fmt.Println(input)
			fmt.Println("Wrong ECL")
			return false
		}

	}

	if prefix == "pid" {
		validPid, _ := regexp.MatchString(`[0-9]{9}`, input)
		if !validPid {
			fmt.Println(input)
			fmt.Println("Wrong PID")
			return false
		}
	}

	if prefix == "byr" {
		if ! validYear(input, 1920, 2002) {
			fmt.Println(input)
			fmt.Println("Wrong BYR")
			return false
		}
	}

	if prefix == "iyr" {
		if !validYear(input, 2010, 2020) {
			fmt.Println(input)
			fmt.Println("Wrong IYR")
			return false
		}
	}

	if prefix == "eyr" {
		if !validYear(input, 2020, 2030) {
			fmt.Println(input)
			fmt.Println("Wrong EYR")
			return false
		}
	}
	return true
}


func validYear(input string, min int, max int) bool {
	yr, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	if !(yr >= min && yr <= max) {
		return false
	}
	return true
}
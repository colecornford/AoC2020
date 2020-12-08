package main

import (
	"io/ioutil"
	"strings"
	"fmt"
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
	return strings.Split(string(content), "\n")
}

func part1() {
	instructionSet := getData()
	linesRan := make([]int,0)
	infLoop := false
	currentInstruction := instructionSet[0]
	line := 0
	acc := 0
	for ! infLoop {
		linesRan = append(linesRan, line)
		acc, line = runInstruction(currentInstruction, line, acc)
		currentInstruction = instructionSet[line - 1]
		for _, l := range linesRan {
			if l == line {
				fmt.Printf("Final Accumulation: %d", acc)
				infLoop = true
			}
		}
	} 

}

func part2() {
	is := getData()
	for x := 0; x < len(is); x++ {
		instructionSet := getData()
		instruction := instructionSet[x]
		jmpORnop := strings.Split(string(instruction), " ")

		if jmpORnop[0] == "jmp" {
			instruction = strings.ReplaceAll(instruction, "jmp", "nop")
		} else if jmpORnop[0] == "nop" {
			instruction = strings.ReplaceAll(instruction, "nop", "jmp")
		}
		linesRan := make([]int,0)
		infLoop := false
		instructionSet[x] = instruction
		currentInstruction := instructionSet[0]
		line := 0
		acc := 0
		for ! infLoop {
			linesRan = append(linesRan, line)
			acc, line = runInstruction(currentInstruction, line, acc)
			if strings.Contains(instruction, "jmp +0"){
				break
			}
			currentInstruction = instructionSet[line - 1]
			for _, l := range linesRan {
				if l == line {
					// fmt.Printf("Final Accumulation: %d", acc)
					infLoop = true
				}
			}
			if line == len(instructionSet) {
				fmt.Printf("Final Accumulation: %d", acc)
				return
			}
		} 
	}
}

func runInstruction(instruction string, line int, acc int) (newAcc int, newLine int) {
	instruction = strings.ReplaceAll(instruction, "+", "+ ")
	instruction = strings.ReplaceAll(instruction, "-", "- ")
	x := strings.Split(string(instruction), " ")

	op := x[0]
	operand := x[1]
	amt, _ := strconv.Atoi(x[2])
	
	if op == "nop" {
		newAcc = acc
		newLine = line + 1
	} else if op == "acc" {
		if operand == "+" {
			newAcc = acc + amt
		} else if operand == "-" {
			newAcc = acc - amt
		}
		newLine = line + 1
	} else if op == "jmp" {
		if operand == "+" {
			newLine = line + amt
		} else if operand == "-" {
			newLine = line - amt
		}
		newAcc = acc
	}
	return
}
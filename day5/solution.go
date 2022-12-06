package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("input.txt")
	//file, err := os.ReadFile("input.test.txt")

	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	fileContent := string(file)

	scanner := bufio.NewScanner(strings.NewReader(fileContent))

	cratesStack1 := [][]string{[]string{}, []string{}, []string{}, []string{}, []string{}, []string{}, []string{}, []string{}, []string{}}
	moveList := [][]string{}

	for scanner.Scan() {
		line := append([]rune(scanner.Text()), ' ')
		if line[1] == '1' {
			break
		}

		for i, j := 0, 0; i < len(line); i, j = i+3+1, j+1 {
			if line[i+1] == ' ' {
				continue
			}
			cratesStack1[j] = append(cratesStack1[j], string(line[i+1]))
		}
	}

	for i, crateStack := range cratesStack1 {
		stack := []string{}
		for j := len(crateStack) - 1; j >= 0; j-- {
			stack = append(stack, crateStack[j])
		}
		cratesStack1[i] = stack
	}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		re := regexp.MustCompile("[0-9]+")
		moveList = append(moveList, re.FindAllString(line, -1))
	}

	sToI := func(i string) int {
		result, _ := strconv.Atoi(i)
		return result
	}

	getInstructions := func(moveSet []string) (int, int, int) {
		return sToI(moveSet[0]), sToI(moveSet[1]) - 1, sToI(moveSet[2]) - 1
	}

	cratesStack2 := [][]string{}
	for i, crateStack := range cratesStack1 {
		cratesStack2 = append(cratesStack2, []string{})
		for _, crate := range crateStack {
			cratesStack2[i] = append(cratesStack2[i], crate)
		}
	}

	println("Solution 1: " + getAnswer(moveCrates(cratesStack1, moveList,
		func(cratesStack [][]string, moveSet []string) [][]string {
			number, start, end := getInstructions(moveSet)
			for i := 0; i < number; i++ {
				if len(cratesStack[start]) == 0 {
					break
				}
				startLength := len(cratesStack[start])
				if startLength == 0 {
					continue
				}
				crate := cratesStack[start][startLength-1]
				if startLength-1 < 0 {
					cratesStack[start] = []string{}
				} else {
					cratesStack[start] = cratesStack[start][:(startLength - 1)]
				}
				cratesStack[end] = append(cratesStack[end], crate)
			}
			return cratesStack
		})))
	println("Solution 2: " + getAnswer(moveCrates(cratesStack2, moveList,
		func(cratesStack [][]string, moveSet []string) [][]string {
			number, start, end := getInstructions(moveSet)
			//fmt.Printf("\n=========\ninstructions: %v", []int{number, start, end})
			//fmt.Printf("\nstart: %v: %v\n", start, cratesStack[start])
			//fmt.Printf("end: %v: %v\n", end, cratesStack[end])
			crates := cratesStack[start][len(cratesStack[start])-number : len(cratesStack[start])]
			//fmt.Printf("crates: %v\n", crates)
			cratesStack[start] = cratesStack[start][:len(cratesStack[start])-number]
			cratesStack[end] = append(cratesStack[end], crates...)
			//fmt.Printf("cratesStack After\n")
			//fmt.Printf("start: %v: %v\n", start, cratesStack[start])
			//fmt.Printf("end: %v: %v\n", end, cratesStack[end])
			//println("--------")

			return cratesStack
		})))
}

func getAnswer(cratesStack [][]string) string {
	answer := ""
	for _, crateStack := range cratesStack {
		if len(crateStack) == 0 {
			answer += " "
			continue
		}

		answer += string(crateStack[len(crateStack)-1])
	}
	return answer
}

func moveCrates(cratesStack [][]string, moveList [][]string, callback func([][]string, []string) [][]string) [][]string {

	for _, moveSet := range moveList {
		cratesStack = callback(cratesStack, moveSet)
	}
	return cratesStack
}

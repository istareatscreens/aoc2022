package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Section struct {
	lower int
	upper int
}

func main() {

	file, err := os.ReadFile("input.txt")
	//file, err := os.ReadFile("input.test.txt")

	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	fileContent := string(file)

	scanner := bufio.NewScanner(strings.NewReader(fileContent))

	count := 0
	count2 := 0

	createSection := func(row string) Section {
		upperAndLowerBound := strings.Split(row, "-")
		lower, _ := strconv.Atoi(upperAndLowerBound[0])
		upper, _ := strconv.Atoi(upperAndLowerBound[1])
		return Section{lower: lower, upper: upper}
	}

	compareSections := func(sections []Section) int {
		if (sections[0].lower >= sections[1].lower && sections[0].upper <= sections[1].upper) ||
			(sections[1].lower >= sections[0].lower && sections[1].upper <= sections[0].upper) {
			return 1
		}
		return 0
	}

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		sections := []Section{createSection(row[0]), createSection(row[1])}
		count += compareSections(sections)

		if (compareSections(sections) == 1) ||
			(sections[0].lower >= sections[1].lower && sections[0].lower <= sections[1].upper) ||
			(sections[1].lower >= sections[0].lower && sections[1].lower <= sections[0].upper) {
			count2 += 1
		}

	}

	println("Solution 1: " + strconv.Itoa(count))
	println("Solution 2: " + strconv.Itoa(count2))

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	fileContent := string(file)

	scanner := bufio.NewScanner(strings.NewReader(fileContent))
	sum := 0
	var sums []int = []int{}
	for scanner.Scan() {
		result := scanner.Text()
		if result == "" {
			sums = append(sums, sum)
			sum = 0
			continue
		}
		value, _ := strconv.Atoi(result)
		sum += value
	}
	//fmt.Printf("%v", sums)

	sort.Ints(sums)

	println("Solution 1: " + strconv.Itoa(sums[len(sums)-1]))
	println("Solution 2: " + strconv.Itoa(sums[len(sums)-1]+sums[len(sums)-2]+sums[len(sums)-3]))

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	orderedmap "github.com/wk8/go-ordered-map"
)

type Item struct {
	count int
}

type Rucksack struct {
	compartment1 orderedmap.OrderedMap
	compartment2 orderedmap.OrderedMap
}

func main() {

	file, err := os.ReadFile("input.txt")
	//file, err := os.ReadFile("input.test.txt")

	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	fileContent := string(file)

	scanner := bufio.NewScanner(strings.NewReader(fileContent))

	convertItemToPriority := func(char rune) int {
		if unicode.IsUpper(char) {
			return int(char-64) + 26
		}
		return int(char - 96)
	}

	priorityTotal := 0

	createCompartmentInvantory := func(rucksack []rune) orderedmap.OrderedMap {
		compartmentInvantory := orderedmap.New()
		for _, item := range rucksack {
			item := convertItemToPriority(item)
			count, present := compartmentInvantory.Get(item)
			if !present {
				compartmentInvantory.Set(item, &Item{1})
				continue
			}
			compartmentInvantory.Set(item, &Item{count.(*Item).count + 1})
		}
		return *compartmentInvantory
	}

	rucksacks := []Rucksack{}
	for scanner.Scan() {
		rucksack := []rune(scanner.Text())
		rucksacks = append(rucksacks, (Rucksack{
			compartment1: createCompartmentInvantory(rucksack[:len(rucksack)/2]),
			compartment2: createCompartmentInvantory(rucksack[len(rucksack)/2:]),
		}))
	}

	for _, rucksack := range rucksacks {
		for pair := rucksack.compartment1.Oldest(); pair != nil; pair = pair.Next() {
			_, present := rucksack.compartment2.Get(pair.Key)
			if present {
				priorityTotal += pair.Key.(int)
			}
		}
	}

	println("Solution 1: " + strconv.Itoa(priorityTotal))

	scanner2 := bufio.NewScanner(strings.NewReader(fileContent))
	var rucksacks2 [][]rune = [][]rune{}
	for scanner2.Scan() {
		rucksacks2 = append(rucksacks2, []rune(scanner2.Text()))
	}

	priorityTotal2 := 0
	for i := 0; i < len(rucksacks2)-2; i += 3 {
	out:
		for _, item1 := range rucksacks2[i] {
			for _, item2 := range rucksacks2[i+1] {
				if item1 == item2 {
					for _, item3 := range rucksacks2[i+2] {
						if item3 == item2 {
							priorityTotal2 += convertItemToPriority(item1)
							break out
						}
					}
				}
			}
		}
	}

	println("Solution 2: " + strconv.Itoa(priorityTotal2))

}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type directory struct {
	name     string
	files    map[string]int
	sub      map[string]*directory
	previous *directory
}

func main() {

	file, err := os.ReadFile("input.txt")
	if len(os.Args) > 1 {
		file, err = os.ReadFile("input.test.txt")
	}

	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	fileContent := string(file)

	scanner := bufio.NewScanner(strings.NewReader(fileContent))

	scanner.Scan()
	scanner.Text()
	root := &directory{name: "/", files: make(map[string]int), sub: make(map[string]*directory), previous: nil}
	currentDirectory := root

	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '$' {
			command := line[2:4]
			if command == "cd" {
				argument := line[5:]
				if argument == ".." {
					currentDirectory = currentDirectory.previous
					continue
				}
				name := (*currentDirectory).name + argument + "/"
				if nil == (*currentDirectory).sub[name] {
					println("FUCK")
				}
				currentDirectory = (*currentDirectory).sub[name]
				continue
			} else if command == "ls" {
				continue
			}
		}
		if line[0:3] == "dir" {
			name := currentDirectory.name + line[4:] + "/"
			_, ok := currentDirectory.sub[name]
			if ok {
				continue
			}
			currentDirectory.sub[name] = &directory{name: name, files: make(map[string]int), sub: make(map[string]*directory), previous: currentDirectory}
			continue
		}

		fileData := strings.Split(line, " ")
		currentDirectory.files[fileData[1]], err = strconv.Atoi(fileData[0])
	}

	sizeMap := make(map[string]int)
	getSize(*root, &sizeMap)

	sum := 0
	for _, size := range sizeMap {
		if size <= 100000 {
			sum += size
		}
	}

	fmt.Printf("Solution 1: %v\n", strconv.Itoa(sum))

	minSizeToDelete := ^uint64(0)
	for _, size := range sizeMap {
		if size >= 30000000-(70000000-sizeMap["/"]) {
			minSizeToDelete = uint64(math.Min(float64(minSizeToDelete), float64(size)))
		}
	}

	fmt.Printf("Solution 2: %v\n", strconv.Itoa(int(minSizeToDelete)))
}

func getSize(root directory, sizeMap *map[string]int) int {
	for _, subDirectory := range root.sub {
		(*sizeMap)[root.name] += getSize(*subDirectory, sizeMap)
	}

	for _, size := range root.files {
		(*sizeMap)[root.name] += size
	}

	return (*sizeMap)[root.name]
}

func ignoreError(val interface{}, err error) interface{} {
	return val
}

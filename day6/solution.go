package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	packet := []rune(scanner.Text())

	getStart := func(packet []rune, size int) int {
		startOfPacket := 0
		letterMap := map[string]int{}
		for i := 0; i < len(packet); i += 1 {
			count := 0
			for _, value := range letterMap {
				if value == 1 {
					count += 1
				}
			}
			if count == size {
				startOfPacket = i
				break
			}
			if i >= size {
				letterMap[string(packet[i-size])]--
			}
			letterMap[string(packet[i])]++

		}
		return startOfPacket
	}

	println("Solution for part 1 is: ", getStart(packet, 4))
	println("Solution for part 2 is: ", getStart(packet, 14))
}

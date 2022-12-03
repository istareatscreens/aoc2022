package main

import (
	"bufio"
	"fmt"
	"os"
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

	/*
		(1 for Rock, 2 for Paper, and 3 for Scissors)
		(0 if you lost, 3 if the round was a draw, and 6 if you won)
		A for Rock, B for Paper, and C for Scissors
		X for Rock, Y for Paper, and Z for Scissors
	*/

	analyzePlay := func(opponentMove int, winningMove int) func(int) int {
		return func(playerMove int) int {
			if opponentMove == playerMove {
				return 3
			} else if playerMove == winningMove {
				return 6
			} else {
				return 0
			}
		}
	}

	winningMap := make(map[int]func(int) int)
	winningMap[1] = analyzePlay(1, 2)
	winningMap[2] = analyzePlay(2, 3)
	winningMap[3] = analyzePlay(3, 1)
	sum := 0
	for scanner.Scan() {
		move := []rune(scanner.Text())
		opponentsMove := move[0] - 'A' + 1
		playersMove := move[2] - 'X' + 1
		sum += winningMap[int(opponentsMove)](int(playersMove)) + int(playersMove)
	}

	println("Solution 1: " + strconv.Itoa(sum))

	scanner = bufio.NewScanner(strings.NewReader(fileContent))
	losingMoveMap := make(map[int]int)
	winningMoveMap := make(map[int]int)
	losingMoveMap[1] = 3
	losingMoveMap[2] = 1
	losingMoveMap[3] = 2
	winningMoveMap[1] = 2
	winningMoveMap[2] = 3
	winningMoveMap[3] = 1
	sum = 0
	/*
			X means you need to lose,
			Y means you need to end the round in a draw,
		 	Z means you need to win
	*/
	for scanner.Scan() {
		move := []rune(scanner.Text())
		opponentsMove := move[0] - 'A' + 1
		playersMove := move[2]
		switch playersMove {
		case 'X':
			sum += 0 + losingMoveMap[int(opponentsMove)]
		case 'Y':
			sum += 3 + int(opponentsMove)
		case 'Z':
			sum += 6 + winningMoveMap[int(opponentsMove)]
		}
	}

	println("Solution 2: " + strconv.Itoa(sum))

}

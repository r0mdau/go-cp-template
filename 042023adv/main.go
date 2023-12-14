package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func eprint(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func print(a ...interface{}) {
	fmt.Println(a...)
}

func contestResponse() {
	var solution int
	lCount := 1
	mScore := make(map[int]int)
	mCardCount := make(map[int]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		spl1 := strings.Split(line, ": ")
		games := strings.Split(spl1[1], " | ")
		winnings := strings.Split(games[0], " ")
		plays := strings.Split(games[1], " ")

		wMap := map[string]int{}
		for _, v := range winnings {
			wMap[v] = 0
		}
		// deleting empty beginning string \o/
		delete(wMap, "")

		var gameScore, scratchScore int
		for _, v := range plays {
			if _, ok := wMap[v]; ok {
				if gameScore == 0 {
					gameScore++
				} else {
					gameScore *= 2
				}
				scratchScore++
			}
		}
		mScore[lCount] = scratchScore

		cardCount := mCardCount[lCount]
		for z := 0; z <= cardCount; z++ {
			for i := 1; i <= scratchScore; i++ {
				mCardCount[lCount+i]++
			}
		}

		solution += gameScore
		lCount++
	}
	print("Part1:", solution)

	solution = 0
	for _, v := range mCardCount {
		solution += v
	}
	print("Part2:", solution+lCount-1)
}

func main() {
	contestResponse()
}

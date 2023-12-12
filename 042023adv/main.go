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

		var gameScore int
		for _, v := range plays {
			if _, ok := wMap[v]; ok {
				if gameScore == 0 {
					gameScore++
				} else {
					gameScore *= 2
				}
			}
		}
		solution += gameScore
	}
	print(solution)
}

func main() {
	contestResponse()
}

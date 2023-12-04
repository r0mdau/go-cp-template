package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func eprint(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func print(a ...interface{}) {
	fmt.Println(a...)
}

var bag map[string]int

func contestResponse() {
	// Game 1: 1 green, 1 blue, 1 red; 3 green, 1 blue, 1 red; 4 green, 3 blue, 1 red; 4 green, 2 blue, 1 red; 3 blue, 3 green
	var solution, gCount int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		gCount++

		line := scanner.Text()
		spl1 := strings.Split(line, ": ")
		games := strings.Split(spl1[1], "; ")

		goodGame := true

		for _, game := range games {
			bag = map[string]int{
				"red":   12,
				"green": 13,
				"blue":  14,
			}

			colors := strings.Split(game, ", ")
			for _, color := range colors {
				print("Color:", color)
				spl2 := strings.Split(color, " ")
				val, _ := strconv.Atoi(spl2[0])
				print("Value : ", val)
				bag[spl2[1]] -= val
			}

			if bag["red"] < 0 || bag["green"] < 0 || bag["blue"] < 0 {
				goodGame = false
				break
			}
		}
		if goodGame {
			solution += gCount
		}
	}

	print(solution)
}

func main() {
	contestResponse()
}

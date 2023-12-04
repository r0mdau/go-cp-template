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

var minBag map[string]int

func contestResponse() {
	// Game 1: 1 green, 1 blue, 1 red; 3 green, 1 blue, 1 red; 4 green, 3 blue, 1 red; 4 green, 2 blue, 1 red; 3 blue, 3 green
	var solution int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		spl1 := strings.Split(line, ": ")
		games := strings.Split(spl1[1], "; ")
		print("Games:", games)

		minBag = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, game := range games {
			colors := strings.Split(game, ", ")
			for _, color := range colors {
				//print("Color:", color)
				spl2 := strings.Split(color, " ")
				val, _ := strconv.Atoi(spl2[0])
				//print("Value : ", val)

				if val > minBag[spl2[1]] {
					minBag[spl2[1]] = val
				}
			}
		}
		solution += minBag["red"] * minBag["green"] * minBag["blue"]
	}

	print(solution)
}

func main() {
	contestResponse()
}

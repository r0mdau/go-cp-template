package main

import (
	"bufio"
	"fmt"
	"os"
)

func eprint(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func print(a ...interface{}) {
	fmt.Println(a...)
}

// 467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..

func contestResponse() {
	var solution, lCount int
	var engine = make([][]byte, 140)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		engine[lCount] = make([]byte, len(line))
		for c := 0; c < len(line); c++ {
			engine[lCount][c] = line[c]
		}
		lCount++
	}
	print(solution)
}

func main() {
	contestResponse()
}

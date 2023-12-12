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

type Coord struct {
	X, Y int
}

func contestResponse() {
	engine := make(map[Coord]rune)
	gears := make(map[Coord][]int)
	var eWidth, lCount, solution int
	var line string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()

		for x, r := range line {
			engine[Coord{x, lCount}] = r
		}
		lCount++
	}
	eWidth = len(line)

	for y := 0; y <= lCount; y++ {
		for x := 0; x <= eWidth; {
			var numLen, pNumber int
			for {
				c := engine[Coord{x + numLen, y}]
				if c < '0' || c > '9' {
					break
				}
				pNumber = 10*pNumber + int(c-'0')
				numLen++
			}

			if numLen == 0 {
				x++
				continue
			}

			var found bool
			for j := y - 1; j <= y+1; j++ {
				for i := x - 1; i <= x+numLen; i++ {
					coord := Coord{i, j}
					ru := engine[coord]
					if (ru >= '0' && ru <= '9') || ru == 0 || ru == '.' {
						continue
					}
					if !found {
						solution += pNumber
						found = true
					}
					if ru == '*' {
						gears[coord] = append(gears[coord], pNumber)
					}
				}
			}
			x += numLen
		}
	}
	print(solution)

	ratio := 0
	for _, v := range gears {
		if len(v) == 2 {
			ratio += v[0] * v[1]
		}
	}
	print(ratio)
}

func main() {
	contestResponse()
}

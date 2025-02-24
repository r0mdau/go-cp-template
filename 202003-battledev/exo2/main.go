package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func eprint(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func print(a ...interface{}) {
	fmt.Println(a...)
}

func contestResponse() {
	eprint("=============== BEGIN ===============")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // first line is maybe not needed
	n := scanner.Text()
	eprint(n)
	// initialize vars

	var previous, occur, maxOccur int
	occur = 1
	for scanner.Scan() {
		current, _ := strconv.Atoi(scanner.Text())
		if current == previous {
			occur++
		} else {
			occur = 1
		}
		previous = current
		if maxOccur < occur {
			maxOccur = occur
		}

		eprint(current)
		// do something with current
	}

	eprint("")

	print(maxOccur)
}

func main() {
	contestResponse()
}

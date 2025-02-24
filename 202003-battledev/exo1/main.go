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

func contestResponse() {
	eprint("=============== BEGIN ===============")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // first line is maybe not needed

	// initialize vars
	m := make(map[string]int)

	for scanner.Scan() {
		current := scanner.Text()
		eprint(current)
		if _, ok := m[current]; ok {
			m[current]++
		} else {
			m[current] = 1
		}
	}
	eprint(m)
	var solution1, solution2 string
	var count1, count2 int
	for key, val := range m {
		if count1 < val {
			count1 = val
			solution1 = key
		}
	}
	delete(m, solution1)
	for key, val := range m {
		if count2 < val {
			count2 = val
			solution2 = key
		}
	}

	eprint("")
	print(solution1, solution2)
}

func main() {
	contestResponse()
}

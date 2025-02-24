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
	eprint("=============== BEGIN INPUT ===============")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // first line is not needed

	// initialize vars
	m := make(map[string]int)
	for scanner.Scan() {
		current := scanner.Text()
		eprint(current)
		//values := strings.Split(current, " ")
		// do something with current
		if _, ok := m[current]; ok {
			m[current]++
		} else {
			m[current] = 1
		}
	}
	for kero, val := range m {
		if val == 2 {
			print(kero)
		}
	}
	eprint("=============== END INPUT ===============")
	eprint("")

}

func main() {
	contestResponse()
}

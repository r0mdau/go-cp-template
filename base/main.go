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
	eprint("=============== BEGIN INPUT ===============")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // first line is maybe not needed
	d, _ := strconv.Atoi(scanner.Text())
	eprint(d)
	scanner.Scan()
	l, _ := strconv.Atoi(scanner.Text())
	eprint(l)
	eprint("=============== END INPUT ===============")
	eprint("")

	solution := d
	for i := 1; i <= l; i++ {
		solution += 5
	}

	print(solution)
}

func main() {
	contestResponse()
}

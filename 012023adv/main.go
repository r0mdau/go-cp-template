package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func eprint(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func print(a ...interface{}) {
	fmt.Println(a...)
}

func getFirstDigit(l string) (int, error) {
	for i := 0; i < len(l); i++ {
		if unicode.IsDigit(rune(l[i])) {
			return strconv.Atoi(string(l[i]))
		}
	}
	return 0, nil
}

func getLastDigit(l string) (int, error) {
	for i := len(l) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(l[i])) {
			return strconv.Atoi(string(l[i]))
		}
	}
	return 0, nil
}

func contestResponse() {
	eprint("=============== BEGIN INPUT ===============")

	var solution int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		first, _ := getFirstDigit(line)
		last, _ := getLastDigit(line)

		solution += first*10 + last
	}

	print(solution)
}

func main() {
	contestResponse()
}

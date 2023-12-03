package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func eprint(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func print(a ...interface{}) {
	fmt.Println(a...)
}

// var strRune = map[string]rune{
// 	"one":   '1',
// 	"two":   '2',
// 	"three": '3',
// 	"four":  '4',
// 	"five":  '5',
// 	"six":   '6',
// 	"seven": '7',
// 	"eight": '8',
// 	"nine":  '9',
// }

var nMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getIntFromLetters(letters string) (int, bool) {
	if len(letters) >= 3 {
		for k, v := range nMap {
			if strings.Contains(letters, k) {
				return v, true
			}
		}
	}
	return 0, false
}

func getFirstDigit(l string) (int, error) {
	var letters string
	for i := 0; i < len(l); i++ {
		if unicode.IsDigit(rune(l[i])) {
			return strconv.Atoi(string(l[i]))
		} else {
			letters += string(l[i])
			number, found := getIntFromLetters(letters)
			if found {
				return number, nil
			}
		}
	}
	return 0, nil
}

func getLastDigit(l string) (int, error) {
	var letters string
	for i := len(l) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(l[i])) {
			return strconv.Atoi(string(l[i]))
		} else {
			letters = string(l[i]) + letters
			number, found := getIntFromLetters(letters)
			if found {
				return number, nil
			}
		}
	}
	return 0, nil
}

func contestResponse() {

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

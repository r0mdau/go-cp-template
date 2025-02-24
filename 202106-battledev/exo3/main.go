package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func eprint(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func print(a ...interface{}) {
	fmt.Println(a...)
}

func holePos(s []string) int {
	for i := range s {
		if s[i] == "." {
			return i
		}
	}
	return 0
}

func contestResponse() {
	eprint("=============== BEGIN INPUT ===============")

	scanner := bufio.NewScanner(os.Stdin)

	// initialize vars
	m := [20][]string{}
	hashtagsCnt := [20]int{}
	i := 0
	for scanner.Scan() {
		current := scanner.Text()
		eprint(current)
		values := strings.Split(current, "")
		m[i] = values
		cnt := 0
		for j := range values {
			if values[j] == "#" {
				cnt++
			}
		}
		hashtagsCnt[i] = cnt
		i++
	}

	eprint("=============== END INPUT ===============")
	eprint("")

	hole := 0
	doAble := false
	Exit:
		for j := range m {
			// if this line and previous 3 count 9 hashtags AND the '.' (hole) is at the same position
			// we looks good, but next there are Exit conditions
			if j >= 3 && hashtagsCnt[j] == 9 && hashtagsCnt[j-1] == 9 && hashtagsCnt[j-2] == 9 && hashtagsCnt[j-3] == 9 && holePos(m[j]) == holePos(m[j-1]) && holePos(m[j]) == holePos(m[j-2]) && holePos(m[j]) == holePos(m[j-3]) {
				hole = holePos(m[j])
				// but exit if there is a hole under...
				if j+1 < len(m) - 1 && m[j+1][hole] == "." {
					break Exit
				}
				// but exit if there is a hashtag above
				for z := j - 4; z >= 0; z-- {
					if m[z][hole] == "#" {
						break Exit
					}
				}
				doAble = true
				print(fmt.Sprintf("BOOM %d", hole + 1))
				break Exit
			}
		}
	if !doAble {
		print("NOPE")
	}
}

func main() {
	contestResponse()
}

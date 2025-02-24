package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func eprint(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func print(a ...interface{}) {
	fmt.Println(a...)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func mapWaste(s []string) map[string]int {
	m := make(map[string]int)
	for i := range s {
		if _, ok := m[s[i]]; ok {
			m[s[i]]++
		} else {
			m[s[i]] = 1
		}
	}
	return m
}

func halveMap(m map[string]int) map[string]int {
	for k := range m {
		m[k] /= 2
	}
	return m
}

func contestResponse() {

	// read inputs with bigger buffer size for large orbits...
	// was stuck here during contest :(
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	orbitLen, _ := strconv.Atoi(strings.TrimSpace(readLine(reader)))
	orbitWaste := strings.Split(strings.TrimSpace(readLine(reader)), "")

	// initialize vars
	orbitMap := halveMap(mapWaste(orbitWaste))
	halfWaste := mapWaste(orbitWaste[:orbitLen/2])


	cnt := 0
	for i := 0; i < orbitLen/2; i++ {
		// check if halfWaste contains half of each waste letter
		if reflect.DeepEqual(orbitMap, halfWaste) {
			cnt++
		}

		// rotate waste
		halfWaste[orbitWaste[i]]--
		halfWaste[orbitWaste[i+(orbitLen/2)]]++
	}

	fmt.Println(cnt*2)
}

func main() {
	contestResponse()
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func sum(s []int) int {
	t := 0
	for i := range s {
		t += s[i]
	}
	return t
}

func max(s []int) int {
	m := 0
	for i := range s {
		if m < s[i] {
			m = s[i]
		}
	}
	return m
}

func maxint(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func contestResponse() {
	// read inputs
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	constraints := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	strAsteroids := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	// initialize vars
	N, _ := strconv.Atoi(constraints[0])
	A, _ := strconv.Atoi(constraints[1])
	C, _ := strconv.Atoi(constraints[2])
	dp := make([]int, N+A+C+1)
	asteroids := make([]int, N)
	for i := range strAsteroids {
		asteroids[i], _ = strconv.Atoi(strAsteroids[i])
	}

	// solve
	absorb := sum(asteroids[:A])
	for i := 0; i < N; i++ {
		dp[i+1] = maxint(dp[i+1], dp[i])
		dp[i+A+C] = maxint(dp[i+A+C+1], dp[i]+absorb)
		if i+A < len(asteroids) {
			absorb += asteroids[i+A]
		}
		absorb -= asteroids[i]
	}

	fmt.Println(sum(asteroids) - max(dp))
}

func main() {
	contestResponse()
}

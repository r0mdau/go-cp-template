package main

import (
	"bufio"
	"fmt"
	"math"
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

func calculateLength(m map[int32][]int32, temp []int32, v int32, visited map[int32]bool) []int32 {
	visited[v] = true
	temp = append(temp, v)
	for _, val := range m[v] {
		if visited[val] == false {
			temp = calculateLength(m, temp, val, visited)
		}
	}
	return temp
}

func componentsInGraph(gb [][]int32) []int32 {
	// Write your code here
	var min, max int32
	min = math.MaxInt32
	edges := make(map[int32][]int32)
	for i := 0; i < len(gb); i++ {
		index := gb[i][0]
		value := gb[i][1]
		edges[index] = append(edges[index], value)
		edges[value] = append(edges[value], index)
	}
	eprint(edges)

	visited := make(map[int32]bool)
	for i := range edges {
		visited[i] = false
	}

	for i := range edges {
		var length int32
		length = 1

		if visited[i] == false {
			temp := []int32{}
			length = int32(len(calculateLength(edges, temp, i, visited)))
		}

		if length < min && length > 1 {
			min = length
		}
		if length > max {
			max = length
		}
	}
	return []int32{min, max}
}

func contestResponse() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // first line is not needed
	n := scanner.Text()
	eprint(n)

	// initialize vars
	var gb [][]int32

	for scanner.Scan() {
		current := scanner.Text()
		values := strings.Split(current, " ")
		value0, _ := strconv.Atoi(values[0])
		value1, _ := strconv.Atoi(values[1])
		gb = append(gb, []int32{int32(value0), int32(value1)})
		//eprint(values)
		// do something with current
	}
	eprint("gb", gb)

	eprint("")
	print(componentsInGraph(gb))
}

func main() {
	contestResponse()
}

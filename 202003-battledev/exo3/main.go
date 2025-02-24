package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func eprint(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func print(a ...interface{}) {
	fmt.Println(a...)
}

func stringToTime(format, str string) time.Time {
	tm, err := time.Parse(format, str)
	if err != nil {
		eprint("Failed to decode time:", err)
	}
	return tm
}

type ByFirstDate [][2]time.Time

func (a ByFirstDate) Len() int {
	return len(a)
}
func (a ByFirstDate) Less(i, j int) bool {
	return a[i][0].Before(a[j][0])
}
func (a ByFirstDate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

const HourMin = "15:04"

func contestResponse() {
	eprint("=============== BEGIN INPUT ===============")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // first line is not needed

	// initialize vars
	m := make(map[int][][2]time.Time)
	// map contains each day as key
	for i := 1; i <= 5; i++ {
		m[i] = [][2]time.Time{}
	}

	// read inputs
	for scanner.Scan() {
		current := scanner.Text()
		values := strings.Split(current, " ")
		day, _ := strconv.Atoi(values[0])
		hours := strings.Split(values[1], "-")
		debut := stringToTime(HourMin, hours[0])
		fin := stringToTime(HourMin, hours[1])
		m[day] = append(m[day], [2]time.Time{debut, fin})
		eprint(current)
	}

	// sort slice by first date for each day in map
	for i := 1; i <= 5; i++ {
		sort.Sort(ByFirstDate(m[i]))
	}
	// for each slice of each day in map, outer current boundaries equal as previous, examples :
	//    |------| 1           |-----------------| 1         |------| 1
	//        |-------| 2           |-----| 2              |-----| 2
	// => |-----------| 2   => |-----------------| 2    => |--------| 2
	for i := 1; i <= 5; i++ {
		for j := 1; j < len(m[i]); j++ {
			if m[i][j-1][0].Before(m[i][j][0]) && m[i][j-1][1].After(m[i][j][0]) {
				m[i][j][0] = m[i][j-1][0]
			}
			if m[i][j-1][0].Before(m[i][j][1]) && m[i][j-1][1].After(m[i][j][1]) {
				m[i][j][1] = m[i][j-1][1]
			}
		}
	}


	eprint("=============== END INPUT ===============")

	eprint(m[2])

	print(solve(m))
}

func solve(m map[int][][2]time.Time) string {
	startDay := stringToTime(HourMin, "08:00")
	endDay := stringToTime(HourMin, "17:59")
	// for each day
	for d := 1; d <= 5; d++ {
		hours := m[d]
		numDay := d
		// for each slice hours in each day
		for i := range hours {
			// if first slice hours, compare to startDay 08:00
			if i == 0 && hours[i][0].After(startDay) {
				sub := hours[i][0].Sub(startDay)
				if sub >= time.Hour {
					eprint("i == 0")
					return fmt.Sprintf("%d %s-%s", numDay, startDay.Format(HourMin), startDay.Add(time.Minute*59).Format(HourMin))
				}
			}
			// compare current slice hour with previous
			if i > 0 && hours[i-1][1].Before(hours[i][0]) {
				previous := hours[i-1]
				sub := hours[i][0].Sub(previous[1])
				if sub >= time.Hour && previous[1].Add(time.Hour).Before(endDay) {
					eprint("i au milieu")
					return fmt.Sprintf("%d %s-%s", numDay, previous[1].Add(time.Minute).Format(HourMin), previous[1].Add(time.Hour).Format(HourMin))
				}
			}
			// if last slice hours, compare to endDay 17:59
			if i == len(hours)-1 && endDay.After(hours[i][1]) && hours[i][1].After(hours[i-1][1]) {
				sub := endDay.Sub(hours[i][1])
				if sub >= time.Hour && hours[i][1].Add(time.Hour).Before(endDay) {
					eprint("i == len last")
					return fmt.Sprintf("%d %s-%s", numDay, hours[i][1].Add(time.Minute).Format(HourMin), hours[i][1].Add(time.Hour).Format(HourMin))
				}
			}
		}
	}
	return "GAME OVER"
}

func main() {
	contestResponse()
}

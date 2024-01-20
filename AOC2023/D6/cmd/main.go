package main

import (
	"fmt"
	"github.com/DF-wu/AdventOfCode-Practice/dfmisc"
	"strconv"
	"strings"
)

func main() {
	input := dfmisc.Filereader("input.txt")
	fmt.Println(p1(input))
	fmt.Println(p2(input))
	//fmt.Println(input)
}

/*
Time:      71530
Distance:  940200
*/
func p2(input string) int {
	time, dist := 0, 0
	// eat all spaces
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "Time:", "")
	input = strings.ReplaceAll(input, "Distance:", "")
	time, _ = strconv.Atoi(strings.Fields(input)[0])
	dist, _ = strconv.Atoi(strings.Fields(input)[1])
	fmt.Println(time, dist)
	ans := 0
	for holdTime := 1; holdTime < time; holdTime++ {
		releaseTime := time - holdTime
		// cal distance   ↓time      ↓speed
		totalDist := releaseTime * holdTime
		if totalDist > dist {
			// count as valid race
			ans++
		}

	}
	fmt.Println(ans)
	return 0
}

/*
Time:      7  15   30
Distance:  9  40  200
*/
func p1(input string) int {
	// try to read token in more general way
	twoLines := strings.Split(input, "\n")
	records := map[int]int{}
	recordTime := strings.Fields(twoLines[0])
	recordDist := strings.Fields(twoLines[1])
	// load data as a map
	for i := 1; i < len(recordTime); i++ {
		//convert string to int
		rt, _ := strconv.Atoi(recordTime[i])
		rd, _ := strconv.Atoi(recordDist[i])
		records[rt] = rd
	}
	// times i Hava is hold time + release time
	// hold time -> car at 0 speed but 1m/s to charge energy
	// release time -> car at hold time m/s to run rest of time
	ans := 0
	currentRace := 0
	for time, dist := range records {
		currentRace = 0
		fmt.Println(time, dist)
		for holdtime := 1; holdtime < time; holdtime++ {
			releaseTime := time - holdtime
			// cal distance   ↓time      ↓speed
			totalDist := releaseTime * holdtime
			if totalDist > dist {
				// count as valid race
				currentRace++
			}
		}
		if ans == 0 {
			ans = currentRace
		} else {
			ans *= currentRace
		}

	}
	fmt.Println(records)
	fmt.Println(ans)
	return 0

}

package day2

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func count(s string) (twos int, threes int) {
	tw := 0
	th := 0

	freq := map[string]int{}
	for _, rn := range s {
		freq[string(rn)]++
	}

	for _, v := range freq {
		if v == 2 {
			tw++
		}
		if v == 3 {
			th++
		}
	}

	return min(1, tw), min(1, th)
}

func partOne(arr []string) {
	twos := 0
	threes := 0

	for _, boxID := range arr {
		tw, th := count(boxID)
		twos += tw
		threes += th
	}

	fmt.Printf("Hash: %d\n", twos*threes)
}

// Number of differing characters
func match(a string, b string) string {
	similar := ""
	// All the box ids are the same length
	for i, rn := range a {
		if string(rn) == string(b[i]) {
			similar += string(rn)
		}
	}
	return similar
}

func longestMatch(a string, arr []string) string {
	longest := ""
	for _, str := range arr {
		m := match(a, str)
		if len(m) > len(longest) {
			longest = m
		}
	}
	return longest
}

func longest(a []string) string {
	str := ""
	for _, s := range a {
		if len(s) > len(str) {
			str = s
		}
	}
	return str
}

func partTwo(arr []string) {
	matches := []string{}
	for index, str := range arr {
		if index < len(arr) && index+1 < len(arr) {
			matches = append(matches, longestMatch(str, arr[(index+1):len(arr)-1]))
		}
	}

	fmt.Printf("Closest Match: %s", longest(matches))
}

// Solve the puzzle for day 2
func Solve() {
	dat, err := ioutil.ReadFile("./day2/data")
	check(err)
	arr := strings.Split(string(dat), "\n")

	partOne(arr)
	partTwo(arr)
}

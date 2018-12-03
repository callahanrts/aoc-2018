package day1

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Solve - solve the puzzle for day 1
func Solve() {
	dat, err := ioutil.ReadFile("./day1/data")
	check(err)
	arr := strings.Split(string(dat), "\n")

	freq := 0
	for _, el := range arr {
		i, _ := strconv.Atoi(el)
		freq += i
	}
	fmt.Printf("Part 1: %d\n", freq)

	table := map[string]bool{}
	index := 0
	freq = 0

	for !table[strconv.Itoa(freq)] {
		table[strconv.Itoa(freq)] = true
		i, _ := strconv.Atoi(arr[index%(len(arr)-1)])
		freq += i
		index++
	}
	fmt.Printf("Part 2: %d", freq)
}

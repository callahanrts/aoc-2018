package day7

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Step - keep track of steps

func parseSteps(filename string) (int, map[string][]string) {
	deps := make(map[string][]string)
	rawLogs, _ := ioutil.ReadFile(filename)
	arr := strings.Split(string(rawLogs), "\n")
	for _, raw := range arr {
		if len(raw) > 0 {
			step := string(raw[36])
			dep := string(raw[5])
			deps[step] = append(deps[step], dep)
		}
	}
	return len(arr) - 1, deps
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func toChar(i int) rune {
	return rune('A' - 1 + i)
}

func intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y]--
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}

// Solve - solve it
func Solve() {
	// Part One
	_, prereq := parseSteps("./day7/steps")
	has := []string{}

	S := []string{}
	for i := 1; i <= 26; i++ {
		S = append(S, string(toChar(i)))
	}

	for i := 0; i < 101; i++ {
		for _, j := range S {
			scheduled := stringInSlice(j, has)
			if !scheduled && equal(intersection(prereq[j], has), prereq[j]) {
				fmt.Printf(j)
				has = append(has, j)
				break
			}
		}
	}

	// Part Two
}

package day12

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type rule struct {
	Note string // ..#.. (Order plant/no-plant)
	Val  string // # or .
}

func parseState(filename string) (string, []rule) {
	f, _ := ioutil.ReadFile(filename)
	arr := strings.Split(string(f), "\n")
	re := regexp.MustCompile("initial state: (.*$)")
	match := re.FindStringSubmatch(arr[0])
	return match[1], parseNotes(arr[2:])
}

func parseNotes(data []string) []rule {
	notes := []rule{}
	for _, str := range data {
		if len(str) > 0 {
			s := strings.Split(str, " => ")
			r := rule{
				Note: s[0],
				Val:  s[1],
			}
			notes = append(notes, r)
		}
	}
	return notes
}

func matchRule(s string, r rule) bool {
	return s == r.Note
}

func iterate(state string, rules []rule) string {
	st := []rune(state)
	nst := []rune("")
	for i := 0; i < len(st)-2; i++ {
		if i < 2 {
			nst = append(nst, '.')
			continue
		}
		match := false
		var mr rule
		for _, r := range rules {
			match = match || matchRule(string(st[i-2:i+3]), r)
			if match {
				mr = r
				break
			}
		}
		if match {
			nst = append(nst, []rune(mr.Val)[0])
		} else {
			nst = append(nst, '.')
		}
	}
	return string(nst)
}

func score(s string) int {
	total := 0
	for i, c := range []rune(s) {
		if string(c) == "#" {
			total += i - 10
		}
	}
	return total
}

// Solve day 12
func Solve() {
	state, rules := parseState("./day12/pots")
	state = ".........." + state + "......"
	prev := 0
	for i := 0; i < 20; i++ {
		fmt.Printf("%d: %d change: %d\n", i, score(state), score(state)-prev)
		prev = score(state)
		state = iterate(state+"...", rules)
	}
	fmt.Print(score(state))

	// Part 2:
	// After 102 generations, the increase is a constant 69
	// 102: 9306
	// (50000000000 - 102) * 69 + 9306
}

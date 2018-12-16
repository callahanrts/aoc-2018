package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func calculateScore(s1, s2 int) []int {
	score := []int{}
	s := strconv.Itoa(s1 + s2)
	t := strings.Split(s, "")
	for i := 0; i < len(t); i++ {
		d, _ := strconv.Atoi(t[i])
		score = append(score, d)
	}
	return score
}

func generate(recipes []int, e1, e2 int) ([]int, int, int) {
	recipes = append(recipes, calculateScore(recipes[e1], recipes[e2])...)
	e1 = (1 + e1 + recipes[e1]) % len(recipes)
	e2 = (1 + e2 + recipes[e2]) % len(recipes)
	return recipes, e1, e2
}

func printScores(recipes []int, e1, e2 int) {
	red := color.New(color.FgRed).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	for i := 0; i < len(recipes); i++ {
		if i == e1 {
			fmt.Print(red(recipes[i]), " ")
		} else if i == e2 {
			fmt.Print(cyan(recipes[i]), " ")
		} else {
			fmt.Print(recipes[i], " ")
		}
	}
	fmt.Print("\n")
}

// Solve day 14
func Solve() {
	e1 := 0
	e2 := 1
	recipes := []int{3, 7}
	printScores(recipes, e1, e2)
	after := 9
	after = 20681901

	// Part 1
	for len(recipes) < 10+after {
		recipes, e1, e2 = generate(recipes, e1, e2)
	}
	fmt.Print(recipes[after : after+10])

	// Part 2
	input := []int{6, 8, 1, 9, 0, 1}
	matched := 0
	idx := 0
	for i := 0; i < len(recipes); i++ {
		if recipes[i] == input[0] {
			matched = 1
			idx = i
		} else if matched > 0 && recipes[i] == input[matched] {
			matched++
		} else {
			idx = 0
			matched = 0
		}
		if matched == len(input) {
			fmt.Print("\nFound: ", idx)
			break
		}
	}
}

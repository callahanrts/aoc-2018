package day5

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readPolymer(filename string) []byte {
	str, _ := ioutil.ReadFile(filename)
	return str
}

func refactorPolymer(bytes []byte) ([]byte, int) {
	result := []byte{}
	changes := 0
	for i := 0; i < len(bytes); i++ {
		refactor := false
		if i+1 < len(bytes) {
			this := string(bytes[i])
			next := string(bytes[i+1])
			if this != next && strings.ToLower(this) == strings.ToLower(next) {
				refactor = true
			}
		}
		if refactor {
			i++
			changes++
		} else {
			result = append(result, bytes[i])
		}
	}
	return result, changes
}

func fullRefactor(bytes []byte) (int, int) {
	result, changes := refactorPolymer(bytes)
	for changes != 0 {
		result, changes = refactorPolymer(result)
	}
	return len(string(result)) - 1, changes
}

// Solve the puzzle for day 3
func Solve() {
	bytes := readPolymer("./day5/polymer")

	l, c := fullRefactor(bytes)
	fmt.Printf("%d | %d\n", l, c)
	polymer := string(bytes)

	var minChar string
	minLength := l
	for i := 1; i <= 26; i++ {
		char := string(rune('A' - 1 + i))
		lower := strings.ToLower(char)
		str := strings.Replace(polymer, char, "", -1)
		str = strings.Replace(str, lower, "", -1)
		length, _ := fullRefactor([]byte(str))
		fmt.Printf("%s: %d\n", char, length)
		if length < minLength {
			minLength = length
			minChar = char
		}

	}
	fmt.Printf("MIN: %s: %d", minChar, minLength)
}

package day11

import (
	"fmt"
	"strconv"
)

func powerAtCell(serial, x, y int) int {
	rackID := x + 10
	power := rackID * y
	power += serial
	power *= rackID
	str := []rune(strconv.Itoa(power))
	power, _ = strconv.Atoi(string(str[len(str)-3]))
	return power - 5
}

func powerGrid(serial, x, y, size int) int {
	sum := 0
	for y2 := 0; y2 < size; y2++ {
		for x2 := 0; x2 < size; x2++ {
			p := powerAtCell(serial, x+x2, y2+y)
			sum += p
		}
	}
	return sum
}

func power(serial, size int) (int, int, int) {
	maxX := 0
	maxY := 0
	max := powerGrid(serial, 0, 0, size)
	for y := 0; y <= 300-size; y++ {
		for x := 0; x <= 300-size; x++ {
			p := powerGrid(serial, x, y, size)
			if p > max {
				maxX = x
				maxY = y
				max = p
			}
		}
	}
	return max, maxX, maxY
}

func maxSize(serial int) (int, int, int) {
	max, maxX, maxY := power(serial, 1)
	maxSize := 0
	for i := 1; i <= 300; i++ {
		val, x, y := power(serial, i)
		if val > max {
			maxSize = i
			max = val
			maxX = x
			maxY = y
			fmt.Printf("Power: %d | size: %d (%d, %d)\n", max, maxSize, maxX, maxY)
		}
	}
	return maxSize, maxX, maxY
}

// Solve day 11
func Solve() {
	max, x, y := power(5791, 300)
	fmt.Printf("(%d, %d) %d", x, y, max)

	// Uncomment for part 2. It takes a while :/
	// fmt.Printf("\n\n(%d, %d) %d == %d\n", x, y, max, 29)
	// fmt.Print(maxSize(5791))
}

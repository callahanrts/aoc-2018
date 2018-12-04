package day3

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// Claim - An elves claim to Santa's suit fabric
type Claim struct {
	ID     int
	Width  int
	Height int
	X      int
	Y      int
}

func toInt(val string) int {
	i, _ := strconv.Atoi(val)
	return i
}

func parseClaims(file string) []Claim {
	claims := []Claim{}
	rawClaims, _ := ioutil.ReadFile(file)
	arr := strings.Split(string(rawClaims), "\n")
	for _, raw := range arr {
		re := regexp.MustCompile("#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)")
		// #ID @ X,Y: WidthxHeight
		match := re.FindStringSubmatch(raw)
		if len(match) > 1 {
			data := match[1:]
			claim := Claim{
				ID:     toInt(data[0]),
				X:      toInt(data[1]),
				Y:      toInt(data[2]),
				Width:  toInt(data[3]),
				Height: toInt(data[4]),
			}
			claims = append(claims, claim)
		}
	}
	return claims
}

func cut(fabric [1000][1000]uint8, c Claim) ([1000][1000]uint8, int) {
	overlap := 0
	for i := c.X; i < c.X+c.Width; i++ {
		for j := c.Y; j < c.Y+c.Height; j++ {
			if fabric[j][i] == 1 {
				overlap++
			}
			fabric[j][i]++
		}
	}
	return fabric, overlap
}

func hasOverlap(fabric [1000][1000]uint8, c Claim) bool {
	for i := c.X; i < c.X+c.Width; i++ {
		for j := c.Y; j < c.Y+c.Height; j++ {
			if fabric[j][i] > 1 {
				return true
			}
		}
	}

	return false
}

// Solve the puzzle for day 3
func Solve() {
	fabric := [1000][1000]uint8{}
	claims := parseClaims("./day3/claims")

	overlap := 0
	for _, claim := range claims {
		over := 0
		fabric, over = cut(fabric, claim)
		overlap += over
	}

	for _, claim := range claims {
		if !hasOverlap(fabric, claim) {
			fmt.Printf("No Overlap: %d\n", claim.ID)
		}
	}

	fmt.Printf("%d\n", overlap)
	fmt.Printf("Total: %d", len(claims))
}

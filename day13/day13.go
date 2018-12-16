package day13

// Not my best work, but it solved the puzzle. It looks like some people were
// using complex numbers to keep track of the states of the cars. Super clever.
// At least this was fun to watch the carts move throughout the track.

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"sort"
	"strings"
)

type cart struct {
	ID        int
	X         int
	Y         int
	NextTurn  string
	Direction string
	Moved     bool
}

func (c *cart) turnRight() {
	switch c.Direction {
	case ">":
		c.Direction = "v"
		break
	case "v":
		c.Direction = "<"
		break
	case "^":
		c.Direction = ">"
		break
	case "<":
		c.Direction = "^"
		break
	}
}

func (c *cart) turnLeft() {
	switch c.Direction {
	case ">":
		c.Direction = "^"
		break
	case "v":
		c.Direction = ">"
		break
	case "^":
		c.Direction = "<"
		break
	case "<":
		c.Direction = "v"
		break
	}
}

func (c *cart) turnNext(dir string) {
	switch dir {
	case "L":
		c.turnLeft()
		c.NextTurn = "S"
		break
	case "S":
		c.NextTurn = "R"
		break
	case "R":
		c.turnRight()
		c.NextTurn = "L"
		break
	}
}

func (c *cart) move(track map[int]map[int]string) {
	switch c.Direction {
	case "^":
		c.Y--
		break
	case ">":
		c.X++
		break
	case "v":
		c.Y++
		break
	case "<":
		c.X--
		break
	}

	s := track[c.Y][c.X]
	if s == "\\" {
		if c.Direction == ">" || c.Direction == "<" {
			c.turnRight()
		} else { // ^
			c.turnLeft()
		}
	}
	if s == "/" {
		if c.Direction == "^" || c.Direction == "v" {
			c.turnRight()
		} else {
			c.turnLeft()
		}
	}
	if s == "+" {
		c.turnNext(c.NextTurn)
	}
	c.Moved = true
}

func parseCart(x, y, id int, c string) cart {
	return cart{
		ID:        id,
		X:         x,
		Y:         y,
		Direction: c,
		NextTurn:  "L",
		Moved:     false,
	}
}

func isCorner(piece string) bool {
	return piece == "/" || piece == "\\" || piece == "+"
}

func isVertical(piece string) bool {
	return isCorner(piece) || piece == "|"
}
func isHorizontal(piece string) bool {
	return isCorner(piece) || piece == "-"
}
func trackPiece(x, y int, track map[int]map[int]string) string {
	l := track[y][x-1]
	r := track[y][x+1]
	t := track[y-1][x]
	b := track[y+1][x]
	if isVertical(t) && isHorizontal(l) && isHorizontal(r) && isVertical(b) {
		return "+"
	} else if isVertical(t) && isVertical(b) {
		return "|"
	} else if isHorizontal(l) && isHorizontal(r) {
		return "-"
	} else if isVertical(b) && isHorizontal(r) { // Top Left Corner
		return "/"
	} else if isVertical(b) && isHorizontal(l) { // Top Right Corner
		return "\\"
	} else if isVertical(t) && isHorizontal(l) { // Bottom Right Corner
		return "/"
	} else if isVertical(t) && isHorizontal(r) { // Bottom Left Corner
		return "\\"
	}
	return "?"
}

func parseTrack(filename string) (map[int]map[int]string, []cart) {
	track := map[int]map[int]string{}
	carts := []cart{}
	f, _ := ioutil.ReadFile(filename)
	arr := strings.Split(string(f), "\n")
	total := 0
	for row, line := range arr {
		if track[row] == nil {
			track[row] = map[int]string{}
		}
		for col, c := range []rune(line) {
			switch c {
			case '>', '<':
				carts = append(carts, parseCart(col, row, total, string(c)))
				total++
				break
			case '^', 'v':
				carts = append(carts, parseCart(col, row, total, string(c)))
				total++
			}
			track[row][col] = string(c)
		}
	}

	for r := 0; r < len(track); r++ {
		for c := 0; c < len(track[r]); c++ {
			s := track[r][c]
			if s == ">" || s == "v" || s == "<" || s == "^" {
				track[r][c] = trackPiece(c, r, track)
			}
		}
	}

	return track, carts
}

func removeAtPoint(x, y int, carts []cart) {
	for i := 0; i < len(carts); i++ {
		if carts[i].X == x && carts[i].Y == y {
			carts[i] = parseCart(-1, -1, -1, ".")
		}
	}
}

func tick(track map[int]map[int]string, carts []cart) {
	sort.Slice(carts, func(i, j int) bool {
		return carts[i].X < carts[j].X
	})
	sort.Slice(carts, func(i, j int) bool {
		return carts[i].Y < carts[j].Y
	})
	for i := 0; i < len(carts); i++ {
		if !carts[i].Moved {
			carts[i].move(track)
		}
		cs := collisions(carts)
		if len(cs) > 0 {
			// fmt.Printf("(%d, %d)\n", carts[i].X, carts[i].Y)
			removeAtPoint(carts[i].X, carts[i].Y, carts)
			// panic("collision") // Part 1
		}
	}
	total := 0
	for i := 0; i < len(carts); i++ {
		if carts[i].X > -1 {
			total++
		}
		carts[i].Moved = false
	}
	if total == 1 {
		fmt.Printf("\n(%d, %d)\n", carts[len(carts)-1].X, carts[len(carts)-1].Y)
		panic("Last Cart")
	}
	return
}

func cartInPos(x, y int, carts []cart) (bool, cart) {
	var crt cart
	for _, c := range carts {
		if c.X == x && c.Y == y {
			return true, c
		}
	}
	return false, crt
}

func printTrack(track map[int]map[int]string, carts []cart, cs []collision) {
	fmt.Print("\n\n")
	red := color.New(color.FgRed).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	for row := 0; row < len(track); row++ {
		for col := 0; col < len(track[row]); col++ {
			inPos, crt := cartInPos(col, row, carts)
			inCollision := false
			for _, c := range cs {
				inCollision = col == c.X && row == c.Y
			}
			if inCollision {
				fmt.Print(red("X"))
			} else if inPos {
				fmt.Print(cyan(crt.Direction))
			} else {
				fmt.Print(track[row][col])
			}
		}
		fmt.Print("\n")
	}
}

type collision struct {
	X int
	Y int
}

func collisions(carts []cart) []collision {
	cs := []collision{}
	for _, c1 := range carts {
		for _, c2 := range carts {
			if c1.ID != c2.ID && c1.X == c2.X && c1.Y == c2.Y {
				cs = append(cs, collision{c1.X, c2.Y})
			}
		}
	}
	return cs
}

// Solve day13
func Solve() {
	track, carts := parseTrack("./day13/track")

	cs := []collision{}
	hit := false
	i := 0
	for !hit {
		tick(track, carts)
		hit = len(cs) > 0
		i++
	}
}

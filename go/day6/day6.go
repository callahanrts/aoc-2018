package day6

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// Point - coordinate from the wrist device
type Point struct {
	X  int
	Y  int
	ID int
}

func parseCoords(filename string) []Point {
	coords := []Point{}
	count := 1
	data, _ := ioutil.ReadFile(filename)
	arr := strings.Split(string(data), "\n")
	for _, el := range arr[0 : len(arr)-1] {
		newEl := strings.Replace(el, " ", "", -1)
		c := strings.Split(newEl, ",")
		x, _ := strconv.Atoi(c[0])
		y, _ := strconv.Atoi(c[1])
		coord := Point{
			X:  x,
			Y:  y,
			ID: count,
		}
		count++
		coords = append(coords, coord)
	}

	return coords
}

func taxiDiff(a Point, b Point) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

func max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxCoord(ps []Point) (int, int) {
	maxX := 0
	maxY := 0
	for _, p := range ps {
		maxX = max(p.X, maxX)
		maxY = max(p.Y, maxY)
	}
	return maxX, maxY
}

func plot(freq map[int]map[int]Point, x, y int) {
	for row := 0; row <= y; row++ {
		for col := 0; col <= x; col++ {
			id := freq[row][col].ID
			if id == -1 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", freq[row][col].ID)
			}
		}
		fmt.Printf("\n")
	}
}

func score(freq map[int]map[int]Point, ps []Point, x, y int) map[int]int {
	score := map[int]int{}
	for row := 0; row <= y; row++ {
		for col := 0; col <= x; col++ {
			point := freq[row][col]
			if isFinite(ps, point) && point.ID != -1 {
				score[freq[row][col].ID] = max(score[freq[row][col].ID]+1, 0)
			}
		}
	}
	return score
}

func closestPoint(ps []Point, x, y int) Point {
	var minPoint Point
	curPoint := Point{x, y, -1}
	minDiff := -1

	for _, p := range ps {
		diff := taxiDiff(curPoint, p)
		if minDiff == -1 || diff < minDiff {
			minDiff = diff
			minPoint = p
		} else if minDiff == diff {
			minDiff = diff
			minPoint = curPoint
		}
	}
	return minPoint
}

func countFreq(ps []Point, x, y int) map[int]map[int]Point {
	score := make(map[int]map[int]Point)
	for row := 0; row <= y; row++ {
		if score[row] == nil {
			score[row] = make(map[int]Point)
		}
		for col := 0; col <= x+1; col++ {
			cp := closestPoint(ps, col, row)
			score[row][col] = cp
		}
	}

	return score
}

func isFinite(ps []Point, p Point) bool {
	var tr, tl, br, bl bool
	for _, k := range ps {
		tr = tr || k.X > p.X && k.Y < p.Y
		tl = tl || k.X < p.X && k.Y < p.Y
		br = br || k.X > p.X && k.Y > p.Y
		bl = bl || k.X < p.X && k.Y > p.Y
	}
	if tr && tl && br && bl {
		return true
	}
	return false
}

func maxScore(score map[int]int) int {
	mx := 0
	for _, val := range score {
		mx = max(val, mx)
	}
	return mx
}

// Solve - Puzzle for day 6
func Solve() {
	coords := parseCoords("./day6/coords")
	x, y := maxCoord(coords)
	freq := countFreq(coords, x, y)
	s := score(freq, coords, x, y)
	max := maxScore(s)
	fmt.Print("Part 1: ", max)
}

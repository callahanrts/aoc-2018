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

func plot(coords []Point, x int, y int) {
	count := 1
	maxX := 0
	maxY := 0
	p := map[int]map[int]string{}
	for _, point := range coords {
		if p[point.Y] == nil {
			p[point.Y] = map[int]string{}
		}
		p[point.Y][point.X] = strconv.Itoa(count)
		maxX = max(maxX, point.X)
		maxY = max(maxY, point.Y)
		count++
	}
	for row := 0; row <= y; row++ {
		for col := 0; col <= x; col++ {
			point := p[row][col]
			if len(point) > 0 {
				fmt.Printf(point)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func removeInfinite(ps []Point) []Point {
	points := []Point{}
	for _, p := range ps {
		t := false
		r := false
		b := false
		l := false
		for _, k := range ps {
			r = r || p.X < k.X
			l = l || p.X > k.X
			b = b || p.Y < k.Y
			t = t || p.Y > k.Y
		}
		if r && l && b && t {
			points = append(points, p)
		}
	}
	return points
}

func scorePoints(ps []Point, x int, y int) map[int]int {
	// fmt.Printf("%d, %d", x, y)
	score := map[int]int{}
	// fmt.Printf("\n")
	for row := 0; row < y+1; row++ {
		for col := 0; col <= x+1; col++ {
			tmp := ps[0]
			minDiff := max(x, y)
			for _, p := range ps {
				diff := taxiDiff(Point{X: col, Y: row}, p)
				if diff < minDiff {
					minDiff = diff
					tmp = p
				} else if diff == minDiff {
					tmp = Point{}
				}
			}
			if tmp.ID > 0 {
				fmt.Print(tmp.ID)
				score[tmp.ID] = score[tmp.ID] + 1
			} else {
				fmt.Print(".")
			}
		}
		fmt.Printf("\n")
	}
	// fmt.Printf("\n")
	// fmt.Println(score)
	return score
}

// Solve - Puzzle for day 6
func Solve() {
	coords := parseCoords("./day6/test")
	x, y := maxCoord(coords)
	fmt.Println(coords)
	plot(coords, x, y)
	fmt.Printf("\n")
	score := scorePoints(coords, x, y)
	cs := removeInfinite(coords)

	mx := 0
	for _, p := range cs {
		mx = max(score[p.ID], mx)
		fmt.Printf("%d: %d\n", p.ID, score[p.ID])
	}
}

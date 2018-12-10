package day10

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type vector struct {
	X        int
	Y        int
	Velocity []int
}

func (v *vector) Move() {
	v.X += v.Velocity[0]
	v.Y += v.Velocity[1]
}

func move(vs []vector) []vector {
	for i := 0; i < len(vs); i++ {
		vs[i].Move()
	}
	return vs
}

func gridSize(vs []vector) (int, int, int, int) {
	maxX := vs[0].X
	maxY := vs[0].Y
	minX := vs[0].X
	minY := vs[0].Y
	for _, v := range vs {
		if maxX < v.X {
			maxX = v.X
		}
		if minX > v.X {
			minX = v.X
		}
		if maxY < v.Y {
			maxY = v.Y
		}
		if minY > v.Y {
			minY = v.Y
		}
	}
	return minX, maxX, minY, maxY
}

func printVectors(vs []vector) {
	mp := map[int]map[int]string{}
	minx, maxx, miny, maxy := gridSize(vs)
	for _, v := range vs {
		if mp[v.X] == nil {
			mp[v.X] = map[int]string{}
		}
		mp[v.X][v.Y] = "#"
	}

	for row := miny; row <= maxy; row++ {
		for col := minx; col <= maxx; col++ {
			if len(mp[col][row]) > 0 {
				fmt.Print(mp[col][row])
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

// position=< 9,  1> velocity=< 0,  2>
func parseVectors(filename string) []vector {
	vs := []vector{}
	data, _ := ioutil.ReadFile(filename)
	arr := strings.Split(string(data), "\n")
	for _, el := range arr[0 : len(arr)-1] {
		re := regexp.MustCompile("position=<\\s*(-*\\d+),\\s*(-*\\d+)>.*velocity=<\\s*(-*\\d+),\\s*(-*\\d+)>")
		match := re.FindStringSubmatch(el)
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		vx, _ := strconv.Atoi(match[3])
		vy, _ := strconv.Atoi(match[4])
		v := vector{
			X:        x,
			Y:        y,
			Velocity: []int{vx, vy},
		}
		vs = append(vs, v)
	}
	return vs
}

func overZero(vs []vector) bool {
	lineHeight := 11
	minY := vs[0].Y
	maxY := vs[0].Y
	for _, v := range vs {
		if v.Y < minY {
			minY = v.Y
		}
		if v.Y > maxY {
			maxY = v.Y
		}
	}
	return maxY-minY < lineHeight
}

// Solve day 10
func Solve() {
	vs := parseVectors("./day10/input")

	count := 0
	for !overZero(vs) {
		vs = move(vs)
		count++
	}
	printVectors(vs)
	fmt.Printf("\nTook %ds", count)
}

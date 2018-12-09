package day9

import (
	"fmt"
)

type node struct {
	Val  int
	Next *node
	Prev *node
}

type list struct {
	Length  int
	Current *node
}

func (l *list) InsertNode(val int) {
	n := new(node)
	n.Val = val
	if l.Current == nil {
		n.Next = n
		n.Prev = n
	} else {
		n.Next = l.Current.Next
		n.Prev = l.Current
		l.Current.Next.Prev = n
		l.Current.Next = n
	}
	l.Current = n
	l.Length++
}

func (l *list) RemoveNode() node {
	n := l.Current
	l.Current.Prev.Next = l.Current.Next
	l.Current = l.Current.Next
	l.Length--
	return *n
}

func (l *list) Next() {
	l.Current = l.Current.Next
}

func (l *list) Prev() {
	l.Current = l.Current.Prev
}

func printList(l list) {
	fmt.Print("[")
	for i := 0; i < l.Length; i++ {
		l.Next()
		fmt.Print(l.Current.Val, ",")
	}
	fmt.Print("]\n")
}

func playGame(players int, lastMarble int) map[int]int {
	scores := map[int]int{}
	l := list{}
	l.InsertNode(0)
	for i := 1; i < lastMarble; i++ {
		player := ((i - 1) % players) + 1
		if i > 0 && i%23 == 0 {
			for j := 1; j <= 7; j++ {
				l.Prev()
			}
			n := l.RemoveNode()
			scores[player] += n.Val + i
		} else {
			l.Next()
			l.InsertNode(i)
		}
	}
	return scores
}

func maxScore(scores map[int]int) int {
	max := 0
	for _, score := range scores {
		if score > max {
			max = score
		}
	}
	return max
}

// Solve day 9
func Solve() {
	scores := playGame(410, 72059)
	fmt.Print(maxScore(scores), "\n")

	scores = playGame(410, 72059*100)
	fmt.Print(maxScore(scores), "\n")
}

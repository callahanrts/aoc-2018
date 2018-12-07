package day7

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Step - keep track of steps

func parseSteps(filename string) (int, map[string][]string) {
	deps := make(map[string][]string)
	rawLogs, _ := ioutil.ReadFile(filename)
	arr := strings.Split(string(rawLogs), "\n")
	for _, raw := range arr {
		if len(raw) > 0 {
			step := string(raw[36])
			dep := string(raw[5])
			deps[step] = append(deps[step], dep)
		}
	}
	return len(arr) - 1, deps
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func toChar(i int) rune {
	return rune('A' - 1 + i)
}

func intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y]--
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}

func stepOrder(prereq map[string][]string) []string {
	// Part One
	has := []string{}

	S := []string{}
	for i := 1; i <= 26; i++ {
		S = append(S, string(toChar(i)))
	}

	for i := 0; i <= len(prereq); i++ {
		for _, j := range S {
			scheduled := stringInSlice(j, has)
			if !scheduled && equal(intersection(prereq[j], has), prereq[j]) {
				fmt.Printf(j)
				has = append(has, j)
				break
			}
		}
	}
  fmt.Printf("\n")
  return has
}

func available(step string, done[]string, prereq map[string][]string) bool {
  avail := true
  for _, dep := range prereq[step] {
    avail = avail && !stringInSlice(dep, done)
  }
  return avail
}

func stepTime (step string) int {
  return int([]rune(step)[0]) - 64
}

func consume (workers [2][]string, step string) ([2][]string, bool) {
  // For each worker
  for w := 0; w < 2; w++ {
    if len(workers[w]) == 0 {
      for j := 0; j < stepTime(step); j++ {
        workers[w] = append(workers[w], step)
        return workers, true
      }
    }
  }

  return workers, false
}

func work (workers [2][]string) [2][]string {
  for w := 0; w < 2; w++ {
    if len(workers[w]) > 0 {
      workers[w] = workers[w][1:]
    }
  }

  return workers
}


// Solve - solve it
func Solve() {
	_, prereq := parseSteps("./day7/test")
  steps := stepOrder(prereq)
  done := []string{}

  workers := [2][]string{}

  t := 0
  for len(steps) > 0 {
    // If the next step is available
    if available(steps[0], done, prereq) {

      // Consume the next task if possible
      var consumed bool
      workers, consumed = consume(workers, steps[0])
      if consumed {
        fmt.Print(consumed)
        done = append(done, steps[0])
        steps = steps[1:]
      }

    }

    workers = work(workers)

    t++
  }

  fmt.Printf("\n%d", t)
}

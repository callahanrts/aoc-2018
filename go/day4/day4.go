package day4

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Log - A log of sleepy elves
type Log struct {
	Date  time.Time
	Event string
	ID    int
}

func toInt(val string) int {
	i, _ := strconv.Atoi(val)
	return i
}

func parseGuardID(log string) int {
	re := regexp.MustCompile("#(\\d+)")
	match := re.FindStringSubmatch(log)
	if len(match) > 0 {
		id, err := strconv.Atoi(match[1])
		if err == nil {
			return id
		}
	}
	return -1
}

func parseLogDate(log string) time.Time {
	re := regexp.MustCompile("\\[(.*)\\]")
	match := re.FindStringSubmatch(log)
	date, _ := time.Parse("2006-01-02 15:04", match[1])
	return date
}

func parseLogEvent(log string) string {
	re := regexp.MustCompile("(wake|sleep|shift)")
	match := re.FindStringSubmatch(log)
	return match[1]
}

func parseLogs(file string) []Log {
	logs := []Log{}
	rawLogs, _ := ioutil.ReadFile(file)
	arr := strings.Split(string(rawLogs), "\n")
	for _, raw := range arr {
		if len(raw) > 0 {
			log := Log{
				Date:  parseLogDate(raw),
				Event: parseLogEvent(raw),
				ID:    parseGuardID(raw),
			}

			logs = append(logs, log)
		}
	}

	sort.Slice(logs, func(i, j int) bool {
		return logs[i].Date.Before(logs[j].Date)
	})
	return logs
}

func sleepTimes(logs []Log) map[int]map[int]int {
	sleep := map[int]map[int]int{}
	guard := -1
	startTime := time.Now()
	for _, log := range logs {
		switch event := log.Event; event {
		case "shift":
			guard = log.ID
		case "sleep":
			startTime = log.Date
		case "wake":
			if sleep[guard] == nil {
				sleep[guard] = map[int]int{}
			}
			for i := startTime.Minute(); i < log.Date.Minute(); i++ {
				sleep[guard][i]++
			}
		}
		// fmt.Println(log)
	}
	return sleep
}

func countDurations(durations []time.Duration) time.Duration {
	var total time.Duration
	for _, dur := range durations {
		total = total + dur
	}
	return total
}

func longestSleeper(sleep map[int]map[int]int) (int, int, int) {
	max := 0
	guard := 0
	mode := 0
	for id, records := range sleep {
		minutes := 0
		maxFreq := 0
		localMode := 0
		for minute, freq := range records {
			if freq > maxFreq {
				maxFreq = freq
				localMode = minute
			}
			minutes += freq
		}
		if minutes > max {
			max = minutes
			guard = id
			mode = localMode
		}
	}
	return guard, max, mode
}

func maxMinute(sleep map[int]map[int]int) (int, int) {
	minute := 0
	guard := 0
	max := 0
	for id, records := range sleep {
		for min, freq := range records {
			if freq > max {
				guard = id
				minute = min
				max = freq
			}
		}
	}

	return guard, minute
}

// Solve the puzzle for day 3
func Solve() {
	logs := parseLogs("./day4/logs")
	sleep := sleepTimes(logs)

	guard, max, mode := longestSleeper(sleep)
	fmt.Printf("guard: %d, max: %d, mode: %d", guard, max, mode)
	fmt.Printf("\nGuard * Mode: %d", guard*mode)

	gd, maxMin := maxMinute(sleep)
	fmt.Printf("\nGuard: %d, minute: %d", gd, maxMin)
	fmt.Printf("\nGuard * Minute: %d", gd*maxMin)

	fmt.Printf("\nTotal: %d", len(logs))
}

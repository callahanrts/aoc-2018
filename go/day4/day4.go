package day4

import (
	"fmt"
	"io/ioutil"
	"regexp"
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

func parseClaims(file string) []Log {
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
	return logs
}

// Solve the puzzle for day 3
func Solve() {
	logs := parseClaims("./day4/test")

	// for _, log := range logs {
	// }

	fmt.Printf("Total: %d", len(logs))
}

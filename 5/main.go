package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	filename := args[0]

	p1(filename)
}

func parseDB(filename string) ([][2]int, []int) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	freshList := make([][2]int, 0)
	checkList := make([]int, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			id, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatalf("failed to parse checklist id: %s", err)
			}
			checkList = append(checkList, id)
			continue
		}
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("failed to parse range start: %s", err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("failed to parse range end: %s", err)
		}
		freshList = append(freshList, [2]int{start, end})
	}

	return freshList, checkList
}

func isFresh(id int, freshList [][2]int) bool {
	for _, r := range freshList {
		if id >= r[0] && id <= r[1] {
			return true
		}
	}
	return false
}

func p1(filename string) {
	freshList, checkList := parseDB(filename)
	freshCount := 0
	for _, id := range checkList {
		if isFresh(id, freshList) {
			freshCount++
		}
	}

	fmt.Println("fresh ingredients:", freshCount)
}

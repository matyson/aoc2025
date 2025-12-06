package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	filename := args[0]

	p1(filename)
	p2(filename)
}

func parseDB(filename string) ([][2]int, []int) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n\n")
	freshList := make([][2]int, 0)
	checkList := make([]int, 0)

	for _, line := range strings.Split(lines[0], "\n") {
		parts := strings.Split(line, "-")

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

	for _, line := range strings.Split(lines[1], "\n") {
		id, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			log.Fatalf("failed to parse ingredient id: %s", err)
		}
		checkList = append(checkList, id)
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

func p2(filename string) {
	freshList, _ := parseDB(filename)

	slices.SortFunc(freshList, func(a, b [2]int) int {
		if a[0] < b[0] {
			return -1
		} else if a[0] > b[0] {
			return 1
		}
		return 0
	})

	mergedList := make([][2]int, 0)
	for _, r := range freshList {
		if len(mergedList) == 0 {
			mergedList = append(mergedList, r)
			continue
		}

		lastRange := mergedList[len(mergedList)-1]
		if r[0] <= lastRange[1]+1 {
			if r[1] > lastRange[1] {
				mergedList[len(mergedList)-1][1] = r[1]
			}
		} else {
			mergedList = append(mergedList, r)
		}
	}

	totalFresh := 0
	for _, r := range mergedList {
		totalFresh += r[1] - r[0] + 1
	}
	fmt.Println("total fresh ingredients:", totalFresh)
}

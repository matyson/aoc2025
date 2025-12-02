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

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	idList := strings.Split(string(data), ",")
	sum := p1(idList)

	fmt.Println("the sum of all invalid IDs is:", sum)
}

func p1(idList []string) int {
	sum := 0
	for _, id := range idList {
		r := strings.Split(id, "-")
		min, err := strconv.Atoi(r[0])
		if err != nil {
			log.Fatalf("failed to convert min to int: %s", err)
		}
		max, err := strconv.Atoi(r[1])
		if err != nil {
			log.Fatalf("failed to convert max to int: %s", err)
		}

		for i := min; i <= max; i++ {
			if !isValid(i) {
				sum += i
			}
		}
	}
	return sum
}

func isValid(id int) bool {
	sId := strconv.Itoa(id)
	n := len(sId)
	if n%2 != 0 {
		return true
	}
	min := sId[0 : n/2]
	max := sId[n/2 : n]

	return min != max
}

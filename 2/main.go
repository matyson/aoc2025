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

	fmt.Println("the sum of all invalid IDs was:", sum)

	sum = p2(idList)
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
			if !isValid(i, 2) {
				sum += i
			}
		}
	}
	return sum
}

func isValid(id int, slices int) bool {
	sId := strconv.Itoa(id)
	n := len(sId)
	if n%slices != 0 {
		return true
	}
	sliceLen := n / slices
	parts := make([]string, slices)
	for i := 0; i < slices; i++ {
		parts[i] = sId[i*sliceLen : (i+1)*sliceLen]
	}
	valid := false
	for i := 1; i < len(parts); i++ {
		if parts[i] != parts[i-1] {
			valid = true
			break
		}
	}
	return valid
}

func p2(idList []string) int {
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
			for slices := 2; slices <= len(strconv.Itoa(i)); slices++ {
				if !isValid(i, slices) {
					sum += i
					break
				}
			}
		}
	}
	return sum
}

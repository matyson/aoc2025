package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	filename := args[0]

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	joltage := 0

	for scanner.Scan() {
		line := scanner.Text()

		max, maxIdx := findMax(line)
		if maxIdx == len(line)-1 {
			secondMax, _ := findMax(line[:maxIdx])
			joltage += secondMax*10 + max
		} else {
			secondMax, _ := findMax(line[maxIdx+1:])
			joltage += max*10 + secondMax
		}

	}
	fmt.Println("maximum joltage is:", joltage)
}

func findMax(line string) (int, int) {
	max := 0
	maxIdx := -1

	for i, char := range line {
		val, err := strconv.Atoi(string(char))
		if err != nil {
			log.Fatalf("failed to convert char to int: %s", err)
		}
		if val > max {
			max = val
			maxIdx = i
		}
	}
	return max, maxIdx

}

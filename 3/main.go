package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	filename := args[0]

	p1(filename)
	p2(filename)
}

func p1(filename string) {
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

func p2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	joltage := 0
	n := 12

	for scanner.Scan() {
		line := scanner.Text()

		sliceIdx := 0
		for i := n; i >= 1; i-- {
			max, maxIdx := findMax(line[sliceIdx : len(line)-i+1])
			joltage += max * int(math.Pow10(i-1))
			sliceIdx += maxIdx + 1
		}

	}
	fmt.Println("PLUS ULTRA:", joltage)

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

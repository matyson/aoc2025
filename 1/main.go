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

	p1(filename)
	p2(filename)

}

func p1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer f.Close()

	pass := 0
	dial := 50
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("failed to convert steps to int: %s", err)
		}

		switch direction {
		case "L":
			dial = ((dial-steps)%100 + 100) % 100
		case "R":
			dial = (dial + steps) % 100
		}

		if dial == 0 {
			pass++
		}

	}
	fmt.Println("the password was:", pass)

}

func p2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer f.Close()

	pass := 0
	dial := 50
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("failed to convert steps to int: %s", err)
		}
		var cycles int
		switch direction {
		case "L":
			if dial == 0 {
				cycles = steps / 100
			}

			diff := dial - steps
			if diff <= 0 && dial > 0 {
				cycles = -diff/100 + 1
			}

			dial = ((dial-steps)%100 + 100) % 100

		case "R":
			cycles = (dial + steps) / 100
			dial = (dial + steps) % 100
		}

		pass += cycles

	}
	fmt.Println("the password is:", pass)
}

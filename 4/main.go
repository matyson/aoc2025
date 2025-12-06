package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	grid := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, 0)
		for _, r := range line {
			row = append(row, string(r))
		}
		grid = append(grid, row)
	}

	paper := "@"
	M := len(grid)    // M rows
	N := len(grid[0]) // N columns
	sum := 0
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == paper {
				n := countPapers(grid, i, j, M, N, paper)
				if n < 4 {
					sum++
				}
			}
		}
	}

	fmt.Println("accessible papers:", sum)
}

func p2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, 0)
		for _, r := range line {
			row = append(row, string(r))
		}
		grid = append(grid, row)
	}

	paper := "@"
	M := len(grid)    // M rows
	N := len(grid[0]) // N columns
	sum := 0
	canRemove := true
	for canRemove {
		p := 0
		for i := 0; i < M; i++ {
			for j := 0; j < N; j++ {
				if grid[i][j] == paper {
					n := countPapers(grid, i, j, M, N, paper)
					if n < 4 {
						p++
						grid[i][j] = "x"
					}
				}
			}
		}
		if p == 0 {
			canRemove = false
		}

		sum += p
	}

	fmt.Println("removed papers:", sum)
}

func countPapers(grid [][]string, x, y, nx, ny int, paper string) int {
	count := 0
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for _, d := range directions {
		newX := x + d[0]
		newY := y + d[1]
		if newX >= 0 && newX < nx && newY >= 0 && newY < ny {
			if grid[newX][newY] == paper {
				count++
			}
		}
	}
	return count
}
